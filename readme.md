### To build and compress app correct to linux and macOS

```
# Remember to build your handler executable for Linux!
GOOS=linux GOARCH=amd64 go build -o main main.go
zip main.zip main
```