# Go Advanced Training Camp Week 3 Homework

## Problem

Implement an HTTP server startup and shutdown based on `errgroup`, along with the registration and handling of Linux signal signals. Ensure that when one service exits, all services exit as well.

## Approach

1. Graceful shutdown

2. When and how to shut down the HTTP server

## TODO

1. [x] Refactor code

2. [x] Add unit tests

## References

`https://github.com/da440dil/go-workgroup/blob/master/template/signal/signal.go`

`https://draveness.me/golang/docs/part3-runtime/ch06-concurrency/golang-sync-primitives/#errgroup`

## Build

`go run github.com/webmin7761/go-school/homework/concurrency`

## Usage

- <http://localhost:8080/>
- <http://127.0.0.1:6060/debug/pprof/>
