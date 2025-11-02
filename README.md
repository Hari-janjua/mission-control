# Mission Control Project

## Overview

**Mission Control** is a secure, asynchronous **command & control system** simulating a military-style operation.

- The **Commander’s Camp** issues orders (missions).
- The **Soldier Worker** executes them asynchronously.
- All communication is **one-way and secure**, using **Kafka** as a message hub.
- Soldiers authenticate via short-lived **JWT tokens** (rotated every 30 seconds).
- **MariaDB** stores mission status.
- The system is **fully containerized** via Docker Compose.

---

## 🧩 Architecture Overview

### Components

| Service            | Description                                                          |
|--------------------|----------------------------------------------------------------------|
| **Commander**      | Issues missions, tracks their progress, validates soldier tokens.    |
| **Soldier Worker** | Receives missions, executes them concurrently, sends status updates. |
| **Kafka**          | Acts as a central, secure communication hub for all message passing. |
| **MariaDB**        | Stores mission data (status, timestamps, command).                   |

---

## 🏗️ Architecture Pattern

The entire codebase follows **Hexagonal (Ports & Adapters)** architecture.

