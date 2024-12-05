# Goredis: Understanding Redis by Building from Scratch

This repository, `goredis`, is part of my journey to understand how Redis works by building a Redis-like implementation from scratch. The inspiration and guidance for this project come from resources like [Build Redis from Scratch](https://www.build-redis-from-scratch.dev).

## Features

- Implements core concepts from Redis.
- Uses the RESP protocol for communication.
- Aimed at gaining a deeper understanding of Redis internals.

## Structure

- **RESP Writer**: Implements the RESP protocol for encoding/decoding messages.
- **Command Parsing**: Handles parsing and execution of simple Redis-like commands.

## How to Use

### Installation

Clone the repository and initialize it as a Go module:

```bash
git clone https://github.com/iangechuki/goredis.git
cd goredis
go mod tidy
```
