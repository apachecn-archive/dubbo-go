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

package reader

import (
	"bytes"
	"testing"
)

import (
	"github.com/dubbogo/gost/encoding/yaml"

	"github.com/stretchr/testify/assert"
)

import (
	"dubbo.apache.org/dubbo-go/v3/protocol/rest/config"
)

func TestRestConfigReaderReadConsumerConfig(t *testing.T) {
	bs, err := yaml.LoadYMLConfig("./testdata/consumer_config.yml")
	assert.NoError(t, err)
	configReader := NewRestConfigReader()
	err = configReader.ReadConsumerConfig(bytes.NewBuffer(bs))
	assert.NoError(t, err)
	assert.NotEmpty(t, config.GetRestConsumerServiceConfigMap())
}

func TestRestConfigReaderReadProviderConfig(t *testing.T) {
	bs, err := yaml.LoadYMLConfig("./testdata/provider_config.yml")
	assert.NoError(t, err)
	configReader := NewRestConfigReader()
	err = configReader.ReadProviderConfig(bytes.NewBuffer(bs))
	assert.NoError(t, err)
	assert.NotEmpty(t, config.GetRestProviderServiceConfigMap())
}
