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

package random

import (
	"math/rand"
	"time"
)

import (
	"dubbo.apache.org/dubbo-go/v3/common"
	"dubbo.apache.org/dubbo-go/v3/common/extension"
	"dubbo.apache.org/dubbo-go/v3/registry"
	"dubbo.apache.org/dubbo-go/v3/registry/servicediscovery/instance"
)

func init() {
	rand.Seed(time.Now().UnixNano())
	extension.SetServiceInstanceSelector("random", NewRandomServiceInstanceSelector)
}

// the ServiceInstanceSelector implementation based on Random algorithm
type RandomServiceInstanceSelector struct{}

func NewRandomServiceInstanceSelector() instance.ServiceInstanceSelector {
	return &RandomServiceInstanceSelector{}
}

func (r *RandomServiceInstanceSelector) Select(url *common.URL, serviceInstances []registry.ServiceInstance) registry.ServiceInstance {
	if len(serviceInstances) == 0 {
		return nil
	}
	if len(serviceInstances) == 1 {
		return serviceInstances[0]
	}
	index := rand.Intn(len(serviceInstances))
	return serviceInstances[index]
}
