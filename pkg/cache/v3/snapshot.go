// Copyright 2018 Envoyproxy Authors
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

package cache

import (
	"encoding/json"
	"errors"
	"fmt"
	cluster "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	corev3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	endpoint "github.com/envoyproxy/go-control-plane/envoy/config/endpoint/v3"
	listenerv3 "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
	routev3 "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	tlsv3 "github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/tls/v3"
	runtime "github.com/envoyproxy/go-control-plane/envoy/service/runtime/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/envoyproxy/go-control-plane/pkg/resource/v3"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

// Snapshot is an internally consistent snapshot of xDS resources.
// Consistency is important for the convergence as different resource types
// from the snapshot may be delivered to the proxy in arbitrary order.
type Snapshot struct {
	Resources [types.UnknownType]Resources

	// VersionMap holds the current hash map of all resources in the snapshot.
	// This field should remain nil until it is used, at which point should be
	// instantiated by calling ConstructVersionMap().
	// VersionMap is only to be used with delta xDS.
	VersionMap map[string]map[string]string

	data []byte
}

var _ ResourceSnapshot = &Snapshot{}

type ResourcesJson struct {
	Version         string   `json:"version"`
	Endpoints       [][]byte `json:"endpoints"`
	Clusters        [][]byte `json:"clusters"`
	Routes          [][]byte `json:"routes"`
	ScopedRoute     [][]byte `json:"scopedRoute"`
	VirtualHostT    [][]byte `json:"virtualHostT"`
	Listeners       [][]byte `json:"listeners"`
	Runtimes        [][]byte `json:"runtimes"`
	Secrets         [][]byte `json:"secrets"`
	ExtensionConfig [][]byte `json:"extensionConfig"`
}

// NewSnapshot creates a snapshot from response types and a version.
// The resources map is keyed off the type URL of a resource, followed by the slice of resource objects.
func NewSnapshot(version string, resources map[resource.Type][]types.Resource) (*Snapshot, error) {
	out := Snapshot{}

	for typ, resource := range resources {
		index := GetResponseType(typ)
		if index == types.UnknownType {
			return nil, errors.New("unknown resource type: " + typ)
		}
		out.Resources[index] = NewResources(version, resource)
	}
	d, err := marshalSnapshot(version, resources)
	if err != nil {
		return nil, err
	}
	out.data = d
	return &out, nil
}

func marshalSnapshot(version string, resources map[resource.Type][]types.Resource) ([]byte, error) {
	var snapShotJson ResourcesJson
	snapShotJson.Version = version
	for typ, resourceArray := range resources {
		var dataArray [][]byte
		for _, v := range resourceArray {
			data, err := MarshalResource(v)
			if err != nil {
				return nil, err
			}
			dataArray = append(dataArray, data)
		}
		switch typ {
		case resource.EndpointType:
			snapShotJson.Endpoints = dataArray
			break
		case resource.ClusterType:
			snapShotJson.Clusters = dataArray
			break
		case resource.RouteType:
			snapShotJson.Routes = dataArray
			break
		case resource.ScopedRouteType:
			snapShotJson.ScopedRoute = dataArray
			break
		case resource.VirtualHostType:
			snapShotJson.VirtualHostT = dataArray
			break
		case resource.ListenerType:
			snapShotJson.Listeners = dataArray
			break
		case resource.RuntimeType:
			snapShotJson.Runtimes = dataArray
			break
		case resource.SecretType:
			snapShotJson.Secrets = dataArray
			break
		case resource.ExtensionConfigType:
			snapShotJson.ExtensionConfig = dataArray
			break
		default:
			return nil, fmt.Errorf("unknown type, type: %v", typ)
		}
	}
	return json.Marshal(snapShotJson)
}

// NewSnapshotWithTTLs creates a snapshot of ResourceWithTTLs.
// The resources map is keyed off the type URL of a resource, followed by the slice of resource objects.
func NewSnapshotWithTTLs(version string, resources map[resource.Type][]types.ResourceWithTTL) (*Snapshot, error) {
	out := Snapshot{}

	for typ, resource := range resources {
		index := GetResponseType(typ)
		if index == types.UnknownType {
			return nil, errors.New("unknown resource type: " + typ)
		}

		out.Resources[index] = NewResourcesWithTTL(version, resource)
	}

	return &out, nil
}

func Unmarshal(data []byte) (*Snapshot, error) {
	var snapShotJson ResourcesJson
	err := json.Unmarshal(data, &snapShotJson)
	if err != nil {
		return nil, err
	}
	resources := make(map[resource.Type][]types.Resource)
	if len(snapShotJson.Endpoints) > 0 {
		var temp []types.Resource
		for _, v := range snapShotJson.Endpoints {
			unmarshalledData := &endpoint.ClusterLoadAssignment{}
			d := &anypb.Any{TypeUrl: resource.EndpointType, Value: v}
			err = anypb.UnmarshalTo(d, unmarshalledData, proto.UnmarshalOptions{})
			if err != nil {
				return nil, err
			}
			temp = append(temp, types.Resource(unmarshalledData))
		}
		resources[resource.EndpointType] = temp
	}
	if len(snapShotJson.Clusters) > 0 {
		var temp []types.Resource
		for _, v := range snapShotJson.Clusters {
			unmarshalledData := &cluster.Cluster{}
			d := &anypb.Any{TypeUrl: resource.ClusterType, Value: v}
			err = anypb.UnmarshalTo(d, unmarshalledData, proto.UnmarshalOptions{})
			if err != nil {
				return nil, err
			}
			temp = append(temp, types.Resource(unmarshalledData))
		}
		resources[resource.ClusterType] = temp
	}
	if len(snapShotJson.Runtimes) > 0 {
		var temp []types.Resource
		for _, v := range snapShotJson.Runtimes {
			unmarshalledData := &runtime.Runtime{}
			d := &anypb.Any{TypeUrl: resource.RuntimeType, Value: v}
			err = anypb.UnmarshalTo(d, unmarshalledData, proto.UnmarshalOptions{})
			if err != nil {
				return nil, err
			}
			temp = append(temp, types.Resource(unmarshalledData))
		}
		resources[resource.RuntimeType] = temp
	}
	if len(snapShotJson.Secrets) > 0 {
		var temp []types.Resource
		for _, v := range snapShotJson.Secrets {
			unmarshalledData := &tlsv3.Secret{}
			d := &anypb.Any{TypeUrl: resource.SecretType, Value: v}
			err = anypb.UnmarshalTo(d, unmarshalledData, proto.UnmarshalOptions{})
			if err != nil {
				return nil, err
			}
			temp = append(temp, types.Resource(unmarshalledData))
		}
		resources[resource.SecretType] = temp
	}
	if len(snapShotJson.VirtualHostT) > 0 {
		var temp []types.Resource
		for _, v := range snapShotJson.VirtualHostT {
			unmarshalledData := &routev3.VirtualHost{}
			d := &anypb.Any{TypeUrl: resource.VirtualHostType, Value: v}
			err = anypb.UnmarshalTo(d, unmarshalledData, proto.UnmarshalOptions{})
			if err != nil {
				return nil, err
			}
			temp = append(temp, types.Resource(unmarshalledData))
		}
		resources[resource.VirtualHostType] = temp
	}
	if len(snapShotJson.ExtensionConfig) > 0 {
		var temp []types.Resource
		for _, v := range snapShotJson.ExtensionConfig {
			unmarshalledData := &corev3.TypedExtensionConfig{}
			d := &anypb.Any{TypeUrl: resource.ExtensionConfigType, Value: v}
			err = anypb.UnmarshalTo(d, unmarshalledData, proto.UnmarshalOptions{})
			if err != nil {
				return nil, err
			}
			temp = append(temp, types.Resource(unmarshalledData))
		}
		resources[resource.ExtensionConfigType] = temp
	}
	if len(snapShotJson.Routes) > 0 {
		var temp []types.Resource
		for _, v := range snapShotJson.Routes {
			unmarshalledData := &routev3.RouteConfiguration{}
			d := &anypb.Any{TypeUrl: resource.RouteType, Value: v}
			err = anypb.UnmarshalTo(d, unmarshalledData, proto.UnmarshalOptions{})
			if err != nil {
				return nil, err
			}
			temp = append(temp, types.Resource(unmarshalledData))
		}
		resources[resource.RouteType] = temp
	}
	if len(snapShotJson.Listeners) > 0 {
		var temp []types.Resource
		for _, v := range snapShotJson.Listeners {
			unmarshalledData := &listenerv3.Listener{}
			d := &anypb.Any{TypeUrl: resource.ListenerType, Value: v}
			err = anypb.UnmarshalTo(d, unmarshalledData, proto.UnmarshalOptions{})
			if err != nil {
				return nil, err
			}
			temp = append(temp, types.Resource(unmarshalledData))
		}
		resources[resource.ListenerType] = temp
	}
	if len(snapShotJson.ScopedRoute) > 0 {
		var temp []types.Resource
		for _, v := range snapShotJson.ScopedRoute {
			unmarshalledData := &routev3.ScopedRouteConfiguration{}
			d := &anypb.Any{TypeUrl: resource.ScopedRouteType, Value: v}
			err = anypb.UnmarshalTo(d, unmarshalledData, proto.UnmarshalOptions{})
			if err != nil {
				return nil, err
			}
			temp = append(temp, types.Resource(unmarshalledData))
		}
		resources[resource.ScopedRouteType] = temp
	}

	return NewSnapshot(snapShotJson.Version, resources)
}

// Consistent check verifies that the dependent resources are exactly listed in the
// snapshot:
// - all EDS resources are listed by name in CDS resources
// - all SRDS/RDS resources are listed by name in LDS resources
// - all RDS resources are listed by name in SRDS resources
//
// Note that clusters and listeners are requested without name references, so
// Envoy will accept the snapshot list of clusters as-is even if it does not match
// all references found in xDS.
func (s *Snapshot) Consistent() error {
	if s == nil {
		return errors.New("nil snapshot")
	}

	referencedResources := GetAllResourceReferences(s.Resources)

	// Loop through each referenced resource.
	referencedResponseTypes := map[types.ResponseType]struct{}{
		types.Endpoint: {},
		types.Route:    {},
	}

	for idx, items := range s.Resources {

		// We only want to check resource types that are expected to be referenced by another resource type.
		// Basically, if the consistency relationship is modeled as a DAG, we only want
		// to check nodes that are expected to have edges pointing to it.
		responseType := types.ResponseType(idx)
		if _, ok := referencedResponseTypes[responseType]; ok {
			typeURL, err := GetResponseTypeURL(responseType)
			if err != nil {
				return err
			}
			referenceSet := referencedResources[typeURL]

			if len(referenceSet) != len(items.Items) {
				return fmt.Errorf("mismatched %q reference and resource lengths: len(%v) != %d",
					typeURL, referenceSet, len(items.Items))
			}

			// Check superset.
			if err := superset(referenceSet, items.Items); err != nil {
				return fmt.Errorf("inconsistent %q reference: %w", typeURL, err)
			}
		}
	}

	return nil
}

func (s *Snapshot) GetMarshaledData() []byte {
	return s.data
}

// GetResources selects snapshot resources by type, returning the map of resources.
func (s *Snapshot) GetResources(typeURL resource.Type) map[string]types.Resource {
	resources := s.GetResourcesAndTTL(typeURL)
	if resources == nil {
		return nil
	}

	withoutTTL := make(map[string]types.Resource, len(resources))

	for k, v := range resources {
		withoutTTL[k] = v.Resource
	}

	return withoutTTL
}

// GetResourcesAndTTL selects snapshot resources by type, returning the map of resources and the associated TTL.
func (s *Snapshot) GetResourcesAndTTL(typeURL resource.Type) map[string]types.ResourceWithTTL {
	if s == nil {
		return nil
	}
	typ := GetResponseType(typeURL)
	if typ == types.UnknownType {
		return nil
	}
	return s.Resources[typ].Items
}

// GetVersion returns the version for a resource type.
func (s *Snapshot) GetVersion(typeURL resource.Type) string {
	if s == nil {
		return ""
	}
	typ := GetResponseType(typeURL)
	if typ == types.UnknownType {
		return ""
	}
	return s.Resources[typ].Version
}

// GetVersionMap will return the internal version map of the currently applied snapshot.
func (s *Snapshot) GetVersionMap(typeURL string) map[string]string {
	return s.VersionMap[typeURL]
}

// ConstructVersionMap will construct a version map based on the current state of a snapshot
func (s *Snapshot) ConstructVersionMap() error {
	if s == nil {
		return fmt.Errorf("missing snapshot")
	}

	// The snapshot resources never change, so no need to ever rebuild.
	if s.VersionMap != nil {
		return nil
	}

	s.VersionMap = make(map[string]map[string]string)

	for i, resources := range s.Resources {
		typeURL, err := GetResponseTypeURL(types.ResponseType(i))
		if err != nil {
			return err
		}
		if _, ok := s.VersionMap[typeURL]; !ok {
			s.VersionMap[typeURL] = make(map[string]string)
		}

		for _, r := range resources.Items {
			// Hash our version in here and build the version map.
			marshaledResource, err := MarshalResource(r.Resource)
			if err != nil {
				return err
			}
			v := HashResource(marshaledResource)
			if v == "" {
				return fmt.Errorf("failed to build resource version: %w", err)
			}

			s.VersionMap[typeURL][GetResourceName(r.Resource)] = v
		}
	}

	return nil
}
