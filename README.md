# FilMeta endec 

[![Lint Code Base](https://github.com/bitrainforest/filmeta-endec/actions/workflows/linter.yml/badge.svg)](https://github.com/bitrainforest/filmeta-endec/actions/workflows/linter.yml)

FilMeta type encoder/decoder for mongodb driver.

## Install

```shell
$ go get -u github.com/bitrainforest/filmeta-endec
```

## Usage:

```go
package main

import (
	endec "github.com/bitrainforest/filmeta-endec"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	url := "mongodb://localhost:27017"
	opt := options.Client().ApplyURI(url)

	// Build and SetRegistry
	rb := endec.BuildDefaultRegistry()
	opt.SetRegistry(rb)

	client, err := mongo.NewClient(opt)
	
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(nil)
}
```