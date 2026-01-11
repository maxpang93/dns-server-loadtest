package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", ":10080")
	if err != nil {
		log.Fatal("Error resolving UDP address: ", err)
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatal("Error listening to UDP address: ", err)
	}
	defer conn.Close()

	buffer := make([]byte, 1024)
	for {
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Printf("Error read from UDP: %v\n", err)
			continue
		}
		data := make([]byte, n)
		copy(data, buffer[:n])
		go handleUDPConnection(conn, addr, data)

	}
}

func handleUDPConnection(conn *net.UDPConn, remoteAddr *net.UDPAddr, data []byte) {

	fmt.Printf("Read %v bytes from %s\n", len(data), remoteAddr)
	resp := fmt.Sprintf("Msg received: %v\n", string(data))
	fmt.Println(resp)

	_, err := conn.WriteToUDP([]byte(resp), remoteAddr)
	if err != nil {
		fmt.Printf("Error writing response to: %v\n", err)
	}
}
