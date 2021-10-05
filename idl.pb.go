// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: idl.proto

package itu

import (
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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_idl_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_idl_proto_rawDescGZIP(), []int{0}
}

type ID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID *string `protobuf:"bytes,1,req,name=ID" json:"ID,omitempty"`
}

func (x *ID) Reset() {
	*x = ID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ID) ProtoMessage() {}

func (x *ID) ProtoReflect() protoreflect.Message {
	mi := &file_idl_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ID.ProtoReflect.Descriptor instead.
func (*ID) Descriptor() ([]byte, []int) {
	return file_idl_proto_rawDescGZIP(), []int{1}
}

func (x *ID) GetID() string {
	if x != nil && x.ID != nil {
		return *x.ID
	}
	return ""
}

type Course struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID     *string `protobuf:"bytes,1,req,name=ID" json:"ID,omitempty"`
	Name   *string `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
	Rating *int32  `protobuf:"varint,3,opt,name=Rating" json:"Rating,omitempty"`
}

func (x *Course) Reset() {
	*x = Course{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Course) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Course) ProtoMessage() {}

func (x *Course) ProtoReflect() protoreflect.Message {
	mi := &file_idl_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Course.ProtoReflect.Descriptor instead.
func (*Course) Descriptor() ([]byte, []int) {
	return file_idl_proto_rawDescGZIP(), []int{2}
}

func (x *Course) GetID() string {
	if x != nil && x.ID != nil {
		return *x.ID
	}
	return ""
}

func (x *Course) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *Course) GetRating() int32 {
	if x != nil && x.Rating != nil {
		return *x.Rating
	}
	return 0
}

type Courses struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Courses []*Course `protobuf:"bytes,1,rep,name=courses" json:"courses,omitempty"`
}

func (x *Courses) Reset() {
	*x = Courses{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Courses) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Courses) ProtoMessage() {}

func (x *Courses) ProtoReflect() protoreflect.Message {
	mi := &file_idl_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Courses.ProtoReflect.Descriptor instead.
func (*Courses) Descriptor() ([]byte, []int) {
	return file_idl_proto_rawDescGZIP(), []int{3}
}

func (x *Courses) GetCourses() []*Course {
	if x != nil {
		return x.Courses
	}
	return nil
}

var File_idl_proto protoreflect.FileDescriptor

var file_idl_proto_rawDesc = []byte{
	0x0a, 0x09, 0x69, 0x64, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x69, 0x74, 0x75,
	0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x14, 0x0a, 0x02, 0x49, 0x44, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x22,
	0x44, 0x0a, 0x06, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x52,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x30, 0x0a, 0x07, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73,
	0x12, 0x25, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0b, 0x2e, 0x69, 0x74, 0x75, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x07,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x32, 0xd5, 0x01, 0x0a, 0x0d, 0x43, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x12, 0x27, 0x0a, 0x0c, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x0b, 0x2e, 0x69, 0x74, 0x75, 0x2e,
	0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x1a, 0x0a, 0x2e, 0x69, 0x74, 0x75, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x12, 0x27, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x12, 0x0b, 0x2e, 0x69, 0x74, 0x75, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x1a,
	0x0a, 0x2e, 0x69, 0x74, 0x75, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x0c, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x07, 0x2e, 0x69, 0x74,
	0x75, 0x2e, 0x49, 0x44, 0x1a, 0x0a, 0x2e, 0x69, 0x74, 0x75, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x12, 0x25, 0x0a, 0x0d, 0x67, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x42, 0x79, 0x49,
	0x44, 0x12, 0x07, 0x2e, 0x69, 0x74, 0x75, 0x2e, 0x49, 0x44, 0x1a, 0x0b, 0x2e, 0x69, 0x74, 0x75,
	0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x0a, 0x67, 0x65, 0x74, 0x43, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x73, 0x12, 0x0a, 0x2e, 0x69, 0x74, 0x75, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x0c, 0x2e, 0x69, 0x74, 0x75, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x42,
	0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x69, 0x74, 0x75,
}

var (
	file_idl_proto_rawDescOnce sync.Once
	file_idl_proto_rawDescData = file_idl_proto_rawDesc
)

func file_idl_proto_rawDescGZIP() []byte {
	file_idl_proto_rawDescOnce.Do(func() {
		file_idl_proto_rawDescData = protoimpl.X.CompressGZIP(file_idl_proto_rawDescData)
	})
	return file_idl_proto_rawDescData
}

var file_idl_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_idl_proto_goTypes = []interface{}{
	(*Empty)(nil),   // 0: itu.Empty
	(*ID)(nil),      // 1: itu.ID
	(*Course)(nil),  // 2: itu.Course
	(*Courses)(nil), // 3: itu.Courses
}
var file_idl_proto_depIdxs = []int32{
	2, // 0: itu.Courses.courses:type_name -> itu.Course
	2, // 1: itu.CourseMethods.updateCourse:input_type -> itu.Course
	2, // 2: itu.CourseMethods.CreateCourse:input_type -> itu.Course
	1, // 3: itu.CourseMethods.deleteCourse:input_type -> itu.ID
	1, // 4: itu.CourseMethods.getCourseByID:input_type -> itu.ID
	0, // 5: itu.CourseMethods.getCourses:input_type -> itu.Empty
	0, // 6: itu.CourseMethods.updateCourse:output_type -> itu.Empty
	0, // 7: itu.CourseMethods.CreateCourse:output_type -> itu.Empty
	0, // 8: itu.CourseMethods.deleteCourse:output_type -> itu.Empty
	2, // 9: itu.CourseMethods.getCourseByID:output_type -> itu.Course
	3, // 10: itu.CourseMethods.getCourses:output_type -> itu.Courses
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_idl_proto_init() }
func file_idl_proto_init() {
	if File_idl_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_idl_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_idl_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ID); i {
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
		file_idl_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Course); i {
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
		file_idl_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Courses); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_idl_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_idl_proto_goTypes,
		DependencyIndexes: file_idl_proto_depIdxs,
		MessageInfos:      file_idl_proto_msgTypes,
	}.Build()
	File_idl_proto = out.File
	file_idl_proto_rawDesc = nil
	file_idl_proto_goTypes = nil
	file_idl_proto_depIdxs = nil
}
