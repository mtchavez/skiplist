# Skiplist

[![Latest Version](http://img.shields.io/github/release/mtchavez/skiplist.svg?style=flat-square)](https://github.com/mtchavez/skiplist/releases)
[![Build Status](https://travis-ci.org/mtchavez/skiplist.svg?branch=master)](https://travis-ci.org/mtchavez/skiplist)
[![Go Documentation](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/mtchavez/skiplist)
[![Go Report Card](https://goreportcard.com/badge/github.com/mtchavez/skiplist)](https://goreportcard.com/report/github.com/mtchavez/skiplist)
[![Maintainability](https://api.codeclimate.com/v1/badges/e7513b7306bbabb3d2b4/maintainability)](https://codeclimate.com/github/mtchavez/skiplist/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/e7513b7306bbabb3d2b4/test_coverage)](https://codeclimate.com/github/mtchavez/skiplist/test_coverage)

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
# Benchmarked on 2017-09-16
goos: darwin
goarch: amd64
pkg: github.com/mtchavez/skiplist
BenchmarkInsert_1000-8              2000            805488 ns/op
BenchmarkInsert_10000-8              200           8370616 ns/op
BenchmarkInsert_100000-8              20          98251825 ns/op
BenchmarkInsert_1000000-8              1        1122310227 ns/op
BenchmarkParallelInsert-8        1000000              1349 ns/op
BenchmarkDelete_1000-8              5000            221056 ns/op
BenchmarkDelete_10000-8              500           3577634 ns/op
BenchmarkDelete_100000-8              30          61547826 ns/op
BenchmarkDelete_1000000-8              2         611290978 ns/op
BenchmarkParallelDelete-8        2000000               802 ns/op
```

## TODO

* Implement a compare interface
* Concurrent skiplist implementation
