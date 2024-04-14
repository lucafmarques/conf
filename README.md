# `conf`: application configuration made simple 
[![Go Reference](https://pkg.go.dev/badge/github.com/lucafmarques/conf.svg)](https://pkg.go.dev/github.com/lucafmarques/conf)
[![Go Report Card](https://goreportcard.com/badge/github.com/lucafmarques/conf)](https://goreportcard.com/report/github.com/lucafmarques/conf)

`conf` makes application configuration simple by providing a small, [opinionated](#opinions) API for building configuration structs using environment variables.

```
go get github.com/lucafmarques/conf
```

`conf` is a simple struct tag wrapper for [`env`](https://github.com/lucafmarques/env) and as such, supports all types `env` does.

## Example

```go
package main

import "github.com/lucafmarques/conf"

type config struct {
	URIs       []string  `env:"CONFIG_URIS"`
	MaxSize    uint16    `env:"CONFIG_MAX_SIZE"`
	LogLevel   string    `env:"CONFIG_LOG_LEVEL"`
	TraceRatio float32   `env:"CONFIG_TRACE_RATIO"`
	CutoffDate time.Time `env:"CONFIG_CUTOFF_DATE"`
}

func main() {
  cfg := &config{}
  err := conf.Build(cfg)
  // ...
}
```
---
### Opinions 

`conf` was designed to be used as-is, without any knobs to tweak. Because of that design decision, `conf` assumes that:

- A field of type `[]T` must have its ENV value be a comma-separated list of items
- A field of type `struct` must either:
  - have a struct tag itself and implement `encoding.TextUnmarshaler`
  - have a struct tag in its fields
