# Skiplist

[![Latest Version](http://img.shields.io/github/release/mtchavez/skiplist.svg?style=flat-square)](https://github.com/mtchavez/skiplist/releases)
[![Build Status](https://travis-ci.org/mtchavez/skiplist.svg)](https://travis-ci.org/mtchavez/skiplist)
[![Go Documentation](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/mtchavez/skiplist)
[![Coverage Status](https://coveralls.io/repos/github/mtchavez/skiplist/badge.svg?branch=master)](https://coveralls.io/github/mtchavez/skiplist?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/mtchavez/skiplist)](https://goreportcard.com/report/github.com/mtchavez/skiplist)

Skiplist implementation in Go. Read more on [Skip Lists](http://en.wikipedia.org/wiki/Skip_list)

Skiplist with various additions and potential variations on the standard list.

## Install

`go get -u github.com/mtchavez/skiplist`

## Usage

Initialize a skiplist

```go
package main

func main() {
    list := skiplist.NewList()
}
```

Insert Nodes

```go
package main

func main() {
    list := skiplist.NewList()
    list.Insert(1, []byte("Node 1"))
    list.Insert(2, []byte("Node 2"))
}
```

Iterate through nodes

```go
package main

import (
    "fmt"
)

func main() {
    list := skiplist.NewList()
    list.Insert(1, []byte("Node 1"))
    list.Insert(2, []byte("Node 2"))
    list.Insert(3, []byte("Node 3"))

    for i := list.Iterator(); i.Next(); {
        // Print Key
        fmt.Println(i.Key())

        // Print Value
        fmt.Println(i.Val())
    }
}
```

## Documentation

Docs are on [GoDoc](http://godoc.org/github.com/mtchavez/skiplist)

## Tests

Run tests with coverage

`go test --cover`

## Benchmarks

Benchmarked with on a 2.3 GHz Intel Core i7.

```
goos: darwin
goarch: amd64
pkg: github.com/mtchavez/skiplist
BenchmarkInsert_1000-8              2000            850617 ns/op
BenchmarkInsert_10000-8              200           9111489 ns/op
BenchmarkInsert_100000-8              10         112362795 ns/op
BenchmarkInsert_1000000-8              1        2612950317 ns/op
BenchmarkParallelInsert-8        1000000              2935 ns/op
BenchmarkDelete_1000-8              5000            216831 ns/op
BenchmarkDelete_10000-8              500           3288757 ns/op
BenchmarkDelete_100000-8              30          44254999 ns/op
BenchmarkDelete_1000000-8              3         432647827 ns/op
BenchmarkParallelDelete-8        2000000               715 ns/op
```

## TODO

* Update to use `interface{}` for key/value
  * With a compare interface
* Concurrent skiplist implementation
