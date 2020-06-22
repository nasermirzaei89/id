# ID
Unique String Generator

[![Build Status](https://travis-ci.org/nasermirzaei89/id.svg?branch=master)](https://travis-ci.org/nasermirzaei89/id)
[![Go Report Card](https://goreportcard.com/badge/github.com/nasermirzaei89/id)](https://goreportcard.com/report/github.com/nasermirzaei89/id)
[![GoDoc](https://godoc.org/github.com/nasermirzaei89/id?status.svg)](https://godoc.org/github.com/nasermirzaei89/id)
[![GitHub license](https://img.shields.io/github/license/nasermirzaei89/id.svg)](https://github.com/nasermirzaei89/id/blob/master/LICENSE)

## Description

This package generates a new unique string (40 Character) that is useful for **id**.

## Install

### As Library

```sh
go get github.com/nasermirzaei89/id
```

### As Binary

```sh
go get github.com/nasermirzaei89/id/...
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/nasermirzaei89/id"
)

func main() {
	fmt.Println(id.New())
}
```
