![](docs/logo.png)

Phosphor is a Distributed Tracing system, similar to [Google's Dapper](research.google.com/pubs/pub36356.html),  [Twitter's Zipkin](twitter.github.io/zipkin), and [Hailo's Trace Service](https://speakerdeck.com/mattheath/observability-in-micro-service-architectures).

It is comprised of a few simple components:

 - Phosphor Client, used to send traces from applications
 - Phosphor Daemon, collects traces and forwards onto the main server
 - Phosphor Server, stores traces and aggregated trace information
 - Phosphor UI, view trace and debug information about your infrastructure

![Phosphor Architecture](docs/phosphor/outline.png)

## Dependencies

 - NSQ is currently used as the delivery transport for traces between PhosphorD and the Phosphor Server

## Caveats

This system is currently in development, and some components are not yet open source. In particular, the persistence layer in this repository is an in-memory mock, and is therefore not appropriate for production usage. Additional storage adaptors will be added in the near future.

## Thanks

Phosphor was based on a similar systems built internally at [Hailo](https://hailoapp.com), and developed by @mattheath with the support of [Starling](https://starlingbank.co.uk).
