package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mondough/phosphor/internal/util"
	"github.com/mondough/phosphor/internal/version"
	"github.com/mondough/phosphor/phosphor"
	"github.com/mreiferson/go-options"
)

func phosphorFlagset() *flag.FlagSet {
	flagSet := flag.NewFlagSet("phosphor", flag.ExitOnError)

	// basic options
	flagSet.Bool("version", false, "print version string")
	flagSet.Bool("verbose", false, "enable verbose logging")
	flagSet.Int64("worker-id", 0, "unique seed for message ID generation (int) in range [0,4096) (will default to a hash of hostname)")
	flagSet.String("https-address", "", "<addr>:<port> to listen on for HTTPS clients")
	flagSet.String("http-address", "0.0.0.0:7750", "<addr>:<port> to listen on for HTTP clients")

	// NSQ Transport options
	nsqLookupdHTTPAddrs := util.StringArray{}
	flagSet.Var(&nsqLookupdHTTPAddrs, "nsqlookupd-http-address", "nsqlookupd HTTP address (may be given multiple times)")
	nsqdHTTPAddrs := util.StringArray{}
	flagSet.Var(&nsqdHTTPAddrs, "nsqd-http-address", "nsqd HTTP address (may be given multiple times)")
	flagSet.String("nsq-topic", "phosphor", "NSQ topic name to recieve traces from")
	flagSet.String("nsq-channel", "phosphor-server", "NSQ channel name to recieve traces from. This should be the same for all instances of the phosphor servers to spread ingestion work.")
	flagSet.Int("nsq-max-inflight", 200, "Number of traces to allow NSQ to keep inflight")
	flagSet.Int("nsq-num-handlers", 10, "Number of concurrent NSQ handlers to run")

	return flagSet
}

func main() {
	flagSet := phosphorFlagset()
	flagSet.Parse(os.Args[1:])

	// Globally seed rand
	rand.Seed(time.Now().UTC().UnixNano())

	if flagSet.Lookup("version").Value.(flag.Getter).Get().(bool) {
		fmt.Println(version.String("phosphor"))
		return
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	opts := phosphor.NewOptions()
	cfg := map[string]interface{}{}
	options.Resolve(opts, flagSet, cfg)

	p := phosphor.New(opts)

	p.Run()
	<-signalChan
	p.Exit()
}
