# Skiplist

Skiplist implementation in Go. Read more on [Skip Lists](http://en.wikipedia.org/wiki/Skip_list)

This implementation is sligtly different in that it allows duplicate
keys for nodes.

## Install

`go get -u github.com/mtchavez/skiplist`

## Usage

Initialize a skiplist

```go
package main

func main() {
    list := skiplist.New()
}
```

Insert Nodes

```go
package main

func main() {
    list := skiplist.New()
    list.Insert(1, "Node 1")
    list.Insert(2, "Node 2")
}
```

Iterate through nodes

```go
package main

func main() {
    list := skiplist.New()
    list.Insert(1, "Node 1")
    list.Insert(2, "Node 2")
    list.Insert(3, "Node 3")

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

* Implement delete node
* Implement generic node key/value
* Benchmarks

## License

Written by Chavez

Released under the MIT License: http://www.opensource.org/licenses/mit-license.php
