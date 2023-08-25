<h1 align="center">
    <img src="resources/logo.svg" alt="go-vee Logo" width="336px" /><br />
    go-vee - Govee API Wrapper for Go
</h1>

go-vee is an unofficial Go wrapper for the Govee API. It makes it easy to control your Govee smart devices from your Go code. It is based on the [official Govee API documentation](https://govee-public.s3.amazonaws.com/developer-docs/GoveeDeveloperAPIReference.pdf). **This project is a work in progress, use in production at your own risk.**

> To get started with the Govee API, you need to request an API key. You can do this from the Govee Home app. Open the app, go to Profile > Settings > Apply for API Key.

## Installation

```bash
go get github.com/loxhill/go-vee
```

## Usage

```go
package main

import (
    "fmt"
    govee "github.com/loxhill/go-vee"
)

client := govee.New("your-api-key")

listRequest := client.ListDevices()
response, err := client.Run(listRequest)
if err != nil {
    panic(err)
}
devices := response.Devices()

controlRequest, _ := client.Device(devices[0].Device, devices[0].Model).TurnOn()
_, err = client.Run(controlRequest)
if err != nil {
    panic(err)
}
```