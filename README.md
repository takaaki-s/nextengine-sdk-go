# NextEngine SDK for Go

nextengine-sdk-go is the NextEngine SDK for the Go programming language.

# Getting Started

## Install

```shell
go get github.com/takaaki-s/nextengine-sdk-go
```

To update the SDK

```shell
go get -u github.com/takaaki-s/nextengine-sdk-go
```

# Example

## Quick Example

```go
package main

import (
	"context"
	"fmt"

	"github.com/takaaki-s/nextengine-sdk-go"
)

func main() {
	nc := nextengine.NewDefaultClient(
		"<CLIENT_ID>",
		"<CLIENT_SECRET>",
		"<REDIRECT_URI>",
		"<ACCESS_TOKEN>",
		"<REFRESH_TOKEN>")
	ctx := context.Background()
	ord := &entity.ReceiveOrderBase{}
	err := nc.APIExecute(ctx, "/api_v1_receiveorder_base/count", map[string]string{"receive_order_id-gte": "1"}, ord)
	if err != nil {
		fmt.Printf("error: %#v", err)
		return
	}
	fmt.Printf("response: %v", ord)
}

```

## NextEngine API Reference

Please see to https://developer.next-engine.com/

# Lisence

MIT