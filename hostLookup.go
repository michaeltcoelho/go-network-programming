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

	hostLookup(arg)

	os.Exit(0)
}

func hostLookup(name string) {
	addrs, err := net.LookupHost(name)
	if err != nil {
		fmt.Println("Resolution error ", err.Error())
		os.Exit(2)
	}

	for _, s := range addrs {
		fmt.Println(s)
	}
}
