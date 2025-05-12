/*
Copyright 2022 The Kruise Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"fmt"
	"github.com/openkruise/kruise-game/cloudprovider/utils"
)

func main() {
	testPorts := []int32{80, 443, 139, 445}
	for _, port := range testPorts {
		fmt.Printf("Port %d is safe: %v\n", port, utils.IsPortSafe(port))
	}
}