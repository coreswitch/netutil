// Copyright 2017 CoreSwitch
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package netutil

import (
	"fmt"
	"testing"
)

func TestPtreeMatch(t *testing.T) {
	ptree := NewPtree(32)
	net1, _ := ParsePrefix("10.0.0.0/25")
	n1 := ptree.Acquire(net1.IP, net1.Length)
	n1.Item = "10"

	net2, _ := ParsePrefix("11.0.0.0/30")
	n2 := ptree.Acquire(net2.IP, net2.Length)
	n2.Item = "11"

	net, _ := ParsePrefix("10.0.0.127/32")
	n := ptree.Match(net.IP, net.Length)
	if n != nil {
		fmt.Println("Found")
	} else {
		fmt.Println("Not Found")
		t.Errorf("test error")
	}

	net, _ = ParsePrefix("10.0.0.128/32")
	n = ptree.Match(net.IP, net.Length)
	if n != nil {
		fmt.Println("Found")
		t.Errorf("test error")
	} else {
		fmt.Println("Not Found")
	}

	net, _ = ParsePrefix("11.0.0.1/32")
	n = ptree.Match(net.IP, net.Length)
	if n != nil {
		fmt.Println("Found")
	} else {
		fmt.Println("Not Found")
		t.Errorf("test error")
	}

	net, _ = ParsePrefix("11.0.0.4/32")
	n = ptree.Match(net.IP, net.Length)
	if n != nil {
		fmt.Println("Found")
		t.Errorf("test error")
	} else {
		fmt.Println("Not Found")
	}
}
