# DotEnv for Go

A lightweight .env loader and parser for Go.

## Installation

As a library

```bash
go get github.com/golamee/dotenv
```

As a bin command

```bash
# Go version > 1.17
go install github.com/golamee/dotenv

# Go version < 1.17
go get github.com/golamee/dotenv
```

## Usage

### Import environment file

Basic import from .env file

```go
package main

import (
    "github.com/golamee/dotenv"
)

func main() {
    err := dotenv.Load()
    if err != nil {
        fmt.Println(err)
    }
}
```

Import from another environment file

```go
package main

import (
    "github.com/golamee/dotenv"
)

func main() {
    err := dotenv.Load(".env.local")
    if err != nil {
        fmt.Println(err)
    }
}
```

Import multiple environment files

```go
package main

import {
    "github.com/golamee/dotenv"
}

func main() {
    err := dotenv.Load(".env.db", ".env.rabbitmq", ".env.redis")
    if err != nil {
        fmt.Println(err)
    }
}
```

### Get environment data

#### Simple get

Example environment file (.env)

```env
NAME=sample
```

Example main.go

```go
package main

import (
    "github.com/golamee/dotenv"
)

func main() {
    err := dotenv.Load()
    if err != nil {
        fmt.Println(err)
    }

    data, err := dotenv.Get("NAME")
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println("NAME =", data)
}
```

### Type-safe getters

```go
// get data type Int64
data, err := dotenv.GetInt64("NAME")
if err != nil {
    fmt.Println(err)
}

fmt.Println("NAME =", data)
```

```go
// get data type Int32
data, err := dotenv.GetInt32("NAME")
if err != nil {
    fmt.Println(err)
}

fmt.Println("NAME =", data)
```

```go
// get data type Float64
data, err := dotenv.GetFloat64("NAME")
if err != nil {
    fmt.Println(err)
}

fmt.Println("NAME =", data)
```

```go
// get data type Float32
data, err := dotenv.GetFloat32("NAME")
if err != nil {
    fmt.Println(err)
}

fmt.Println("NAME =", data)
```

```go
// get data type Bool
data, err := dotenv.GetBool("NAME")
if err != nil {
    fmt.Println(err)
}

fmt.Println("NAME =", data)
```

```go
// get data type Complex64
data, err := dotenv.GetComplex64("NAME")
if err != nil {
    fmt.Println(err)
}

fmt.Println("NAME =", data)
```
