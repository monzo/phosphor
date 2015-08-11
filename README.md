# PhosphorD

PhosphorD is a local forwarder, like StatsD, which receives traces from the Phosphor client, and forwards to the [Phosphor server](https://github.com/mondough/phosphor).

Currently this receives Traces over UDP, which prevents clients blocking, but is reasonably reliable on a local machine. In the event this blocks, traces will be dropped and lost.

A future improvement would make this configurable to read from local files, mirroring the behaviour of Dapper Daemons as described in the [Google Dapper](https://research.google.com/pubs/pub36356.html) paper.

## Usage

`go get github.com/mondough/phosphord`

### Command line options

```
  -buffer-size=200: set the maximum number of traces buffered per worker before batch sending
  -nsq-topic="trace": nsq topic to forward traces to
  -nsqd-tcp-address=: nsqd TCP address (may be given multiple times)
  -num-forwarders=20: set the number of workers which buffer and forward traces
  -verbose=false: enable verbose logging
```
