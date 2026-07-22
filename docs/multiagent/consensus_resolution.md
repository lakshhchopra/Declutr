# Declutr Multi-Agent Consensus & Conflict Resolution Guide

When specialist agents produce conflicting results or recommendations, the `CoordinatorAgent` evaluates outputs using confidence scores (0.0–1.0) and empirical evidence.

If consensus confidence falls below the system threshold (0.80), the Coordinator escalates the conflict to the user for human review.
