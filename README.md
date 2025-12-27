# Mini-Redi-Lite

A Redis-like in-memory database written in Go.

## Features
- just for fun and education how redis server works so i build in powerfull performance of golang
- TCP server (port 6379) with Redis-style commands
- HTTP API for GET/SET operations later i add more commands
- In-memory storage
- Binary file persistence (data.bin)
- Concurrent clients supported

## Quick Start

```bash
git clone https://github.com/<your-username>/MiniRedis.git
cd MiniRedis
go run .
