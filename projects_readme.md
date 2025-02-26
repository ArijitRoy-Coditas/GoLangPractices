# Todo List Project

## Overview
The Todo List project is a simple command-line application built with Go. It allows users to add, modify, delete, and read tasks from a CSV file.

## Project Structure

- `Actions/` - Contains functions to add, delete, modify, and read tasks.
- `data/` - Stores the `tasks.csv` file where task data is saved.
- `ModifyOption/` - Contains helper functions for modifying specific aspects of tasks.
- `utils/` - Utility functions for handling console operations.
- `go.mod` & `go.sum` - Dependency management files.
- `main.go` - Entry point of the application.

## How to Run

### Prerequisites
Make sure you have Go installed. You can check with:
```sh
 go version
```

### Steps to Run
1. Navigate to the project directory:
   ```sh
   cd Projects/Todo_list/
   ```
2. Run the application:
   ```sh
   go run main.go
   ```

## Usage
Once running, the application will prompt you with options to:
- Add a task
- Modify a task
- Delete a task
- View tasks

Follow the on-screen instructions to interact with the app.

## Notes
- The tasks are stored in `data/tasks.csv`.
- Ensure that the `tasks.csv` file exists before running the app, or the application will create it automatically.

Enjoy using the Todo List application!

