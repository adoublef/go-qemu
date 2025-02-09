// Copyright 2016 The go-qemu Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	hypervisor "github.com/adoublef/go-qemu/hypervisor"
)

var (
	network = flag.String("network", "unix", `Named network used to connect on. For Unix sockets -network=unix, for TCP connections: -network=tcp`)
	address = flag.String("address", "/var/run/libvirt/libvirt-sock", `Address of the hypervisor. This could be in the form of Unix or TCP sockets. For TCP connections: -address="host:16509"`)
	timeout = flag.Duration("timeout", 2*time.Second, "Connection timeout. Another valid value could be -timeout=500ms")
)

func main() {
	flag.Parse()

	fmt.Printf("\nConnecting to %s://%s\n", *network, *address)
	newConn := func() (net.Conn, error) {
		return net.DialTimeout(*network, *address, *timeout)
	}

	driver := hypervisor.NewRPCDriver(newConn)
	hv := hypervisor.New(driver)

	fmt.Printf("\n**********Domains**********\n")
	domains, err := hv.Domains()
	if err != nil {
		log.Fatalf("Unable to get domains from hypervisor: %v", err)
	}
	for _, dom := range domains {
		fmt.Printf("%s\n", dom.Name)
	}
	fmt.Printf("\n***************************\n")
}
