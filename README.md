# matching-engine

A low-latency matching engine / order book written in Go, built incrementally
as a study of exchange internals and systems performance.

## What it does

Maintains a two-sided limit order book and matches incoming orders against
resting liquidity under **price-time priority** — best price first, then
first-in-first-out within a price level. Supports limit and market orders,
partial fills, and a market data feed of book and trade events.

## Design goals

- **Deterministic, single-threaded core.** Every command is processed in
  sequence-number order, so the engine is replayable and auditable — the same
  input always yields the same trades. Concurrency lives only at ingestion and
  publication, never in the match loop.
- **No allocation on the hot path.** Intrusive FIFO queues, pooled nodes, and
  reused buffers keep the critical path garbage-free.
- **Layered architecture.** `domain` holds pure types; `matching` is the
  concrete, interface-free hot path; `application` and `infrastructure` handle
  wiring and I/O at the edges.

## Architecture

    cmd/engine            Composition root (wiring, config)
    internal/domain       Pure types: Order, Trade, Side, OrderType, ...
    internal/matching     Hot path: price levels, order book, match engine
    internal/application  Command/event model, intake sequencer
    internal/marketdata   Market data feed / fan-out
    internal/infrastructure  Config, publishers, FIX gateway
    bench                 Load and latency benchmarks

## Requirements

Go 1.23+

## License

MIT
EOF
