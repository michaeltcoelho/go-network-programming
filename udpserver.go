/* UDP daytime server
 */

package main

import (
	"fmt"
	"os"
	"net"
	"time"
)

func main() {
	service := ":1200"

	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	checkError(err)

	conn, err := net.ListenUDP("udp", udpAddr)
	checkError(err)

	for {
		fmt.Println("Listining on UDP...")
		handleClient(conn)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s ", err.Error())
		os.Exit(1)
	}
}

func handleClient(conn *net.UDPConn) {
	var buf [512]byte
	
	_, addr, err := conn.ReadFromUDP(buf[0:])

	fmt.Println("Reading from ", addr.String())

	if err != nil {
		return
	}

	daytime := time.Now().String()

	conn.WriteToUDP([]byte(daytime), addr)
}