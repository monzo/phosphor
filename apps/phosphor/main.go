package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mondough/phosphor/internal/version"
	"github.com/mondough/phosphor/phosphor"
	"github.com/mreiferson/go-options"
)

func phosphorFlagset() *flag.FlagSet {
	flagSet := flag.NewFlagSet("phosphor", flag.ExitOnError)

	// basic options
	flagSet.Bool("version", false, "print version string")
	flagSet.Bool("verbose", false, "enable verbose logging")

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
	// p.Exit()
}
