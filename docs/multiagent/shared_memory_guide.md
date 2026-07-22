# Declutr Shared Multi-Agent Memory Guide

The Shared Memory workspace (`SharedMemoryItem`) acts as a central repository for cross-agent collaboration state during goal execution.

## Memory Categories

- `TASK_RESULT`: Output payloads produced by completed specialist tasks.
- `USER_PREFERENCE`: Inferred or explicitly stated user preferences.
- `CONTEXT`: Goal execution context shared across agents.
- `STATE`: Pipeline execution status markers.
