# bthesisAssistant
Custom assistant for Thesis

```shell
GOOS=windows go build -ldflags -H=windowsgui && mv bthesisAssistant.exe ~/Share
go build -v -ldflags="-X 'main.Version=0.1-linux'" ./cmd/bthesisAssistant/
```

