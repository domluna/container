lru
----

### Installation

```sh
$ go get github.com/domluna/container/lru
```

### Usage

```go
package main

import (
  "github.com/domluna/container/lru"
)

func main() {
  c := lru.New(10) // lru of size 10
  c.Add("Hello", 1)
  c.Add("blah", 2)

  v, ok := c.Get("blah") // 2, true
  v, ok := c.Get("foo") // nil, false
}

```
