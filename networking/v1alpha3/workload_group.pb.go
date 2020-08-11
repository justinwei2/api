// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: networking/v1alpha3/workload_group.proto

// `WorkloadGroup` describes a collection of workload instances.
// It provides a specification that the workload instances can use to bootstrap
// their proxies, including the metadata and identity. It is only intended to
// be used with non-k8s workloads like Virtual Machines, and is meant to mimic
// the existing sidecar injection and deployment specification model used for
// Kubernetes workloads to bootstrap Istio proxies.
//
// The following example declares a workload group representing a collection
// of workloads that will be registered under `reviews` in namespace
// `bookinfo`. The set of labels will be associated with each workload
// instance during the bootstrap process, and the ports 3550 and 8080
// will be associated with the workload group and use service account `default`.
// `app.kubernetes.io/version` is just an arbitrary example of a label.
//
// {{<tabset category-name="example">}}
// {{<tab name="v1alpha3" category-value="v1alpha3">}}
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: WorkloadGroup
// metadata:
//   name: reviews
//   namespace: bookinfo
// spec:
//   template:
//     metadata:
//       labels:
//         app.kubernetes.io/name: reviews
//         app.kubernetes.io/version: "1.3.4"
//     template:
//       ports:
//         grpc: 3550
//         http: 8080
//       serviceAccount: default
// ```
// {{</tab>}}
// {{</tabset>}}
//

package v1alpha3

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	_ "istio.io/gogo-genproto/googleapis/google/api"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// `WorkloadGroup` enables specifying the properties of a single workload for bootstrap and
// provides a template for `WorkloadEntry`, similar to how `Deployment` specifies properties
// of workloads via `Pod` templates.
// `WorkloadGroup` has no relationship to resources which control service registry like `ServiceEntry`
// and as such doesn't configure host name for these workloads.
//
// <!-- crd generation tags
// +cue-gen:WorkloadGroup:groupName:networking.istio.io
// +cue-gen:WorkloadGroup:version:v1alpha3
// +cue-gen:WorkloadGroup:storageVersion
// +cue-gen:WorkloadGroup:subresource:status
// +cue-gen:WorkloadGroup:scope:Namespaced
// +cue-gen:WorkloadGroup:resource:categories=istio-io,networking-istio-io,shortNames=wg,plural=workloadgroups
// +cue-gen:WorkloadGroup:printerColumn:name=Age,type=date,JSONPath=.metadata.creationTimestamp,description="CreationTimestamp is a timestamp
// representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations.
// Clients may not set this value. It is represented in RFC3339 form and is in UTC.
// Populated by the system. Read-only. Null for lists. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata"
// -->
//
// <!-- go code generation tags
// +kubetype-gen
// +kubetype-gen:groupVersion=networking.istio.io/v1alpha3
// +genclient
// +k8s:deepcopy-gen=true
// -->
type WorkloadGroup struct {
	// Template to be used for the generation of WorkloadEntry resources that belong
	// to this WorkloadGroup
	Template             *WorkloadEntryTemplate `protobuf:"bytes,1,opt,name=template,proto3" json:"template,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *WorkloadGroup) Reset()         { *m = WorkloadGroup{} }
func (m *WorkloadGroup) String() string { return proto.CompactTextString(m) }
func (*WorkloadGroup) ProtoMessage()    {}
func (*WorkloadGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_621fb79d92a8ed09, []int{0}
}
func (m *WorkloadGroup) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WorkloadGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WorkloadGroup.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WorkloadGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkloadGroup.Merge(m, src)
}
func (m *WorkloadGroup) XXX_Size() int {
	return m.Size()
}
func (m *WorkloadGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkloadGroup.DiscardUnknown(m)
}

var xxx_messageInfo_WorkloadGroup proto.InternalMessageInfo

func (m *WorkloadGroup) GetTemplate() *WorkloadEntryTemplate {
	if m != nil {
		return m.Template
	}
	return nil
}

// `WorkloadEntryTemplate` is the underlying workload template that is associated with the workload group.
// A `WorkloadGroup` can have more than one `WorkloadEntry`.
type WorkloadEntryTemplate struct {
	// Metadata that will be used for all corresponding WorkloadEntries.
	Metadata *ObjectMeta `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Describes the WorkloadEntry resources that will be created.
	// Please note that `address` and `labels` fields should not be set in the template. The `serviceAccount`
	// field MUST be specified. The workload identities (mTLS certificates) will be bootstrapped using the
	// specified service account's token. Workload entries in this group will be in the same namespace as the
	// workload group, and inherit the labels, and annotations from the group object.
	Template             *WorkloadEntry `protobuf:"bytes,2,opt,name=template,proto3" json:"template,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *WorkloadEntryTemplate) Reset()         { *m = WorkloadEntryTemplate{} }
func (m *WorkloadEntryTemplate) String() string { return proto.CompactTextString(m) }
func (*WorkloadEntryTemplate) ProtoMessage()    {}
func (*WorkloadEntryTemplate) Descriptor() ([]byte, []int) {
	return fileDescriptor_621fb79d92a8ed09, []int{1}
}
func (m *WorkloadEntryTemplate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WorkloadEntryTemplate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WorkloadEntryTemplate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WorkloadEntryTemplate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkloadEntryTemplate.Merge(m, src)
}
func (m *WorkloadEntryTemplate) XXX_Size() int {
	return m.Size()
}
func (m *WorkloadEntryTemplate) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkloadEntryTemplate.DiscardUnknown(m)
}

var xxx_messageInfo_WorkloadEntryTemplate proto.InternalMessageInfo

func (m *WorkloadEntryTemplate) GetMetadata() *ObjectMeta {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *WorkloadEntryTemplate) GetTemplate() *WorkloadEntry {
	if m != nil {
		return m.Template
	}
	return nil
}

// From k8s.io.apimachinery.pkg.apis.meta.v1.ObjectMeta.
type ObjectMeta struct {
	Labels               map[string]string `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Annotations          map[string]string `protobuf:"bytes,2,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ObjectMeta) Reset()         { *m = ObjectMeta{} }
func (m *ObjectMeta) String() string { return proto.CompactTextString(m) }
func (*ObjectMeta) ProtoMessage()    {}
func (*ObjectMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_621fb79d92a8ed09, []int{2}
}
func (m *ObjectMeta) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ObjectMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ObjectMeta.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ObjectMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjectMeta.Merge(m, src)
}
func (m *ObjectMeta) XXX_Size() int {
	return m.Size()
}
func (m *ObjectMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjectMeta.DiscardUnknown(m)
}

var xxx_messageInfo_ObjectMeta proto.InternalMessageInfo

func (m *ObjectMeta) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *ObjectMeta) GetAnnotations() map[string]string {
	if m != nil {
		return m.Annotations
	}
	return nil
}

func init() {
	proto.RegisterType((*WorkloadGroup)(nil), "istio.networking.v1alpha3.WorkloadGroup")
	proto.RegisterType((*WorkloadEntryTemplate)(nil), "istio.networking.v1alpha3.WorkloadEntryTemplate")
	proto.RegisterType((*ObjectMeta)(nil), "istio.networking.v1alpha3.ObjectMeta")
	proto.RegisterMapType((map[string]string)(nil), "istio.networking.v1alpha3.ObjectMeta.AnnotationsEntry")
	proto.RegisterMapType((map[string]string)(nil), "istio.networking.v1alpha3.ObjectMeta.LabelsEntry")
}

func init() {
	proto.RegisterFile("networking/v1alpha3/workload_group.proto", fileDescriptor_621fb79d92a8ed09)
}

var fileDescriptor_621fb79d92a8ed09 = []byte{
	// 356 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xd1, 0x4a, 0x02, 0x41,
	0x14, 0x86, 0x99, 0x95, 0x44, 0x67, 0x09, 0x64, 0x28, 0xb0, 0xbd, 0x30, 0x11, 0x82, 0xbd, 0x9a,
	0x4d, 0x85, 0xa8, 0x2e, 0x02, 0x85, 0x08, 0xc1, 0x08, 0x96, 0xa0, 0x08, 0x42, 0xce, 0xe6, 0xb4,
	0x4e, 0x8e, 0x3b, 0xcb, 0x3a, 0x1a, 0x3e, 0x50, 0x77, 0x3d, 0x48, 0x97, 0x3d, 0x82, 0xf8, 0x24,
	0xb1, 0xb3, 0xab, 0x2e, 0x61, 0x62, 0x77, 0xbb, 0x87, 0xff, 0xff, 0xfe, 0x7f, 0x0e, 0x07, 0xdb,
	0x01, 0x53, 0xef, 0x32, 0x1a, 0xf2, 0xc0, 0x77, 0xa6, 0x75, 0x10, 0xe1, 0x00, 0x9a, 0x4e, 0x3c,
	0x10, 0x12, 0xfa, 0x3d, 0x3f, 0x92, 0x93, 0x90, 0x86, 0x91, 0x54, 0x92, 0x1c, 0xf1, 0xb1, 0xe2,
	0x92, 0xae, 0xf5, 0x74, 0xa9, 0xb7, 0x8e, 0x7d, 0x29, 0x7d, 0xc1, 0x1c, 0x08, 0xb9, 0xf3, 0xca,
	0x99, 0xe8, 0xf7, 0x3c, 0x36, 0x80, 0x29, 0x97, 0x51, 0xe2, 0xb5, 0xb6, 0xa7, 0xb0, 0x40, 0x45,
	0xb3, 0x44, 0x59, 0x7b, 0xc6, 0xfb, 0x0f, 0xe9, 0xfc, 0x26, 0x0e, 0x27, 0x5d, 0x5c, 0x50, 0x6c,
	0x14, 0x0a, 0x50, 0xac, 0x8c, 0xaa, 0xc8, 0x36, 0x1b, 0xa7, 0xf4, 0xcf, 0x26, 0x74, 0xe9, 0xbd,
	0x8e, 0x91, 0xf7, 0xa9, 0xcf, 0x5d, 0x11, 0x6a, 0x1f, 0x08, 0x1f, 0x6e, 0xd4, 0x90, 0x16, 0x2e,
	0x8c, 0x98, 0x82, 0x3e, 0x28, 0x48, 0x73, 0x4e, 0xb6, 0xe4, 0xdc, 0x79, 0x6f, 0xec, 0x45, 0xdd,
	0x32, 0x05, 0xee, 0xca, 0x46, 0x3a, 0x99, 0xaa, 0x86, 0x46, 0xd8, 0xbb, 0x56, 0x6d, 0xe7, 0xe6,
	0x2d, 0x23, 0xd3, 0xf3, 0xd3, 0xc0, 0x78, 0x9d, 0x41, 0x3a, 0x38, 0x2f, 0xc0, 0x63, 0x62, 0x5c,
	0x46, 0xd5, 0x9c, 0x6d, 0x36, 0xea, 0x3b, 0x55, 0xa3, 0x5d, 0xed, 0xd1, 0x01, 0x6e, 0x0a, 0x20,
	0x8f, 0xd8, 0x84, 0x20, 0x90, 0x0a, 0x14, 0x97, 0xc1, 0xb8, 0x6c, 0x68, 0xde, 0xd9, 0x6e, 0xbc,
	0xd6, 0xda, 0x98, 0x40, 0xb3, 0x28, 0xeb, 0x02, 0x9b, 0x99, 0x40, 0x52, 0xc2, 0xb9, 0x21, 0x9b,
	0xe9, 0x5d, 0x16, 0xdd, 0xf8, 0x93, 0x1c, 0xe0, 0xbd, 0x29, 0x88, 0x49, 0xb2, 0x9c, 0xa2, 0x9b,
	0xfc, 0x5c, 0x1a, 0xe7, 0xc8, 0xba, 0xc2, 0xa5, 0xdf, 0xec, 0xff, 0xf8, 0xdb, 0xf4, 0x6b, 0x51,
	0x41, 0xdf, 0x8b, 0x0a, 0x9a, 0x2f, 0x2a, 0xe8, 0xa9, 0x9a, 0x3c, 0x86, 0x4b, 0x7d, 0x90, 0x1b,
	0x4e, 0xcf, 0xcb, 0xeb, 0x63, 0x6b, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0x8c, 0x5f, 0x21, 0x35,
	0xfe, 0x02, 0x00, 0x00,
}

func (m *WorkloadGroup) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WorkloadGroup) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WorkloadGroup) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Template != nil {
		{
			size, err := m.Template.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintWorkloadGroup(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *WorkloadEntryTemplate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WorkloadEntryTemplate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WorkloadEntryTemplate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Template != nil {
		{
			size, err := m.Template.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintWorkloadGroup(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Metadata != nil {
		{
			size, err := m.Metadata.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintWorkloadGroup(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ObjectMeta) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ObjectMeta) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ObjectMeta) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Annotations) > 0 {
		for k := range m.Annotations {
			v := m.Annotations[k]
			baseI := i
			i -= len(v)
			copy(dAtA[i:], v)
			i = encodeVarintWorkloadGroup(dAtA, i, uint64(len(v)))
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintWorkloadGroup(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintWorkloadGroup(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Labels) > 0 {
		for k := range m.Labels {
			v := m.Labels[k]
			baseI := i
			i -= len(v)
			copy(dAtA[i:], v)
			i = encodeVarintWorkloadGroup(dAtA, i, uint64(len(v)))
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintWorkloadGroup(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintWorkloadGroup(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintWorkloadGroup(dAtA []byte, offset int, v uint64) int {
	offset -= sovWorkloadGroup(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *WorkloadGroup) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Template != nil {
		l = m.Template.Size()
		n += 1 + l + sovWorkloadGroup(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *WorkloadEntryTemplate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Metadata != nil {
		l = m.Metadata.Size()
		n += 1 + l + sovWorkloadGroup(uint64(l))
	}
	if m.Template != nil {
		l = m.Template.Size()
		n += 1 + l + sovWorkloadGroup(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ObjectMeta) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Labels) > 0 {
		for k, v := range m.Labels {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovWorkloadGroup(uint64(len(k))) + 1 + len(v) + sovWorkloadGroup(uint64(len(v)))
			n += mapEntrySize + 1 + sovWorkloadGroup(uint64(mapEntrySize))
		}
	}
	if len(m.Annotations) > 0 {
		for k, v := range m.Annotations {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovWorkloadGroup(uint64(len(k))) + 1 + len(v) + sovWorkloadGroup(uint64(len(v)))
			n += mapEntrySize + 1 + sovWorkloadGroup(uint64(mapEntrySize))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovWorkloadGroup(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozWorkloadGroup(x uint64) (n int) {
	return sovWorkloadGroup(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *WorkloadGroup) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWorkloadGroup
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: WorkloadGroup: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WorkloadGroup: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Template", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkloadGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthWorkloadGroup
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWorkloadGroup
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Template == nil {
				m.Template = &WorkloadEntryTemplate{}
			}
			if err := m.Template.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWorkloadGroup(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWorkloadGroup
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthWorkloadGroup
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *WorkloadEntryTemplate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWorkloadGroup
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: WorkloadEntryTemplate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WorkloadEntryTemplate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkloadGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthWorkloadGroup
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWorkloadGroup
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = &ObjectMeta{}
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Template", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkloadGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthWorkloadGroup
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWorkloadGroup
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Template == nil {
				m.Template = &WorkloadEntry{}
			}
			if err := m.Template.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWorkloadGroup(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWorkloadGroup
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthWorkloadGroup
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ObjectMeta) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWorkloadGroup
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ObjectMeta: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ObjectMeta: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Labels", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkloadGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthWorkloadGroup
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWorkloadGroup
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Labels == nil {
				m.Labels = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowWorkloadGroup
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowWorkloadGroup
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthWorkloadGroup
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthWorkloadGroup
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowWorkloadGroup
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthWorkloadGroup
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthWorkloadGroup
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipWorkloadGroup(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthWorkloadGroup
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Labels[mapkey] = mapvalue
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Annotations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkloadGroup
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthWorkloadGroup
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWorkloadGroup
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Annotations == nil {
				m.Annotations = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowWorkloadGroup
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowWorkloadGroup
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthWorkloadGroup
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthWorkloadGroup
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowWorkloadGroup
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthWorkloadGroup
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthWorkloadGroup
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipWorkloadGroup(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthWorkloadGroup
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Annotations[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWorkloadGroup(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWorkloadGroup
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthWorkloadGroup
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipWorkloadGroup(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowWorkloadGroup
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowWorkloadGroup
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowWorkloadGroup
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthWorkloadGroup
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthWorkloadGroup
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowWorkloadGroup
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipWorkloadGroup(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthWorkloadGroup
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthWorkloadGroup = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowWorkloadGroup   = fmt.Errorf("proto: integer overflow")
)
