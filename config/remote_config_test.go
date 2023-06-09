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

package config

import (
	"testing"
	"time"
)

import (
	"github.com/stretchr/testify/assert"
)

import (
	"dubbo.apache.org/dubbo-go/v3/common/constant"
)

func TestRemoteConfig_GetParam(t *testing.T) {
	rc := &RemoteConfig{
		Params: make(map[string]string, 1),
	}

	def := "default value"
	key := "key"
	value := rc.GetParam(key, def)
	assert.Equal(t, def, value)

	actualVal := "actual value"
	rc.Params[key] = actualVal

	value = rc.GetParam(key, def)
	assert.Equal(t, actualVal, value)
}

func TestNewRemoteConfigBuilder(t *testing.T) {
	config := NewRemoteConfigBuilder().
		SetProtocol("nacos").
		SetAddress("127.0.0.1:8848").
		SetTimeout("10s").
		SetUsername("nacos").
		SetPassword("nacos").
		SetParams(map[string]string{"timeout": "3s"}).
		AddParam("timeout", "15s").
		Build()

	values := config.getUrlMap()
	assert.Equal(t, values.Get("timeout"), "15s")

	url, err := config.ToURL()
	assert.NoError(t, err)
	assert.Equal(t, url.GetParam("timeout", "3s"), "15s")

	err = config.Init()
	assert.NoError(t, err)

	timeout := config.GetTimeout()
	assert.Equal(t, timeout, 10*time.Second)
	assert.Equal(t, config.Prefix(), constant.RemotePrefix)

	param := config.GetParam("timeout", "3s")
	assert.Equal(t, param, "15s")
}
