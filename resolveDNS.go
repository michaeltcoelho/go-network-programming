package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Invalid parameters\n")
		os.Exit(1)
	}
	arg := os.Args[1]

	resolveDNS(arg)

	os.Exit(0)
}

func resolveDNS(name string) {
	addr, err := net.ResolveIPAddr("ip", name)
	if err != nil {
		fmt.Println("Resolution error ", err.Error())
		os.Exit(1)
	}
	fmt.Println("Resolved address is ", addr.String())
}
