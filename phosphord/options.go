package phosphord

type Options struct {
	// basic options
	Verbose       bool   `flag:"verbose"`
	UDPAddress    string `flag:"udp-address"`
	NumForwarders int    `flag:"num-forwarders"`
	BufferSize    int    `flag:"buffer-size"`
	FlushInterval int    `flag:"flush-interval"`

	// NSQ Transport options
	NSQDTCPAddresses []string `flag:"nsqd-tcp-address"`
	NSQTopicName     string   `flag:"nsq-topic"`
	NSQMaxInflight   int
	NSQNumHandlers   int
}

func NewOptions() *Options {
	return &Options{
		UDPAddress:    "0.0.0.0:7760",
		NumForwarders: 20,
		BufferSize:    200,
		FlushInterval: 2000,

		NSQTopicName:   "phosphor",
		NSQMaxInflight: 200,
		NSQNumHandlers: 10,
	}
}
