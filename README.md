# Goredis: Understanding Redis by Building from Scratch

This repository, `goredis`, is part of my journey to understand how Redis works by building a Redis-like implementation from scratch. The inspiration and guidance for this project come from resources like [Build Redis from Scratch](https://www.build-redis-from-scratch.dev).

## Features

- Implements core concepts from Redis
- Uses the RESP protocol for communication
- Aimed at gaining a deeper understanding of Redis internals

## Limitations

1. **Concurrency**:
   - This implementation **does not support concurrency**. It only handles one client at a time. Redis achieves its concurrency model using an event loop and handles multiple clients simultaneously. However, this implementation lacks proper synchronization for concurrent access to shared resources.
2. **Redis Server Conflict**:

   - If the official Redis server is running on port `6379`, you must stop it before running `goredis` to avoid port conflicts. You can stop the Redis server with the following command:
     ```bash
     sudo systemctl stop redis
     ```
   - Alternatively, you can modify the port number in the code to avoid conflicts with the existing Redis server.

3. **Incomplete Feature Set**:

   - This project is a learning exercise and only implements a subset of Redis commands. Advanced Redis features such as replication, persistence (AOF or RDB), clustering, and complex data structures like sorted sets and hyperloglogs are not yet implemented.

4. **Command Parsing**:
   - The command parsing logic assumes valid RESP messages and may not handle all edge cases.

## Using `goredis` with Redis CLI (`redis-cli`)

You can use the official `redis-cli` client to communicate with `goredis`. Here's how:

1. **Start `goredis`**:

   ```bash
   make run
   ```

2. **Use `redis-cli`**:

   - Once the server is running, connect to it using the Redis CLI:
     ```bash
     redis-cli
     ```
   - You can then send commands to the `goredis` server just like with Redis. Example:
     ```bash
     PING
     SET key value
     GET key
     HSET users u1 Ian
     HSET posts u1 Hello World
     HGET users u1
     HGETALL users
     ```

3. **Stop the Official Redis Server**:
   - Ensure that the official Redis server is stopped to avoid port conflicts since both `goredis` and Redis use the default port `6379`. You can stop Redis with the command:
     ```bash
     sudo systemctl stop redis
     ```

## How to Use

### Installation

Clone the repository and initialize it as a Go module:

```bash
git clone https://github.com/iangechuki/goredis.git
cd goredis
go mod tidy
```

### Running the Server

Start the server with:

```bash
make run
```

You should see the following output, indicating the server is ready:

```
Listening on port :6379
```

### Sending Commands

You can use any RESP-compliant client, including `redis-cli`, to send commands. Basic supported commands include:

- `SET key value` – Store a value under a key
- `GET key` – Retrieve a value by its key
- `HSET hash key value` – Set a field in a hash
- `HGET hash key` – Get a field value from a hash
- `PING` – Return a `PONG`
- `HGETALL hash` – Retrieve all fields and values of a hash

## Contributing

This project is a learning tool, and contributions are welcome! You can fork the repository, make improvements, and submit pull requests.

**Disclaimer**: This project is not intended to replace Redis in production environments. It is a simplified educational tool for understanding Redis internals.
