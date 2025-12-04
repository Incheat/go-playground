# sqlc Guide

This guide shows how to use **sqlc** to generate type-safe Go code from raw SQL.

sqlc basically does three things:

1. You write SQL schema & queries.
2. You run `sqlc generate`.
3. It outputs Go types & methods for you to call from your app. :contentReference[oaicite:0]{index=0}  

---

## 1. Install sqlc

You need **Go 1.21+** if you install via `go install`. :contentReference[oaicite:1]{index=1}  

Common options:

```bash
# Using Go (works anywhere with Go installed)
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

# macOS (Homebrew)
brew install sqlc

# Ubuntu (Snap)
sudo snap install sqlc

# Docker
docker pull sqlc/sqlc
# Example usage:
docker run --rm -v "$(pwd)":/src -w /src sqlc/sqlc generate
