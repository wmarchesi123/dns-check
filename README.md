# DNS Check
A simple tool to measure DNS response frequency.

---

# Quick Start

### Installation
DNS Check is platform agnostic and can be installed on any system with the following command:
```bash
CGO_ENABLED=0 go install github.com/wmarchesi123/dns-check@latest
```

### Usage

To run DNS Check, use the following command:
```bash
dns-check <domain> <number of requests>
```

Example:
```bash
dns-check google.com 1000
```

Example run:
<img src="assets/ex1.png" alt="A screenshot showing a successful dns-check run." style="zoom:50%;" />

---

[![Release](https://img.shields.io/github/release/wmarchesi123/dns-check.svg?style=flat-square)](https://github.com/wmarchesi123/dns-check/releases/latest) [![Go Report Card](https://goreportcard.com/badge/github.com/wmarchesi123/dns-check?style=flat-square)](https://goreportcard.com/report/github.com/wmarchesi123/dns-check) [![Go Reference](https://pkg.go.dev/badge/github.com/wmarchesi123/dns-check.svg)](https://pkg.go.dev/github.com/wmarchesi123/dns-check) 

