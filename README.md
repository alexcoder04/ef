# ef

**E**asy **F**iles - a library that makes working with files in Go enjoyable.

## Install

```sh
go get github.com/alexcoder04/ef
```

## Use

```go
package main

import (
    ...

    "github.com/alexcoder04/ef
)

func main(){
    f := ef.NewFile("assets", "image.png")
    if !f.Exists {
        println("file does not exist")
        return
    }

    println(f.PathAbs())

    err := f.Copy("new/image.png")
    if err != nil {
        ...
    }
}
```

## Documentation

Please refer to [pkg.go.dev](https://pkg.go.dev/alexcoder04/ef).
