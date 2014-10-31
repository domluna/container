Queue
====

View the [docs](https://godoc.org/github.com/domluna/container/queue)

### Installation

```sh
$ go get github.com/domluna/container/queue
```

### Example

```go

import (
  "github.com/domluna/container/queue"
)

func main() {
  q := queue.New()
  q.Push(1)
  q.Push(2)
  q.Push(3)
  v := t.Pop() // v = 1
  v := t.Pop() // v = 2
  v := t.Pop() // v = 3
  v := t.Pop() // v = nil
}

```

