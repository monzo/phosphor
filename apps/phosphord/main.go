package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	log "github.com/cihub/seelog"
	"github.com/mreiferson/go-options"

	"github.com/mondough/phosphor/internal/util"
	"github.com/mondough/phosphor/internal/version"
	"github.com/mondough/phosphor/phosphord"
)

func phosphordFlagSet() *flag.FlagSet {
	flagSet := flag.NewFlagSet("phosphord", flag.ExitOnError)

	// basic options
	flagSet.Bool("version", false, "print version string")
	flagSet.Bool("verbose", false, "enable verbose logging")

	// forwarder options
	flagSet.String("udp-address", "0.0.0.0:7760", "<addr>:<port> to listen for UDP traces")
	flagSet.Int("num-forwarders", 20, "set the number of workers which buffer and forward traces")
	flagSet.Int("buffer-size", 200, "set the maximum number of traces buffered per worker before batch sending")
	flagSet.Int("flush-interval", 2000, "set the maximum flush interval in ms")

	// NSQ Transport options
	flagSet.String("nsq-topic", "phosphor", "NSQ topic name to recieve traces from")
	nsqdTCPAddrs := util.StringArray{}
	flagSet.Var(&nsqdTCPAddrs, "nsqd-tcp-address", "nsqd TCP address (may be given multiple times)")

	return flagSet
}

func main() {
	flagSet := phosphordFlagSet()
	flagSet.Parse(os.Args[1:])

	defer log.Flush()

	// Globally seed rand
	rand.Seed(time.Now().UTC().UnixNano())

	// Use ALL the CPUs
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Immediately print and exit the version number
	if flagSet.Lookup("version").Value.(flag.Getter).Get().(bool) {
		fmt.Println(version.String("phosphord"))
		return
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	opts := phosphord.NewOptions()
	cfg := map[string]interface{}{}
	options.Resolve(opts, flagSet, cfg)

	p := phosphord.New(opts)

	p.Run()
	<-signalChan
	p.Exit()
}
