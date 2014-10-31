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
  q.Push(22)
  v := t.Pop() // v = 22
  v := t.Pop() // v = nil
}

```

