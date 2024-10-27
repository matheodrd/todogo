# 🗒️Todogo

> [!WARNING]
> This project is still in development and unfinished. Also keep in mind that this is a toy project to experiment and teach myself Go. Lookup [TODO.md](./TODO.md) 😉

A simple stateful CLI application to manage your todos, written in Go.

## Features

- Add, list, select, and update todos
- Persistent state to simplify commands
- Modular and extendable CLI with Cobra
- Todos are stored in a JSON file

## Installation

To install Todogo, clone the repository and build it:

```bash
git clone https://github.com/matheodrd/todogo.git
cd todogo
make build
```

This will create an executable named `todogo`.

## Usage

The CLI has various commands, each with specific options:

### 1. Adding Todos

To add a new todo with a title and a description:

```bash
./todogo add "New Task" -d "Description of the task"
```

### 2. Listing Todos

To list all your todos:

```bash
./todogo list
```

### 3. Selecting a Todo

To select a todo by ID (for example, to update it later):

```bash
./todogo select <id>
```

### 4. Updating Todos

Once a todo is selected, you can update it using the `update` command with one of its subcommands:

- **Title**: Update the title of the selected todo
  ```bash
  ./todogo update title "Updated Title"
  ```

- **Description**: Update the description
  ```bash
  ./todogo update description "New description text"
  ```

- **Status**: Update the status (`todo`, `doing`, `done`)
  ```bash
  ./todogo update status done
  ```

## Configuration and Caching

The CLI stores the selected todo ID in a cache file located at `$XDG_CACHE_HOME/todogo/cache.yml` (usually `~/.cache/todogo/cache.yml`). This allows commands like `update` to apply changes to the currently selected todo without needing to specify its ID every time. I was inspired by the [kubectl](https://github.com/kubernetes/kubectl) project and [rwxrob](https://rwxrob.github.io/zet/1729/).

## Testing

To run unit tests, use the following command:

```bash
go test ./...
```

## License

This project is licensed under the Apache 2.0 License.
