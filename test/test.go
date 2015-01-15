package main

import (
	pb "github.com/bankpossible/iamdev/shared/messages"
	"github.com/mattheath/goprotobuf/proto"

	"fmt"
	"net"
	"time"
)

const (
	MAX_PACKET_SIZE = 1500 - 8 - 20 // 8-byte UDP header, 20-byte IP header
)

func main() {

	// Make example trace frame
	t := &pb.TraceFrame{
		TraceId:         proto.String("aasldjaskjdlsakjdkasjdklasjdlasjdkljdas"),
		RequestId:       proto.String("8yf8sdg76sg897b98fbuys8b9s6rvs6ducghkfhi27tuw"),
		ParentRequestId: proto.String("8yf8sdg76sg897b98fbuys8b9s6rvs6ducghkfhi27tuw"),
		Type:            proto.String("CLIENT_IN"),
		Timestamp:       proto.String(""),
		Duration:        proto.Int64(1231312),
		Hostname:        proto.String("nest"),
		From:            proto.String("some.api"),
		To:              proto.String("some.service"),
		Payload:         proto.String(`{"boop":123}`),
	}

	// Marshal to bytes

	b, err := proto.Marshal(t)
	if err != nil {
		panic(err)
	}

	fmt.Println("Encoded: %s", string(b))
	fmt.Println("Encoded bytes: %v", b)

	// Send via UDP!

	// Get a conn
	c, err := net.DialTimeout("udp", "192.168.59.103:8130", time.Second)
	// c, err := net.DialTimeout("udp", "localhost:8130", time.Second)
	if err != nil {
		panic(err)
	}

	// Write into the connection
	var i int

	for j := 0; j < 20; j++ {
		for i = 0; i < 500; i++ {
			_, err := c.Write([]byte(b))
			if err != nil {
				panic(err)
			}
		}
		time.Sleep(500 * time.Millisecond)
	}

	fmt.Println("Sent", i, "messages")

	// fmt.Println("Sent %v bytes", n)
}
