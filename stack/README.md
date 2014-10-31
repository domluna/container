stack
====

View the [docs](https://godoc.org/github.com/domluna/container/stack)

### Installation

```sh
$ go get github.com/domluna/container/stack
```

### Example

```go

import "github.com/domluna/container/stack"

func main() {
  s := stack.New()
  s.Push(22)
  s.Push(55)
  v := t.Pop() // v = 55
  v := t.Pop() // v = 22
  v := t.Get() // v = nil
}

```

