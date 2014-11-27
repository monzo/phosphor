package main

import (
	"bytes"
	"net"
	"os"
	"time"

	log "github.com/cihub/seelog"

	"github.com/bankpossible/iamdev/phosphord/forwarder"
)

const (
	UDP = "udp"
)

var (
	packetSize  = 512
	bindAddress = "0.0.0.0:8130"

	numForwarders = 20
	bufferSize    = 200
)

func main() {
	log.Infof("Phosphor started at %v", time.Now())

	// @todo parse flags

	// Make a channel to pass around trace frames
	ch := make(chan []byte)

	// Fire up a number of forwarders to process inbound messages
	forwarder.Start(ch, numForwarders, bufferSize)

	// Bind and listen to UDP traffic
	if err := listen(ch); err != nil {
		os.Exit(1)
	}

}

// listen on a UDP socket for trace frames
func listen(ch chan []byte) error {

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

	// Listen loop
	log.Infof("Listening on %s for UDP trace frames", address.String())
	for {
		message := make([]byte, packetSize)
		n, _, error := listener.ReadFrom(message)
		if error != nil {
			continue
		}
		buf := bytes.NewBuffer(message[0:n])
		// log.Infof("Packet received from %s: %s", remaddr, string(message[0:n]))

		// Attempt to push into our channel to be processed by a worker
		select {
		case ch <- buf.Bytes():
			// log.Infof("Wrote message to channel")
		default:
			// abort!
			// log.Infof("Dropped message")
		}
	}
}
