# gosec Guide

## What is gosec?

**gosec** is a Static Application Security Testing (SAST) tool for **Go (Golang)**.  
It analyzes Go source code to identify **potential security vulnerabilities** before runtime.

gosec is commonly used in:
- Secure Go development
- Code reviews
- CI/CD security pipelines

---

## What gosec detects

gosec scans source code for a wide range of common security issues, including:

- Hardcoded credentials (passwords, API keys, tokens)
- SQL injection vulnerabilities
- Command injection (`os/exec` misuse)
- Weak cryptographic algorithms (`md5`, `sha1`)
- Insecure file permissions (`chmod 0777`)
- TLS and certificate misconfigurations
- Unsafe random number generation
- Unsafe deserialization

---

## Installation

### Prerequisites
- Go **1.20+**
- `$GOPATH/bin` or `$HOME/go/bin` added to your `PATH`

### Install using `go install` (recommended)

```bash
go install github.com/securego/gosec/v2/cmd/gosec@latest
```

Verify installation:

```bash
gosec --version
```

---

## Basic Usage

### Scan an entire project

From the root of your Go module:

```bash
gosec ./...
```

This recursively scans all Go files.

---

### Scan a specific directory or file

```bash
gosec ./cmd/api
gosec main.go
```

---

## Example Output

```text
G101: Potential hardcoded credentials
  > main.go:12
```

Each finding includes:
- Rule ID (e.g., `G101`)
- Description
- File name and line number
- Severity and confidence (in detailed formats)

---

## Common Options

### Output formats

```bash
gosec -fmt=json ./...
```

Supported formats:
- `text` (default)
- `json`
- `yaml`
- `sarif` (recommended for GitHub Security)

---

### Generate a report file

```bash
gosec -fmt=sarif -out=gosec.sarif ./...
```

---

### Exclude specific rules

```bash
gosec -exclude=G101,G204 ./...
```

---

### Exclude directories

```bash
gosec -exclude-dir=vendor,testdata ./...
```

---

### Filter by severity

```bash
gosec -severity=high ./...
```

Only reports **high severity** issues.

---

## Suppressing False Positives

### Inline suppression (recommended)

```go
// #nosec G204 -- command input is validated earlier
cmd := exec.Command(userInput)
```

### File-level suppression

```go
// #nosec
```

⚠️ Use suppressions sparingly and always document why the code is safe.

---

## Using gosec in CI/CD (GitHub Actions example)

```yaml
- name: Run gosec
  run: |
    go install github.com/securego/gosec/v2/cmd/gosec@latest
    gosec ./...
```

By default, the pipeline fails if issues are detected.

---

## Limitations

- gosec performs **static analysis only**
- False positives are possible
- Does not replace:
  - Code review
  - Dependency scanning
  - Runtime security testing

---

## Summary

| Feature | Description |
|------|------|
| Tool type | Static security analyzer |
| Language | Go |
| Installation | `go install` |
| Basic usage | `gosec ./...` |
| Best use | Early detection of common Go security issues |

---

**Recommendation:**  
Use gosec early in development and combine it with dependency scanning and manual code reviews for best security coverage.
