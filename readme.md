# 🔗 gRPC URL Shortener (Go + gRPC + Gin)

A simple URL shortener built using **Golang**, **gRPC**, and **Gin**. It supports gRPC-based URL shortening and a Gin-based HTTP server for redirecting users via short codes.

![Go](https://img.shields.io/badge/Go-1.20+-00ADD8?style=for-the-badge&logo=go)
![gRPC](https://img.shields.io/badge/gRPC-Enabled-green?style=for-the-badge&logo=grpc)
![Gin](https://img.shields.io/badge/Gin-HTTP_Server-blueviolet?style=for-the-badge&logo=go)
---

## 📦 Folder Structure

```bash

📁 grpc-url-shortener
├── 📁 client
│   └── main.go                  # gRPC client code
├── 📁 server
│   └── main.go                  # gRPC server implementation
├── 📁 http_redirect
│   └── main.go                  # Gin-based HTTP redirect server
├── 📁 proto
│   ├── urlshortener.proto       # gRPC service definition
│   ├── urlshortener.pb.go       # Generated Protobuf code
│   └── urlshortener_grpc.pb.go  # Generated gRPC service code
├── 📄 go.mod                    # Go module definition
└── 📄 README.md                 # Project documentation

```
---

## 🚀 Features

- Shorten long URLs via gRPC
- Resolve shortened URLs back to original
- HTTP redirect support using Gin (e.g., `http://localhost:8080/x7Yz3A`)
- Simple in-memory storage (no database required)

---

## 🔧 Setup & Installation

### Clone the Repository

```bash
git clone https://github.com/ayushgupta-lang/gRPC-URL-Shortener.git
cd grpc-url-shortener
```
### Install Dependencies
```bash
go mod tidy
go get google.golang.org/grpc
go get google.golang.org/protobuf
go get github.com/gin-gonic/gin
```
### Install protoc Plugins 
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
export PATH="$PATH:$(go env GOPATH)/bin"
```

### Generate Go Code from .proto
```bash
protoc --go_out=. --go-grpc_out=. proto/urlshortener.proto
```
## Run the Project
### Start the gRPC Server
```bash
go run server/main.go
```
### Run the gRPC Client
```bash
go run client/main.go
```
#### Example output:
#### Short code: x7Yz3A
#### Original URL: https://example.com/my-long-url

### Start the Gin HTTP Redirect Server
```bash
go run http_redirect/main.go
```

#### Open in your browser:
#### http://localhost:8080/short/x7Yz3A
#### It redirects you to https://example.com/my-long-url.













