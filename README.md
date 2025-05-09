# Madakaripura

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


## 🧱 Architecture

```text
+-------------+          +------------------+          +-------------------+
| Core System |   --->   | Madakaripura     |   --->   | Third-Party API   |
+-------------+          +------------------+          +-------------------+
                            |         |
                            |         +--> Request Mapper (e.g. Goja)
                            |
                            +--> Adapter (HTTP, future: SOAP, gRPC)
                            |
                            +--> Response Mapper