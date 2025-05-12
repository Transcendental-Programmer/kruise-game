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

package utils

import (
    "testing"
)

func TestIsPortSafe(t *testing.T) {
    safePorts := []int32{80, 443, 8080, 9090, 1000, 2000, 3000}
    for _, port := range safePorts {
        if !IsPortSafe(port) {
            t.Errorf("Port %d should be safe but was marked as unsafe", port)
        }
    }


    for _, port := range HighRiskPorts {
        if IsPortSafe(port) {
            t.Errorf("Port %d should be unsafe but was marked as safe", port)
        }
    }
}
