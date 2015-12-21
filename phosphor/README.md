# Phosphor

The Phosphor server receives traces from PhosphorD via NSQ and stores these for later retrieval via its API.

## Usage

```
-http-address string
    <addr>:<port> to listen on for HTTP clients (default "0.0.0.0:7750")
-https-address string
    <addr>:<port> to listen on for HTTPS clients
-nsq-channel string
    NSQ channel name to recieve traces from. This should be the same for all instances of the phosphor servers to spread ingestion work. (default "phosphor-server")
-nsq-max-inflight int
    Number of traces to allow NSQ to keep inflight (default 200)
-nsq-num-handlers int
    Number of concurrent NSQ handlers to run (default 10)
-nsq-topic string
    NSQ topic name to recieve traces from (default "phosphor")
-nsqd-http-address value
    nsqd HTTP address (may be given multiple times)
-nsqlookupd-http-address value
    nsqlookupd HTTP address (may be given multiple times)
-verbose
    enable verbose logging
-version
    print version string
-worker-id int
    unique seed for message ID generation (int) in range [0,4096) (will default to a hash of hostname)
```
