# PhosphorD

PhosphorD is a local forwarder, like StatsD, which receives traces from the Phosphor client, and forwards to the [Phosphor server](https://github.com/mondough/phosphor/tree/master/phosphor).

Currently this receives Traces over UDP, which prevents clients blocking, but is reasonably reliable on a local machine. In the event this blocks, traces will be dropped and lost.

A future improvement would make this configurable to read from local files, mirroring the behaviour of Dapper Daemons as described in the [Google Dapper](https://research.google.com/pubs/pub36356.html) paper.

## Usage

```
  -buffer-size int
    	set the maximum number of traces buffered per worker before batch sending (default 200)
  -flush-interval int
    	set the maximum flush interval in ms (default 2000)
  -nsq-topic string
    	NSQ topic name to recieve traces from (default "phosphor")
  -nsqd-tcp-address value
    	nsqd TCP address (may be given multiple times)
  -num-forwarders int
    	set the number of workers which buffer and forward traces (default 20)
  -udp-address string
    	<addr>:<port> to listen for UDP traces (default "0.0.0.0:7760")
  -verbose
    	enable verbose logging
  -version
    	print version string
```
