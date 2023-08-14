package main

import (
	"fmt"
	"net"
)

const (
	TargetServerAddress = "127.0.0.1"
	TargetServerPort    = 19132
)

func main() {
	socket, err := net.Dial("udp", fmt.Sprintf("%s:%d", TargetServerAddress, TargetServerPort))
	if err != nil {
		fmt.Println("Error creating socket:", err)
		return
	}
	defer socket.Close()

	protocol := byte(10) // Currently 11. But you shouldn't do 11 if you want an error in the console.
	mtu := byte(46)      // The MTU sent in the response appears to be somewhere around the size of this padding + 46 (28 udp overhead, 1 packet id, 16 magic, 1 protocol version). This padding seems to be used to discover the maximum packet size the network can handle.

	request := []byte{0x05}
	request = append(request, []byte{0x00, 0xff, 0xff, 0x00, 0xfe, 0xfe, 0xfe, 0xfe, 0xfd, 0xfd, 0xfd, 0xfd, 0x12, 0x34, 0x56, 0x78}...) // Always those hex bytes, corresponding to RakNet's default OFFLINE_MESSAGE_DATA_ID
	request = append(request, protocol)
	request = append(request, mtu)

	_, err = socket.Write(request)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}

	fmt.Println("\nSending Open Connection Request 1...")
	response := make([]byte, 512)
	_, err = socket.Read(response)
	if err != nil {
		fmt.Println("Error receiving response:", err)
		return
	}

	fmt.Println("\nReceived Open Connection Reply 1.\n\n")
	fmt.Printf("%v\n", response)
}
