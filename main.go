package main

import (
	"bytes"
	"net"
	"os"
	"time"

	log "github.com/cihub/seelog"
)

const (
	UDP = "udp"
)

var (
	packetSize  = 512
	bindAddress = "0.0.0.0:8130"
)

func main() {
	log.Infof("Phosphor started at %v", time.Now())

	// @todo parse flags

	if err := listen(); err != nil {
		os.Exit(1)
	}

}

// listen on a UDP socket for trace frames
func listen() error {

	// Resolve bind address
	address, err := net.ResolveUDPAddr(UDP, bindAddress)
	if err != nil {
		log.Errorf("Failed to resolve address: %s", err.Error())
		return err
	}

	// Take the resolved address and attempt to listen on the UDP socket
	listener, err := net.ListenUDP(UDP, address)
	if err != nil {
		log.Errorf("ListenUDP error: %s", err.Error())
		return err
	}
	defer listener.Close()

	// Make a channel
	ch := make(chan []byte)

	// Listening loop
	log.Infof("Listening on %s for UDP trace frames", address.String())
	for {
		message := make([]byte, packetSize)
		n, remaddr, error := listener.ReadFrom(message)
		if error != nil {
			continue
		}
		buf := bytes.NewBuffer(message[0:n])
		log.Infof("Packet received from %s: %s", remaddr, string(message[0:n]))

		// Attempt to push into our channel to be processed by a worker
		select {
		case ch <- buf.Bytes():
		default:
			// abort!
			log.Infof("Dropped message")
		}
	}
}
