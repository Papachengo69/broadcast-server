# Broadcast Server

A simple CLI WebSocket broadcast server and client implemented in Go.

## Overview

This project demonstrates real-time communication using WebSockets in Go.  
- The **server** accepts multiple client connections and broadcasts any message received to all connected clients.  
- The **client** connects to the server, sends messages, and receives broadcasts in real-time.

---

## Features

- CLI-based server and client using the same binary.
- Broadcast messages to all connected clients.
- Graceful handling of client connections and disconnections.
- Uses [Gorilla WebSocket](https://github.com/gorilla/websocket) for WebSocket communication.
- Command parsing using [Cobra](https://github.com/spf13/cobra).

---

## Getting Started

### Prerequisites

- Go 1.20 or later installed
- Git

### Installation

Clone this repository:

```bash
git clone https://github.com/Papachengo69/broadcast-server.git
cd broadcast-server
go mod tidy
