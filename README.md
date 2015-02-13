# PhosphorD

PhosphorD is a local forwarder, which receives traces from the Phosphor client, and forwards to the [Phosphor server](https://github.com/mattheath/phosphor).

## Usage

`go get github.com/mattheath/phosphord`

### Command line options

```
  -buffer-size=200: set the maximum number of traces buffered per worker before batch sending
  -nsq-topic="trace": nsq topic to forward traces to
  -nsqd-tcp-address=: nsqd TCP address (may be given multiple times)
  -num-forwarders=20: set the number of workers which buffer and forward traces
  -verbose=false: enable verbose logging
```
