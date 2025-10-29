# CLI Todo

A minimal command-line todo app written in Go. Tasks are stored locally in `todos.json`.

## Requirements
- Go (to build or run from source)
- Windows users can also run the included `todo.exe` binary

## Build
```bash
# from project root
go build -o todo.exe
```

## Run
You can either run the compiled binary or run from source.

```bash
# run the binary (Windows PowerShell)
./todo.exe [command] [args]

# or run from source
go run main.go [command] [args]
```

## Commands
- `add <task> [category] [deadline]`: Add a new task
  - `deadline` format: `YYYY-MM-DD`
  - Duplicate prevention: if a todo with the same `task`, `category`, and `deadline` exists, it won't be added
- `list`: Show all tasks (displays status, optional category, and optional deadline)
- `done <number>`: Mark the numbered task as done
- `delete <number>`: Delete the numbered task

## Examples
```bash
# add tasks (just a task)
./todo.exe add "Buy milk"

# add with category
./todo.exe add "Write report" Work

# add with category and deadline
./todo.exe add "Submit taxes" Finance 2025-04-15

# list tasks (example output)
./todo.exe list
# 1. [ ] Buy milk
# 2. [ ] Write report (Work)
# 3. [ ] Submit taxes (Finance) - due 2025-04-15

# mark task 1 as done
./todo.exe done 1

# delete task 2
./todo.exe delete 2
```

## Data File
- Tasks are persisted in `todos.json` in the project directory.
- If the file does not exist, it will be created on first save.

## Notes
- Usage help shown when no command is provided currently prints:
  - `Usage: go run main.go [add|list|done|delete] [task]`
- The `add` command also accepts optional `category` and `deadline` (as documented above).
- The `done` and `delete` commands take the task number as shown in `list`.
