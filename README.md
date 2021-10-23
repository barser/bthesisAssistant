# bthesisAssistant
Custom assistant for Thesis

```shell
GOOS=windows go build -ldflags -H=windowsgui && mv bthesisAssistant.exe ~/Share
go build -v -ldflags="-X 'main.Version=0.1-linux'" ./cmd/bthesisAssistant/
```

[![Go Report Card](https://goreportcard.com/badge/github.com/barser/bthesisAssistant)](https://goreportcard.com/report/github.com/barser/bthesisAssistant)
