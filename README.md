# ID
Unique String Generator

[![Build Status](https://travis-ci.org/nasermirzaei89/id.svg?branch=master)](https://travis-ci.org/nasermirzaei89/id)

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
