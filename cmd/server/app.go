package main

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"syscall"

	server "github.com/Falokut/grpc_rest_server"
	"github.com/Falokut/image_processing_service/internal/config"
	"github.com/Falokut/image_processing_service/internal/handler"
	"github.com/Falokut/image_processing_service/internal/imageprocessing/defaultprocessing"
	"github.com/Falokut/image_processing_service/internal/service"
	image_service "github.com/Falokut/image_processing_service/pkg/image_processing_service/v1/protos"
	jaegerTracer "github.com/Falokut/image_processing_service/pkg/jaeger"
	"github.com/Falokut/image_processing_service/pkg/logging"
	"github.com/Falokut/image_processing_service/pkg/metrics"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

func main() {
	logging.NewEntry(logging.FileAndConsoleOutput)
	logger := logging.GetLogger()

	cfg := config.GetConfig()
	log_level, err := logrus.ParseLevel(cfg.LogLevel)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Logger.SetLevel(log_level)
	var metric metrics.Metrics
	shutdown := make(chan error, 1)
	if cfg.EnableMetrics {
		tracer, closer, err := jaegerTracer.InitJaeger(cfg.JaegerConfig)
		if err != nil {
			logger.Errorf("cannot create tracer %v", err)
			return
		}
		logger.Info("Jaeger connected")
		defer closer.Close()

		opentracing.SetGlobalTracer(tracer)

		logger.Info("Metrics initializing")
		metric, err = metrics.CreateMetrics(cfg.PrometheusConfig.Name)
		if err != nil {
			logger.Errorf("error while creating metrics %v", err)
			return
		}

		go func() {
			logger.Info("Metrics server running")
			if err = metrics.RunMetricServer(cfg.PrometheusConfig.ServerConfig); err != nil {
				logger.Errorf("Shutting down, error while running metrics server %v", err)
				shutdown <- err
				return
			}
		}()
	} else {
		metric = &metrics.EmptyMetrics{}
	}

	logger.Info("Service initializing")
	s := service.NewImagesProcessingService(logger.Logger, defaultprocessing.ImageProcessing{})
	h := handler.NewImageProcessingServiceHandler(s)

	logger.Info("Server initializing")
	serv := server.NewServer(logger.Logger, h)
	go func() {
		if err := serv.Run(getListenServerConfig(cfg), metric, nil, nil); err != nil {
			logger.Errorf("Shutting down, error while running server %s", err.Error())
			shutdown <- err
			return
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGTERM)

	select {
	case <-quit:
		break
	case <-shutdown:
		break
	}

	serv.Shutdown()
}

const (
	kb = 8 << 10
	mb = kb << 10
)

func getListenServerConfig(cfg *config.Config) server.Config {
	return server.Config{
		Host:        cfg.Listen.Host,
		Port:        cfg.Listen.Port,
		Mode:        cfg.Listen.Mode,
		ServiceDesc: &image_service.ImageProcessingServiceV1_ServiceDesc,
		RegisterRestHandlerServer: func(ctx context.Context, mux *runtime.ServeMux, service any) error {
			serv, ok := service.(image_service.ImageProcessingServiceV1Server)
			if !ok {
				return errors.New("error while creating images processing service")
			}
			return image_service.RegisterImageProcessingServiceV1HandlerServer(ctx, mux, serv)
		},
		MaxRequestSize:  cfg.Listen.MaxRequestSize * mb,
		MaxResponceSize: cfg.Listen.MaxResponseSize * mb,
	}
}
