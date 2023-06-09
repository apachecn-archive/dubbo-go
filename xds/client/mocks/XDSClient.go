/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	structpb "google.golang.org/protobuf/types/known/structpb"
)

import (
	bootstrap "dubbo.apache.org/dubbo-go/v3/xds/client/bootstrap"
	load "dubbo.apache.org/dubbo-go/v3/xds/client/load"
	resource "dubbo.apache.org/dubbo-go/v3/xds/client/resource"
)

// XDSClient is an autogenerated mock type for the XDSClient type
type XDSClient struct {
	mock.Mock
}

// BootstrapConfig provides a mock function with given fields:
func (_m *XDSClient) BootstrapConfig() *bootstrap.Config {
	ret := _m.Called()

	var r0 *bootstrap.Config
	if rf, ok := ret.Get(0).(func() *bootstrap.Config); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bootstrap.Config)
		}
	}

	return r0
}

// Close provides a mock function with given fields:
func (_m *XDSClient) Close() {
	_m.Called()
}

// DumpCDS provides a mock function with given fields:
func (_m *XDSClient) DumpCDS() map[string]resource.UpdateWithMD {
	ret := _m.Called()

	var r0 map[string]resource.UpdateWithMD
	if rf, ok := ret.Get(0).(func() map[string]resource.UpdateWithMD); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]resource.UpdateWithMD)
		}
	}

	return r0
}

// DumpEDS provides a mock function with given fields:
func (_m *XDSClient) DumpEDS() map[string]resource.UpdateWithMD {
	ret := _m.Called()

	var r0 map[string]resource.UpdateWithMD
	if rf, ok := ret.Get(0).(func() map[string]resource.UpdateWithMD); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]resource.UpdateWithMD)
		}
	}

	return r0
}

// DumpLDS provides a mock function with given fields:
func (_m *XDSClient) DumpLDS() map[string]resource.UpdateWithMD {
	ret := _m.Called()

	var r0 map[string]resource.UpdateWithMD
	if rf, ok := ret.Get(0).(func() map[string]resource.UpdateWithMD); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]resource.UpdateWithMD)
		}
	}

	return r0
}

// DumpRDS provides a mock function with given fields:
func (_m *XDSClient) DumpRDS() map[string]resource.UpdateWithMD {
	ret := _m.Called()

	var r0 map[string]resource.UpdateWithMD
	if rf, ok := ret.Get(0).(func() map[string]resource.UpdateWithMD); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]resource.UpdateWithMD)
		}
	}

	return r0
}

// ReportLoad provides a mock function with given fields: server
func (_m *XDSClient) ReportLoad(server string) (*load.Store, func()) {
	ret := _m.Called(server)

	var r0 *load.Store
	if rf, ok := ret.Get(0).(func(string) *load.Store); ok {
		r0 = rf(server)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*load.Store)
		}
	}

	var r1 func()
	if rf, ok := ret.Get(1).(func(string) func()); ok {
		r1 = rf(server)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(func())
		}
	}

	return r0, r1
}

// SetMetadata provides a mock function with given fields: _a0
func (_m *XDSClient) SetMetadata(_a0 *structpb.Struct) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*structpb.Struct) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WatchCluster provides a mock function with given fields: _a0, _a1
func (_m *XDSClient) WatchCluster(_a0 string, _a1 func(resource.ClusterUpdate, error)) func() {
	ret := _m.Called(_a0, _a1)

	var r0 func()
	if rf, ok := ret.Get(0).(func(string, func(resource.ClusterUpdate, error)) func()); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(func())
		}
	}

	return r0
}

// WatchEndpoints provides a mock function with given fields: clusterName, edsCb
func (_m *XDSClient) WatchEndpoints(clusterName string, edsCb func(resource.EndpointsUpdate, error)) func() {
	ret := _m.Called(clusterName, edsCb)

	var r0 func()
	if rf, ok := ret.Get(0).(func(string, func(resource.EndpointsUpdate, error)) func()); ok {
		r0 = rf(clusterName, edsCb)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(func())
		}
	}

	return r0
}

// WatchListener provides a mock function with given fields: _a0, _a1
func (_m *XDSClient) WatchListener(_a0 string, _a1 func(resource.ListenerUpdate, error)) func() {
	ret := _m.Called(_a0, _a1)

	var r0 func()
	if rf, ok := ret.Get(0).(func(string, func(resource.ListenerUpdate, error)) func()); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(func())
		}
	}

	return r0
}

// WatchRouteConfig provides a mock function with given fields: _a0, _a1
func (_m *XDSClient) WatchRouteConfig(_a0 string, _a1 func(resource.RouteConfigUpdate, error)) func() {
	ret := _m.Called(_a0, _a1)

	var r0 func()
	if rf, ok := ret.Get(0).(func(string, func(resource.RouteConfigUpdate, error)) func()); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(func())
		}
	}

	return r0
}
