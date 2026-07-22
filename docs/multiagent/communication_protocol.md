# Declutr Multi-Agent Communication Protocol

Specialist agents NEVER directly invoke each other or exchange raw prompts. All inter-agent communication flows through the event-driven `MessageBus`.

## Message Structure (`AgentMessage`)

- `ID`: Unique message identifier (`msg-...`).
- `CorrelationID`: Transaction correlation ID.
- `GoalID` & `TaskID`: Associated goal and task node references.
- `Sender`: Agent ID of origin (`agt-COORDINATOR_AGENT`, `agt-SEARCH_AGENT`, etc.).
- `Receiver`: Target agent ID or `BROADCAST`.
- `MessageType`: `REQUEST`, `RESPONSE`, `STATUS`, `CONSENSUS_PROPOSAL`.
- `Payload`: Strongly-typed JSON parameter/result payload.
- `Context`: Metadata and execution trace context.
