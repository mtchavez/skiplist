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

## TODO

* Update to use `interface{}` for key/value
  * With a compare interface
* Concurrent skiplist implementation
