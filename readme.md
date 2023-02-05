### To build and compress app correct to linux and macOS

```
# Remember to build your handler executable for Linux!
# Build the project
GOOS=linux GOARCH=amd64 go build -o main main.go
# Zip the executable
zip main.zip main
```
