# DrawSync
DrawSync is a high-performance, real-time multiplayer drawing and guessing application inspired by Skribbl. Built with Go, it leverages WebSockets for low-latency state synchronization across multiple clients.

## Features
* **Real-time Synchronization:** Low-latency canvas broadcast using WebSockets.
* **Room Management:** Support for concurrent game instances and private lobbies.
* **Automated Game Loop:** Logic for round rotation, timer management, and word selection.
* **Graceful Lifecycle Management:** Backend implementation handles system signals to ensure clean connection termination and resource cleanup.

## Technical Stack
* **Language:** Go 1.21+
* **Communication:** Gorilla WebSocket
* **Frontend:** HTML5 Canvas / Modern JavaScript
* **Concurrency:** Goroutine-based event loop

## Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/yourusername/drawsync.git
   cd drawsync
   ```

2. **Initialize modules**
   ```bash
   go mod tidy
   ```

3. **Execute the binary**
   ```bash
   go run cmd/drawsync/main.go
   ```

The service defaults to port `8080`.

## Architecture Overview
DrawSync utilizes a centralized Hub pattern to manage client connections. Each game room operates as a separate state machine, broadcasting drawing coordinates and chat messages to participants. The backend is designed to handle graceful shutdowns, ensuring that the `srv.Shutdown(ctx)` pattern is utilized to close active WebSocket connections and clear memory buffers before process exit.

## Configuration
Environment variables can be used to configure the service:
* `PORT`: The port on which the server listens (default: 8080).
* `WRITE_TIMEOUT`: Maximum duration before timing out writes of the response.
* `IDLE_TIMEOUT`: Maximum amount of time to wait for the next request when keep-alives are enabled.

## License
Distributed under the MIT License. See `LICENSE` for more information.

---
