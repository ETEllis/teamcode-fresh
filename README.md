# TeamCode

⌬ Terminal-based AI coding team assistant with multi-agent collaboration support.

## Overview

TeamCode is a fork of OpenCode that adds first-class support for multi-agent team collaboration. It enables:

- **Team Context**: Shared charter, roles, goals, and working agreements
- **Task Board**: Kanban-style task tracking with WIP limits
- **Handoff Protocol**: Explicit task transitions between agents
- **Decision Log**: Searchable record of architectural decisions
- **Escalation Queue**: Structured blocked task resolution
- **Shared Artifacts**: Team docs, specs, and diagrams
- **Retrospectives**: Sprint retrospectives for continuous improvement
- **Private Inboxes**: Each agent has their own message inbox

## Architecture

```
TeamCode (Go)
    │
    └── internal/team/        # Go team package
            │
            └── pywrapper.go  # Python bridge
                    │
                    └── claude_teams  # Python library
                            │
                            └── ~/.claude/teams/<team>/
                                    ├── team-context.json
                                    ├── task-board.json
                                    ├── handoffs/
                                    ├── inboxes/      # Private per agent
                                    ├── decisions/
                                    └── ...
```

## Building

```bash
go build -o teamcode .
```

## Usage

```bash
# Interactive mode
./teamcode

# Non-interactive mode
./teamcode -p "Explain this codebase"

# With debug logging
./teamcode -d
```

## Team Tools

When operating in team mode, the following tools become available:

- `team_create_context` - Create team with charter and roles
- `team_add_role` - Add role definitions
- `team_assign_role` - Assign agents to roles
- `task_create` - Add tasks to the board
- `task_move` - Move tasks between columns
- `handoff_create` - Create task handoffs
- `handoff_accept` - Accept pending handoffs
- `inbox_read` - Read your messages

## License

MIT
