// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: networking/v1alpha3/workload_group.proto

// `WorkloadGroup` describes a collection of non-Kubernetes (VM) workload
// instances that will serve as identical applications. It provides a
// specification that the workload instances can use to bootstrap their
// proxies, including the metadata and identity.
//
// The following example declares a workload group representing a collection
// of VM workloads that will be registered under `details` in namespace
// `bootstore`. The set of labels will be associated with each workload
// instance during the VM bootstrap process, and the workloads will expose the
// ports 3550 and 8080 to Istio. The workload group will be under the network
// `remote` and service account `default`.
//
// {{<tabset category-name="example">}}
// {{<tab name="v1alpha3" category-value="v1alpha3">}}
// ```yaml
// apiVersion: networking.istio.io/v1alpha3
// kind: WorkloadGroup
// metadata:
//   name: details
//   namespace: bookstore
// spec:
//   labels:
//     app.kubernetes.io/name: details-legacy
//   network: remote
//   ports:
//     grpc: 3550
//     http: 8080
//   serviceAccount: default
// ```
// {{</tab>}}
// {{</tabset>}}
//

package v1alpha3

import (
	bytes "bytes"
	fmt "fmt"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "istio.io/gogo-genproto/googleapis/google/api"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for WorkloadGroup
func (this *WorkloadGroup) MarshalJSON() ([]byte, error) {
	str, err := WorkloadGroupMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for WorkloadGroup
func (this *WorkloadGroup) UnmarshalJSON(b []byte) error {
	return WorkloadGroupUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	WorkloadGroupMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	WorkloadGroupUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)
