# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Overview

This is a multi-language code examples repository for the [Intelligent API](https://api.intelligent-api.com). Each endpoint is implemented in 8 languages (C#, cURL, Go, Java, Node.js, Python, Ruby, Rust) with two auth variants each. All examples are self-contained with minimal external dependencies.

## Repository Structure

```
/{language}/{endpoint-name}/
  ├── basic-auth.{ext}     # Base64-encoded client_id:client_secret in Authorization header
  └── token-auth.{ext}     # OAuth token with scope parameter
```

**Exception — C#**: each auth variant is a separate subdirectory with `Program.cs` + `.csproj`.

Languages: `csharp/`, `cURL/`, `golang/`, `java/`, `nodejs/`, `python/`, `ruby/`, `rust/`

## Running Examples

All examples use placeholder values `[[like_this]]` that must be replaced before execution.

| Language | Prerequisites | Command |
|----------|--------------|---------|
| Python | Python 3 | `python3 basic-auth.py` |
| Node.js | Node.js v20+ | `node basic-auth.js` |
| C# | .NET 8 SDK | `dotnet run` (from the `basic-auth/` or `token-auth/` project folder) |
| Go | Go | `go run basic-auth.go` |
| Java | JDK | `javac BasicAuth.java && java BasicAuth` |
| Ruby | Ruby | `ruby basic-auth.rb` |
| Rust | Cargo | `cargo run` (from the `basic-auth/` or `token-auth/` project folder) |
| cURL | curl | `bash basic-auth.sh` |

## Template Variables

| Variable | Description |
|----------|-------------|
| `[[client_id]]` | Client ID from the API dashboard |
| `[[client_secret]]` | Client secret from the API dashboard |
| `[[scopes]]` | Space-separated OAuth scopes (token-auth only) |
| `[[text_body]]` | Text input for text/language/PII endpoints |
| `[[source_language]]` | BCP-47 source language code (e.g. `en`) |
| `[[target_language]]` | BCP-47 target language code (e.g. `es`) |
| `[[filepath]]` | Absolute path to file for document/image/speech endpoints |
| `[[content_type]]` | MIME type for file uploads (e.g. `application/pdf`, `image/jpeg`) |
| `[[minimum_score]]` | 0.0–1.0 confidence threshold for PII detection |
| `[[url]]` | Target URL for web scraper endpoints |

## API Endpoints

| Folder name | Endpoint |
|-------------|----------|
| `detect-language` | `POST /v1/language/detect` |
| `translate-language` | `POST /v1/language/translate` |
| `detect-pii` | `POST /v1/pii/detect` |
| `redact-pii` | `POST /v1/pii/redact` |
| `analyse-sentiment` | `POST /v1/text/sentiment` |
| `summarise-text` | `POST /v1/text/summarize` |
| `count-calories-text` | `POST /v1/text/calories` |
| `text-classify` | `POST /v1/text/classify` |
| `text-classify-multiple` | `POST /v1/text/classify/multiple` |
| `text-expenses` | `POST /v1/text/expenses` |
| `extract-document-text` | `POST /v1/document/text` |
| `extract-document-expenses` | `POST /v1/document/expenses` |
| `extract-document-identity` | `POST /v1/document/identity` |
| `list-images` | `POST /v1/image/list` |
| `count-calories-image` | `POST /v1/image/calories` |
| `transcribe-speech` | `POST /v1/speech/transcribe` |
| `scrape-markdown` | `POST /v1/scraper/markdown` |
| `scrape-summarise` | `POST /v1/scraper/summarize` |

## Authentication Patterns

**Basic auth** — all languages encode `client_id:client_secret` in Base64 and send it as `Authorization: Basic {encoded}`.

**Token auth** — exchanges credentials + scopes for a Bearer token via `POST /v1/oauth/token`, then uses `Authorization: Bearer {token}` on the target endpoint. The required scope per endpoint is documented in the language-specific READMEs.

## Adding a New Endpoint

Follow the existing pattern: create a folder under each language directory named after the endpoint action, add `basic-auth.{ext}` and `token-auth.{ext}` (or subdirectories for C#/Rust), and update the root `README.md` and each language's `README.md` with the new entry.
