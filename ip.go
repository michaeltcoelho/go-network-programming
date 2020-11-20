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

	ip(arg)

	os.Exit(0)
}

func ip(ip string) {
	addr := net.ParseIP(ip)
	if addr == nil {
		fmt.Println("Invalid address")
	}

	mask := addr.DefaultMask()
	network := addr.Mask(mask)
	ones, bits := mask.Size()

	fmt.Println("Address is ", addr.String())
	fmt.Println("Default mask length is ", bits)
	fmt.Println("Leading ones count is ", ones)
	fmt.Println("Mask is (hex) ", mask.String())
	fmt.Println("Network is ", network.String())
}
