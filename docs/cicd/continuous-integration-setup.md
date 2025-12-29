# Continuous Integration Setup for Go Projects  
### Go Test + golangci-lint + Pre-commit Hooks + GitHub Actions

This document explains how to enforce code quality in a Go project using:

- **golangci-lint**  
- **Go tests**  
- **Local pre-commit hooks**  
- **GitHub Actions CI**  
- **Branch protection rules**

This ensures no bad code can be committed or merged.

---

# 1. What is golangci-lint?

`golangci-lint` is a fast, powerful Go linter aggregator that runs 30+ linters in parallel, including:

- `govet`
- `staticcheck`
- `revive`
- `ineffassign`
- `unused`
- `errcheck`
- `gofmt`
- `goimports`

### Install:

```bash
brew install golangci-lint
```
or
```bash
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s
```
---

# 2. Local CI: Pre-commit Hook

Block commits if lint or tests fail

## Option A — Simple Git pre-commit hook

### Create:

```bash
.git/hooks/pre-commit
```

### Add:

```bash
#!/bin/sh

echo "Running golangci-lint..."
golangci-lint run ./...
if [ $? -ne 0 ]; then
    echo "❌ golangci-lint failed! Commit aborted."
    exit 1
fi

echo "Running go test..."
go test ./...
if [ $? -ne 0 ]; then
    echo "❌ Tests failed! Commit aborted."
    exit 1
fi

echo "✔ All checks passed. Commit allowed."
exit 0
```

### Make executable:

```bash
chmod +x .git/hooks/pre-commit
```

## Option B — Using pre-commit framework (recommended)

### Create:

```arduino
.pre-commit-config.yaml
```

### Add:

```yaml
repos:
  - repo: local
    hooks:
      - id: golangci-lint
        name: Run golangci-lint
        entry: golangci-lint run ./...
        language: system
        pass_filenames: false

      - id: go-test
        name: Run go test
        entry: go test ./...
        language: system
        pass_filenames: false
```

### Install:

```bash
pip install pre-commit
pre-commit install
```

Now git commit runs lint + tests automatically.
---

# 3. Remote CI: GitHub Actions

## Automatically run lint + tests on push and PR

### Create:

```bash
.github/workflows/ci.yml
```

### Add:

```yaml
name: CI

on:
  push:
    branches:
      - main
      - develop
  pull_request:

jobs:
  build-and-test:
    runs-on: ubuntu-latest

    steps:
      - name: Checkout code
        uses: actions/checkout@v4

      - name: Set up Go
        uses: actions/setup-go@v5
        with:
          go-version: '1.22'

      - name: Cache Go modules
        uses: actions/cache@v4
        with:
          path: |
            ~/go/pkg/mod
            ~/.cache/go-build
          key: ${{ runner.os }}-go-${{ hashFiles('**/go.sum') }}
          restore-keys: |
            ${{ runner.os }}-go-

      - name: Download dependencies
        run: go mod download

      - name: Run golangci-lint
        uses: golangci/golangci-lint-action@v6
        with:
          version: latest
          args: ./...

      - name: Run Go tests
        run: go test ./... -v
```

This ensures lint + tests run automatically for every push and pull request.
---

# 4. Enforce CI with Branch Protection

### Go to:

## GitHub → Settings → Branches → Branch Protection Rules

### Enable:

- Require pull request
- Require status checks to pass before merging
- Select your CI job (e.g., “CI”)
- Optional: Require branches to be up to date

This ensures the main branch remains clean and stable.
---

# 5. Recommended Workflow Summary

| Stage         | Tool              | Purpose                               |
|---------------|-------------------|----------------------------------------|
| Local commit  | Pre-commit hook   | Prevent committing failed lint/test    |
| Push / PR     | GitHub Actions    | Run full CI checks                     |
| Merge to main | Branch Protection | Prevent merging broken code            |


# 6. Suggested Repo Files

```bash
.git/hooks/pre-commit               # (optional local hook)
.pre-commit-config.yaml             # (recommended)
.golangci.yml                       # golangci-lint config
.github/workflows/ci.yml            # GitHub Actions pipeline
```

### Example .golangci.yml:

```yaml
run:
  timeout: 5m

linters:
  enable:
    - govet
    - revive
    - staticcheck
    - errcheck
    - goimports
    - ineffassign
    - gosimple
    - unused

issues:
  exclude-use-default: false
```

# Conclusion

With this setup:
- Bad code cannot be committed
- PRs must pass lint + tests
- The main branch stays clean
- Your Go project achieves professional-level CI quality