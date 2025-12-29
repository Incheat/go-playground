# koanf Guide

> Lightweight, extensible configuration management for Go  
> Repo: `github.com/knadh/koanf/v2` :contentReference[oaicite:0]{index=0}  

---

## 1. What is koanf?

**koanf** is a Go library for reading configuration from multiple sources (files, env vars, flags, remote stores, etc.) and merging them into a single config tree. It supports many formats: JSON, YAML, TOML, HCL, dotenv, etc.:contentReference[oaicite:1]{index=1}  

Key ideas:

- **Providers** – *where* config comes from (file, env, flags, S3, etc.).
- **Parsers** – *how* raw bytes are turned into a nested `map[string]interface{}` (YAML, JSON, TOML, …).
- **Delimiter** – paths like `app.server.port` let you access nested configuration.:contentReference[oaicite:2]{index=2}  

---

## 2. Installation

Install core + at least one provider and parser:

```bash
# Core
go get -u github.com/knadh/koanf/v2

# File provider (example)
go get -u github.com/knadh/koanf/providers/file

# YAML parser (example)
go get -u github.com/knadh/koanf/parsers/yaml
