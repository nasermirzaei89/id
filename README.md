# ID
Unique String Generator

## Description

This package generates a new unique string (40 Character) that is useful for **id**.

## Install

### As Library

```sh
go get github.com/nasermirzaei89/id
```

### As Binary

```sh
go get github.com/nasermirzaei89/id/cmd/id
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/nasermirzaei89/id"
)

func main() {
	fmt.Print(id.New())
}
```
