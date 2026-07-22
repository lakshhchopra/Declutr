# Declutr Multi-Agent Coordinator Guide

The Coordinator Agent (`CoordinatorAgent`) serves as the orchestrator of multi-agent intelligence.

```
User Goal → Coordinator Agent → Task Planner → Specialist Agents → Shared Memory → Execution → Review → Response
```

## Coordinator Responsibilities

1. Parse user goal objectives.
2. Invoke `MultiAgentTaskPlanner` to decompose goals into parallel & sequential DAG task execution graphs.
3. Dispatch task execution requests to registered specialist agents via the `MessageBus`.
4. Monitor execution state, handle task retries, and recover from specialist agent failures.
5. Aggregate specialist outputs in `SharedMemory` and evaluate consensus using `ConsensusResult`.
