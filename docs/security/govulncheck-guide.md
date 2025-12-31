# govulncheck Guide

`govulncheck` is a **security analysis tool for Go** that scans your Go code and dependencies for **known vulnerabilities** using the official **Go vulnerability database**.

It is maintained by the Go team and is the **recommended way** to check Go projects for vulnerabilities.

---

## What govulncheck Does

`govulncheck` analyzes:

- Direct and transitive dependencies
- Your actual code paths (not just `go.mod`)
- Known vulnerabilities from https://vuln.go.dev

It reports:

- Vulnerability IDs (for example `GO-2023-1987`)
- Whether the vulnerability is actually **reachable**
- Call stacks showing how your code reaches vulnerable functions
- Suggested fixes or upgrade versions

This makes it more precise than dependency-only scanners.

---

## Requirements

- Go **1.18+** (latest Go version recommended)

---

## Installation

Install using the Go toolchain:

```bash
go install golang.org/x/vuln/cmd/govulncheck@latest
```

Verify installation:

```bash
govulncheck -h
```

If the command is not found, ensure your GOPATH bin directory is in your PATH:

```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

---

## Basic Usage

Scan the current module:

```bash
govulncheck ./...
```

This scans all packages in the module and reports vulnerabilities that affect reachable code.

---

## Common Usage Patterns

### Scan a Specific Package

```bash
govulncheck ./cmd/server
```

### Scan a Single File

```bash
govulncheck main.go
```

### Dependency-Only Scan (Faster)

```bash
govulncheck -mode=module
```

This mode checks dependencies without analyzing call paths. It is useful for quick checks or CI runs.

---

## Output Formats

### Human-Readable Output (Default)

```bash
govulncheck ./...
```

### JSON Output (CI / Automation)

```bash
govulncheck -json ./... > vuln-report.json
```

---

## Understanding the Output

A typical report includes:

- Vulnerability ID
- Affected module and versions
- Call stack from your code to the vulnerable function
- Fixed versions (if available)

Example (simplified):

```
Vulnerability: GO-2023-1987
Module: golang.org/x/net
Call stack:
  main.main → handler → vulnerableFunc
Fixed in: v0.17.0
```

If no vulnerabilities are found:

```
No vulnerabilities found.
```

---

## Using govulncheck in CI

Example GitHub Actions step:

```yaml
- name: Run govulncheck
  run: |
    go install golang.org/x/vuln/cmd/govulncheck@latest
    govulncheck ./...
```

The command exits with a non-zero status if vulnerabilities are detected, allowing CI to fail builds.

---

## When to Use govulncheck vs Other Tools

| Tool | Purpose |
|-----|--------|
| govulncheck | Precise vulnerability detection based on code paths |
| go list -m -u | Find outdated dependencies |
| dependabot | Automated dependency updates |
| staticcheck | Code quality and bug detection |

---

## Best Practices

- Run `govulncheck` locally before releases
- Integrate it into CI pipelines
- Regularly update dependencies
- Combine with other static analysis tools

---

## Summary

- `govulncheck` is the official Go vulnerability scanner
- It detects real, reachable vulnerabilities
- Easy to install and automate
- Suitable for both local development and CI pipelines
