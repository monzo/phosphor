package main

import (
	"fmt"
	"net"
	"time"

	"github.com/mattheath/goprotobuf/proto"

	pb "github.com/mattheath/phosphor/proto"
)

const (
	MAX_PACKET_SIZE = 1500 - 8 - 20 // 8-byte UDP header, 20-byte IP header
)

func main() {

	// Make example trace frame
	t := &pb.TraceFrame{
		TraceId:     proto.String("aasldjaskjdlsakjdkasjdklasjdlasjdkljdas"),
		SpanId:      proto.String("8yf8sdg76sg897b98fbuys8b9s6rvs6ducghkfhi27tuw"),
		ParentId:    proto.String("97as8d7s9a7a7dv32hrkqehfkuh23hq8d7h4g7iygs7ih"),
		Type:        pb.FrameType_OUT.Enum(),
		Timestamp:   proto.Int64(time.Now().UnixNano()),
		Duration:    proto.Int64(1231312),
		Hostname:    proto.String("somehostname"),
		Origin:      proto.String("some.api"),
		Destination: proto.String("some.service"),
		Payload:     proto.String(`{"boop":123}`),
	}

	// Marshal to bytes

	b, err := proto.Marshal(t)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Encoded: %s\n", string(b))
	fmt.Printf("Encoded bytes: %v\n", b)

	// Send via UDP!

	// Get a conn
	c, err := net.DialTimeout("udp", "localhost:7760", time.Second)
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
