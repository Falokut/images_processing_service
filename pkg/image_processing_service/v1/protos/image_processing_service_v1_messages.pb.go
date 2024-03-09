// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: image_processing_service_v1_messages.proto

package protos

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ResampleFilter int32

const (
	// A high-quality resampling filter for photographic images yielding sharp results.
	ResampleFilter_Lanczos ResampleFilter = 0
	// A sharp cubic filter that is faster than Lanczos filter while providing similar results.
	ResampleFilter_CatmullRom ResampleFilter = 1
	// A cubic filter that produces smoother results with less ringing artifacts than CatmullRom.
	ResampleFilter_MitchellNetravali ResampleFilter = 2
	// Bilinear resampling filter, produces a smooth output. Faster than cubic filters.
	ResampleFilter_Linear ResampleFilter = 3
	// Simple and fast averaging filter appropriate for downscaling.
	// When upscaling it's similar to NearestNeighbor.
	ResampleFilter_Box ResampleFilter = 4
	// Fastest resampling filter, no antialiasing.
	ResampleFilter_NearestNeighbor ResampleFilter = 5
)

// Enum value maps for ResampleFilter.
var (
	ResampleFilter_name = map[int32]string{
		0: "Lanczos",
		1: "CatmullRom",
		2: "MitchellNetravali",
		3: "Linear",
		4: "Box",
		5: "NearestNeighbor",
	}
	ResampleFilter_value = map[string]int32{
		"Lanczos":           0,
		"CatmullRom":        1,
		"MitchellNetravali": 2,
		"Linear":            3,
		"Box":               4,
		"NearestNeighbor":   5,
	}
)

func (x ResampleFilter) Enum() *ResampleFilter {
	p := new(ResampleFilter)
	*p = x
	return p
}

func (x ResampleFilter) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResampleFilter) Descriptor() protoreflect.EnumDescriptor {
	return file_image_processing_service_v1_messages_proto_enumTypes[0].Descriptor()
}

func (ResampleFilter) Type() protoreflect.EnumType {
	return &file_image_processing_service_v1_messages_proto_enumTypes[0]
}

func (x ResampleFilter) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResampleFilter.Descriptor instead.
func (ResampleFilter) EnumDescriptor() ([]byte, []int) {
	return file_image_processing_service_v1_messages_proto_rawDescGZIP(), []int{0}
}

type BlurRequest_BlurMethod int32

const (
	BlurRequest_Box      BlurRequest_BlurMethod = 0
	BlurRequest_Gaussian BlurRequest_BlurMethod = 1
)

// Enum value maps for BlurRequest_BlurMethod.
var (
	BlurRequest_BlurMethod_name = map[int32]string{
		0: "Box",
		1: "Gaussian",
	}
	BlurRequest_BlurMethod_value = map[string]int32{
		"Box":      0,
		"Gaussian": 1,
	}
)

func (x BlurRequest_BlurMethod) Enum() *BlurRequest_BlurMethod {
	p := new(BlurRequest_BlurMethod)
	*p = x
	return p
}

func (x BlurRequest_BlurMethod) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BlurRequest_BlurMethod) Descriptor() protoreflect.EnumDescriptor {
	return file_image_processing_service_v1_messages_proto_enumTypes[1].Descriptor()
}

func (BlurRequest_BlurMethod) Type() protoreflect.EnumType {
	return &file_image_processing_service_v1_messages_proto_enumTypes[1]
}

func (x BlurRequest_BlurMethod) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BlurRequest_BlurMethod.Descriptor instead.
func (BlurRequest_BlurMethod) EnumDescriptor() ([]byte, []int) {
	return file_image_processing_service_v1_messages_proto_rawDescGZIP(), []int{8, 0}
}

type ImageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Image file as bytes
	Image []byte `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *ImageResponse) Reset() {
	*x = ImageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_image_processing_service_v1_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageResponse) ProtoMessage() {}

func (x *ImageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_image_processing_service_v1_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageResponse.ProtoReflect.Descriptor instead.
func (*ImageResponse) Descriptor() ([]byte, []int) {
	return file_image_processing_service_v1_messages_proto_rawDescGZIP(), []int{0}
}

func (x *ImageResponse) GetImage() []byte {
	if x != nil {
		return x.Image
	}
	return nil
}

type Image struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Image file as bytes (supports base64 encoding)
	Image []byte `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *Image) Reset() {
	*x = Image{}
	if protoimpl.UnsafeEnabled {
		mi := &file_image_processing_service_v1_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Image) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Image) ProtoMessage() {}

func (x *Image) ProtoReflect() protoreflect.Message {
	mi := &file_image_processing_service_v1_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Image.ProtoReflect.Descriptor instead.
func (*Image) Descriptor() ([]byte, []int) {
	return file_image_processing_service_v1_messages_proto_rawDescGZIP(), []int{1}
}

func (x *Image) GetImage() []byte {
	if x != nil {
		return x.Image
	}
	return nil
}

type CropRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image *Image `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	// rectangle coordinate x0
	StartX uint32 `protobuf:"varint,2,opt,name=StartX,json=start_x,proto3" json:"StartX,omitempty"`
	// rectangle coordinate y0
	StartY uint32 `protobuf:"varint,3,opt,name=StartY,json=start_y,proto3" json:"StartY,omitempty"`
	// rectangle coordinate x1
	EndX uint32 `protobuf:"varint,4,opt,name=EndX,json=end_x,proto3" json:"EndX,omitempty"`
	// rectangle coordinate y1
	EndY uint32 `protobuf:"varint,5,opt,name=EndY,json=end_y,proto3" json:"EndY,omitempty"`
}

func (x *CropRequest) Reset() {
	*x = CropRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_image_processing_service_v1_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CropRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CropRequest) ProtoMessage() {}

func (x *CropRequest) ProtoReflect() protoreflect.Message {
	mi := &file_image_processing_service_v1_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CropRequest.ProtoReflect.Descriptor instead.
func (*CropRequest) Descriptor() ([]byte, []int) {
	return file_image_processing_service_v1_messages_proto_rawDescGZIP(), []int{2}
}

func (x *CropRequest) GetImage() *Image {
	if x != nil {
		return x.Image
	}
	return nil
}

func (x *CropRequest) GetStartX() uint32 {
	if x != nil {
		return x.StartX
	}
	return 0
}

func (x *CropRequest) GetStartY() uint32 {
	if x != nil {
		return x.StartY
	}
	return 0
}

func (x *CropRequest) GetEndX() uint32 {
	if x != nil {
		return x.EndX
	}
	return 0
}

func (x *CropRequest) GetEndY() uint32 {
	if x != nil {
		return x.EndY
	}
	return 0
}

type ResizeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image          *Image         `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	ResampleFilter ResampleFilter `protobuf:"varint,2,opt,name=ResampleFilter,json=resample_filter,proto3,enum=image_processing_service.ResampleFilter" json:"ResampleFilter,omitempty"`
	Width          int32          `protobuf:"varint,3,opt,name=Width,json=width,proto3" json:"Width,omitempty"`
	Height         int32          `protobuf:"varint,4,opt,name=Height,json=height,proto3" json:"Height,omitempty"`
}

func (x *ResizeRequest) Reset() {
	*x = ResizeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_image_processing_service_v1_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResizeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResizeRequest) ProtoMessage() {}

func (x *ResizeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_image_processing_service_v1_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResizeRequest.ProtoReflect.Descriptor instead.
func (*ResizeRequest) Descriptor() ([]byte, []int) {
	return file_image_processing_service_v1_messages_proto_rawDescGZIP(), []int{3}
}

func (x *ResizeRequest) GetImage() *Image {
	if x != nil {
		return x.Image
	}
	return nil
}

func (x *ResizeRequest) GetResampleFilter() ResampleFilter {
	if x != nil {
		return x.ResampleFilter
	}
	return ResampleFilter_Lanczos
}

func (x *ResizeRequest) GetWidth() int32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *ResizeRequest) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

type ValidateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image     *Image `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	MaxWidth  *int32 `protobuf:"varint,2,opt,name=MaxWidth,json=max_width,proto3,oneof" json:"MaxWidth,omitempty"`
	MaxHeight *int32 `protobuf:"varint,3,opt,name=MaxHeight,json=max_height,proto3,oneof" json:"MaxHeight,omitempty"`
	MinWidth  *int32 `protobuf:"varint,4,opt,name=MinWidth,json=min_width,proto3,oneof" json:"MinWidth,omitempty"`
	MinHeight *int32 `protobuf:"varint,5,opt,name=MinHeight,json=min_height,proto3,oneof" json:"MinHeight,omitempty"`
	// List of image types, checking for the occurrence of an image type in the list
	SupportedTypes []string `protobuf:"bytes,6,rep,name=SupportedTypes,json=supported_types,proto3" json:"SupportedTypes,omitempty"`
}

func (x *ValidateRequest) Reset() {
	*x = ValidateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_image_processing_service_v1_messages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateRequest) ProtoMessage() {}

func (x *ValidateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_image_processing_service_v1_messages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateRequest.ProtoReflect.Descriptor instead.
func (*ValidateRequest) Descriptor() ([]byte, []int) {
	return file_image_processing_service_v1_messages_proto_rawDescGZIP(), []int{4}
}

func (x *ValidateRequest) GetImage() *Image {
	if x != nil {
		return x.Image
	}
	return nil
}

func (x *ValidateRequest) GetMaxWidth() int32 {
	if x != nil && x.MaxWidth != nil {
		return *x.MaxWidth
	}
	return 0
}

func (x *ValidateRequest) GetMaxHeight() int32 {
	if x != nil && x.MaxHeight != nil {
		return *x.MaxHeight
	}
	return 0
}

func (x *ValidateRequest) GetMinWidth() int32 {
	if x != nil && x.MinWidth != nil {
		return *x.MinWidth
	}
	return 0
}

func (x *ValidateRequest) GetMinHeight() int32 {
	if x != nil && x.MinHeight != nil {
		return *x.MinHeight
	}
	return 0
}

func (x *ValidateRequest) GetSupportedTypes() []string {
	if x != nil {
		return x.SupportedTypes
	}
	return nil
}

type ValidateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImageValid bool `protobuf:"varint,1,opt,name=ImageValid,json=image_valid,proto3" json:"ImageValid,omitempty"`
	// Returns when image not passed, a message describing why the image failed validation
	Details *string `protobuf:"bytes,2,opt,name=Details,json=details,proto3,oneof" json:"Details,omitempty"`
}

func (x *ValidateResponse) Reset() {
	*x = ValidateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_image_processing_service_v1_messages_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateResponse) ProtoMessage() {}

func (x *ValidateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_image_processing_service_v1_messages_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateResponse.ProtoReflect.Descriptor instead.
func (*ValidateResponse) Descriptor() ([]byte, []int) {
	return file_image_processing_service_v1_messages_proto_rawDescGZIP(), []int{5}
}

func (x *ValidateResponse) GetImageValid() bool {
	if x != nil {
		return x.ImageValid
	}
	return false
}

func (x *ValidateResponse) GetDetails() string {
	if x != nil && x.Details != nil {
		return *x.Details
	}
	return ""
}

type UserErrorMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=Message,json=message,proto3" json:"Message,omitempty"`
}

func (x *UserErrorMessage) Reset() {
	*x = UserErrorMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_image_processing_service_v1_messages_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserErrorMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserErrorMessage) ProtoMessage() {}

func (x *UserErrorMessage) ProtoReflect() protoreflect.Message {
	mi := &file_image_processing_service_v1_messages_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserErrorMessage.ProtoReflect.Descriptor instead.
func (*UserErrorMessage) Descriptor() ([]byte, []int) {
	return file_image_processing_service_v1_messages_proto_rawDescGZIP(), []int{6}
}

func (x *UserErrorMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type HueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image *Image `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	// between -360 and 360
	Hue int32 `protobuf:"varint,2,opt,name=hue,proto3" json:"hue,omitempty"`
}

func (x *HueRequest) Reset() {
	*x = HueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_image_processing_service_v1_messages_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HueRequest) ProtoMessage() {}

func (x *HueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_image_processing_service_v1_messages_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HueRequest.ProtoReflect.Descriptor instead.
func (*HueRequest) Descriptor() ([]byte, []int) {
	return file_image_processing_service_v1_messages_proto_rawDescGZIP(), []int{7}
}

func (x *HueRequest) GetImage() *Image {
	if x != nil {
		return x.Image
	}
	return nil
}

func (x *HueRequest) GetHue() int32 {
	if x != nil {
		return x.Hue
	}
	return 0
}

type BlurRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image      *Image                 `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	BlurRadius float64                `protobuf:"fixed64,2,opt,name=blur_radius,json=blurRadius,proto3" json:"blur_radius,omitempty"`
	Method     BlurRequest_BlurMethod `protobuf:"varint,3,opt,name=method,proto3,enum=image_processing_service.BlurRequest_BlurMethod" json:"method,omitempty"`
}

func (x *BlurRequest) Reset() {
	*x = BlurRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_image_processing_service_v1_messages_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlurRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlurRequest) ProtoMessage() {}

func (x *BlurRequest) ProtoReflect() protoreflect.Message {
	mi := &file_image_processing_service_v1_messages_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlurRequest.ProtoReflect.Descriptor instead.
func (*BlurRequest) Descriptor() ([]byte, []int) {
	return file_image_processing_service_v1_messages_proto_rawDescGZIP(), []int{8}
}

func (x *BlurRequest) GetImage() *Image {
	if x != nil {
		return x.Image
	}
	return nil
}

func (x *BlurRequest) GetBlurRadius() float64 {
	if x != nil {
		return x.BlurRadius
	}
	return 0
}

func (x *BlurRequest) GetMethod() BlurRequest_BlurMethod {
	if x != nil {
		return x.Method
	}
	return BlurRequest_Box
}

var File_image_processing_service_v1_messages_proto protoreflect.FileDescriptor

var file_image_processing_service_v1_messages_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69,
	0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x76, 0x31, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x25, 0x0a, 0x0d, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x22, 0x1d, 0x0a, 0x05, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x22, 0xa0, 0x01, 0x0a, 0x0b, 0x43,
	0x72, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x05, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x12, 0x17, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x72, 0x74, 0x58, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x78, 0x12, 0x17, 0x0a, 0x06, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x59, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x5f, 0x79, 0x12, 0x13, 0x0a, 0x04, 0x45, 0x6e, 0x64, 0x58, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x05, 0x65, 0x6e, 0x64, 0x5f, 0x78, 0x12, 0x13, 0x0a, 0x04, 0x45, 0x6e, 0x64, 0x59,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x65, 0x6e, 0x64, 0x5f, 0x79, 0x22, 0xc7, 0x01,
	0x0a, 0x0d, 0x52, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x35, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e,
	0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52,
	0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x51, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28,
	0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e,
	0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x0f, 0x72, 0x65, 0x73, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x57, 0x69, 0x64,
	0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12,
	0x16, 0x0a, 0x06, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0xb3, 0x02, 0x0a, 0x0f, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x05, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x05, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x12, 0x20, 0x0a, 0x08, 0x4d, 0x61, 0x78, 0x57, 0x69, 0x64, 0x74, 0x68, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x09, 0x6d, 0x61, 0x78, 0x5f, 0x77, 0x69, 0x64, 0x74,
	0x68, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x09, 0x4d, 0x61, 0x78, 0x48, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x48, 0x01, 0x52, 0x0a, 0x6d, 0x61, 0x78, 0x5f, 0x68,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a, 0x08, 0x4d, 0x69, 0x6e, 0x57,
	0x69, 0x64, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x48, 0x02, 0x52, 0x09, 0x6d, 0x69,
	0x6e, 0x5f, 0x77, 0x69, 0x64, 0x74, 0x68, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x09, 0x4d, 0x69,
	0x6e, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x48, 0x03, 0x52,
	0x0a, 0x6d, 0x69, 0x6e, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x88, 0x01, 0x01, 0x12, 0x27,
	0x0a, 0x0e, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x73,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65,
	0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x4d, 0x61, 0x78, 0x57,
	0x69, 0x64, 0x74, 0x68, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x4d, 0x61, 0x78, 0x48, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x4d, 0x69, 0x6e, 0x57, 0x69, 0x64, 0x74, 0x68, 0x42,
	0x0c, 0x0a, 0x0a, 0x5f, 0x4d, 0x69, 0x6e, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x5e, 0x0a,
	0x10, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1f, 0x0a, 0x0a, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x12, 0x1d, 0x0a, 0x07, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x88, 0x01,
	0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x2c, 0x0a,
	0x10, 0x55, 0x73, 0x65, 0x72, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x55, 0x0a, 0x0a, 0x48,
	0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x05, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x68, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x68,
	0x75, 0x65, 0x22, 0xd4, 0x01, 0x0a, 0x0b, 0x42, 0x6c, 0x75, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x35, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x6c, 0x75,
	0x72, 0x5f, 0x72, 0x61, 0x64, 0x69, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a,
	0x62, 0x6c, 0x75, 0x72, 0x52, 0x61, 0x64, 0x69, 0x75, 0x73, 0x12, 0x48, 0x0a, 0x06, 0x6d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x30, 0x2e, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x6c, 0x75, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x42, 0x6c, 0x75, 0x72, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x06, 0x6d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x22, 0x23, 0x0a, 0x0a, 0x42, 0x6c, 0x75, 0x72, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x12, 0x07, 0x0a, 0x03, 0x42, 0x6f, 0x78, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x47,
	0x61, 0x75, 0x73, 0x73, 0x69, 0x61, 0x6e, 0x10, 0x01, 0x2a, 0x6e, 0x0a, 0x0e, 0x52, 0x65, 0x73,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x0b, 0x0a, 0x07, 0x4c,
	0x61, 0x6e, 0x63, 0x7a, 0x6f, 0x73, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x61, 0x74, 0x6d,
	0x75, 0x6c, 0x6c, 0x52, 0x6f, 0x6d, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x4d, 0x69, 0x74, 0x63,
	0x68, 0x65, 0x6c, 0x6c, 0x4e, 0x65, 0x74, 0x72, 0x61, 0x76, 0x61, 0x6c, 0x69, 0x10, 0x02, 0x12,
	0x0a, 0x0a, 0x06, 0x4c, 0x69, 0x6e, 0x65, 0x61, 0x72, 0x10, 0x03, 0x12, 0x07, 0x0a, 0x03, 0x42,
	0x6f, 0x78, 0x10, 0x04, 0x12, 0x13, 0x0a, 0x0f, 0x4e, 0x65, 0x61, 0x72, 0x65, 0x73, 0x74, 0x4e,
	0x65, 0x69, 0x67, 0x68, 0x62, 0x6f, 0x72, 0x10, 0x05, 0x42, 0x24, 0x5a, 0x22, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_image_processing_service_v1_messages_proto_rawDescOnce sync.Once
	file_image_processing_service_v1_messages_proto_rawDescData = file_image_processing_service_v1_messages_proto_rawDesc
)

func file_image_processing_service_v1_messages_proto_rawDescGZIP() []byte {
	file_image_processing_service_v1_messages_proto_rawDescOnce.Do(func() {
		file_image_processing_service_v1_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_image_processing_service_v1_messages_proto_rawDescData)
	})
	return file_image_processing_service_v1_messages_proto_rawDescData
}

var file_image_processing_service_v1_messages_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_image_processing_service_v1_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_image_processing_service_v1_messages_proto_goTypes = []interface{}{
	(ResampleFilter)(0),         // 0: image_processing_service.ResampleFilter
	(BlurRequest_BlurMethod)(0), // 1: image_processing_service.BlurRequest.BlurMethod
	(*ImageResponse)(nil),       // 2: image_processing_service.ImageResponse
	(*Image)(nil),               // 3: image_processing_service.Image
	(*CropRequest)(nil),         // 4: image_processing_service.CropRequest
	(*ResizeRequest)(nil),       // 5: image_processing_service.ResizeRequest
	(*ValidateRequest)(nil),     // 6: image_processing_service.ValidateRequest
	(*ValidateResponse)(nil),    // 7: image_processing_service.ValidateResponse
	(*UserErrorMessage)(nil),    // 8: image_processing_service.UserErrorMessage
	(*HueRequest)(nil),          // 9: image_processing_service.HueRequest
	(*BlurRequest)(nil),         // 10: image_processing_service.BlurRequest
}
var file_image_processing_service_v1_messages_proto_depIdxs = []int32{
	3, // 0: image_processing_service.CropRequest.image:type_name -> image_processing_service.Image
	3, // 1: image_processing_service.ResizeRequest.image:type_name -> image_processing_service.Image
	0, // 2: image_processing_service.ResizeRequest.ResampleFilter:type_name -> image_processing_service.ResampleFilter
	3, // 3: image_processing_service.ValidateRequest.image:type_name -> image_processing_service.Image
	3, // 4: image_processing_service.HueRequest.image:type_name -> image_processing_service.Image
	3, // 5: image_processing_service.BlurRequest.image:type_name -> image_processing_service.Image
	1, // 6: image_processing_service.BlurRequest.method:type_name -> image_processing_service.BlurRequest.BlurMethod
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_image_processing_service_v1_messages_proto_init() }
func file_image_processing_service_v1_messages_proto_init() {
	if File_image_processing_service_v1_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_image_processing_service_v1_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_image_processing_service_v1_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Image); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_image_processing_service_v1_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CropRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_image_processing_service_v1_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResizeRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_image_processing_service_v1_messages_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_image_processing_service_v1_messages_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_image_processing_service_v1_messages_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserErrorMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_image_processing_service_v1_messages_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HueRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_image_processing_service_v1_messages_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlurRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_image_processing_service_v1_messages_proto_msgTypes[4].OneofWrappers = []interface{}{}
	file_image_processing_service_v1_messages_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_image_processing_service_v1_messages_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_image_processing_service_v1_messages_proto_goTypes,
		DependencyIndexes: file_image_processing_service_v1_messages_proto_depIdxs,
		EnumInfos:         file_image_processing_service_v1_messages_proto_enumTypes,
		MessageInfos:      file_image_processing_service_v1_messages_proto_msgTypes,
	}.Build()
	File_image_processing_service_v1_messages_proto = out.File
	file_image_processing_service_v1_messages_proto_rawDesc = nil
	file_image_processing_service_v1_messages_proto_goTypes = nil
	file_image_processing_service_v1_messages_proto_depIdxs = nil
}
