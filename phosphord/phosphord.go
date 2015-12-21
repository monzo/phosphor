package phosphord

import (
	"bytes"
	"net"
	"os"
	"runtime"
	"time"

	log "github.com/cihub/seelog"

	"github.com/mondough/phosphor/phosphord/transport"
)

const (
	UDP = "udp"
)

var (
	packetSize = 65536 - 8 - 20 // 8-byte UDP header, 20-byte IP header
)

type PhosphorD struct {
	opts *Options
	tr   transport.Transport

	traceChan chan []byte

	exitChan chan struct{}
}

func New(opts *Options) *PhosphorD {
	// Initialise our transport
	// TODO ensure this doesn't connect until we Run()
	tr, err := transport.NewNSQTransport(opts.NSQTopicName, opts.NSQDTCPAddresses)
	if err != nil {
		log.Criticalf(err.Error())
		os.Exit(1)
	}

	return &PhosphorD{
		opts:      opts,
		tr:        tr,
		traceChan: make(chan []byte),

		exitChan: make(chan struct{}),
	}
}

func (p *PhosphorD) Run() {
	log.Infof("PhosphorD started at %v using %v CPUs", time.Now(), runtime.NumCPU())

	// Fire up a number of forwarders to process inbound messages
	log.Infof("Starting %v forwarders with buffer size of %v", p.opts.NumForwarders, p.opts.BufferSize)
	for i := 0; i < p.opts.NumForwarders; i++ {
		go p.forward(i)
	}

	// Bind and listen to UDP traffic
	go p.listen()
}

// Exit and shut down
func (p *PhosphorD) Exit() {
	log.Infof("PhosphorD exiting")
	select {
	case <-p.exitChan: // check if already closed
	default:
		close(p.exitChan)
	}
}

// listen on a UDP socket for trace frames
func (p *PhosphorD) listen() {

	// Resolve bind address
	address, err := net.ResolveUDPAddr(UDP, p.opts.UDPAddress)
	if err != nil {
		log.Errorf("Failed to resolve address: %s", err.Error())
		return
	}

	// Take the resolved address and attempt to listen on the UDP socket
	listener, err := net.ListenUDP(UDP, address)
	if err != nil {
		log.Errorf("ListenUDP error: %s", err.Error())
		return
	}
	defer listener.Close()

	// Listen loop
	log.Infof("Listening on %s for UDP trace frames", address.String())
	for {
		message := make([]byte, packetSize)
		n, _, err := listener.ReadFrom(message)
		if err != nil {
			continue
		}
		buf := bytes.NewBuffer(message[0:n])
		// log.Infof("Packet received from %s: %s", remaddr, string(message[0:n]))

		// Attempt to push into our channel to be processed by a worker
		select {

		// Successfully write inbound message to queue
		case p.traceChan <- buf.Bytes():

		// Stop listening and shut down
		case <-p.exitChan:
			return

		// Drop message to prevent blocking
		default:
		}
	}
}
