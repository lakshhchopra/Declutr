# Declutr Multi-Agent Task Planner Guide

The `MultiAgentTaskPlanner` constructs Directed Acyclic Graph (DAG) task execution plans (`TaskGraph`).

## Execution Node Types

- **Parallel Execution (`PARALLEL`)**: Tasks without mutual dependencies (e.g. Search Agent and Knowledge Agent scanning simultaneously) run in parallelgoroutines.
- **Sequential Execution (`SEQUENTIAL`)**: Dependent steps waiting on prerequisite output task completion.
- **Approval Checkpoints**: Tasks marked as sensitive wait for human confirmation before execution.
