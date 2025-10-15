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
- `add <task>`: Add a new task
- `list`: Show all tasks
- `done <number>`: Mark the numbered task as done
- `delete <number>`: Delete the numbered task

Examples:
```bash
# add tasks
./todo.exe add "Buy milk"
./todo.exe add "Write report"

# list tasks
./todo.exe list

# mark task 1 as done
./todo.exe done 1

# delete task 2
./todo.exe delete 2
```

## Data File
- Tasks are persisted in `todos.json` in the project directory.
- If the file does not exist, it will be created on first save.

## Notes
- Usage help (when no command is provided):
  - `Usage: go run main.go [add|list|done|delete] [task]`
- The `done` and `delete` commands take the task number as shown in `list`.
