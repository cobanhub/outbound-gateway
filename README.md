# Madakaripura

[![Go Version](https://img.shields.io/badge/go-1.21+-brightgreen.svg)](https://golang.org/)
[![License](https://img.shields.io/github/license/cobanhub/outbound-gateway)](./LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/cobanhub/outbound-gateway)](https://goreportcard.com/report/github.com/cobanhub/outbound-gateway)
[![Project Status](https://img.shields.io/badge/status-active--development-yellow.svg)]()


A lightweight and configurable middleware service that acts as a dynamic outbound gateway between core systems and third-party partner APIs.

## 🚀 Overview

The **Madakaripura** decouples your internal core system from external partner APIs by handling:

- Request transformation (core format → partner-specific)
- Forwarding logic with configurable endpoints and headers
- Response transformation (partner → core format)
- Retry, timeout, and logging

This pattern is ideal for B2B integrations where each partner has their own API structure, headers, and behavior.

---

## 🚧 Project Status: Ongoing Development

This project is currently under active development.  
Features may change, and contributions or feedback are very welcome.

Check out the progress on this board: \
[![Trello Board](https://img.shields.io/badge/Trello-Board-blue?logo=trello&style=for-the-badge)](https://trello.com/b/1wuLhpMq/madakaripura)


## 🧱 Architecture

```text
+-------------+          +------------------+          +-------------------+
| Core System |   --->   | Madakaripura     |   --->   | Third-Party API   |
+-------------+          +------------------+          +-------------------+
                            |         |
                            |         +--> Request Mapper (e.g. Goja)
                            |
                            +--> Adapter (HTTP, future: SOAP, gRPC, graphQL, etc)
                            |
                            +--> Response Mapper
```

## 🤝 Contributing

Contributions are welcome!


To contribute:
1. Create backlog ticket on this board [Madakaripura](https://trello.com/b/1wuLhpMq/madakaripura)
2. Fork this repository
3. Create your feature branch: `git checkout -b my-feature`
4. Commit your changes: `git commit -am 'Add my feature'`
5. Push to the branch: `git push origin my-feature`
6. Open a pull request

Please follow Go best practices and ensure your changes pass linting and tests.

---

## 🛠️ Local Development

### Prerequisites

- Go 1.21+
- (Optional) `air` or `reflex` for live reload
- `curl` or Postman for testing

### Run the app locally

```bash
git clone https://github.com/cobanhub/outbound-gateway.git
cd outbound-gateway
go mod tidy
go run ./cmd/main.go
```