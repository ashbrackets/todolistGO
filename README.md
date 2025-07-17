# todolistGO
A Command-Line Interface (CLI) To-Do List Application

## Setup

To set up todolistGO on your machine, you'll need Git and Go installed.

### Installation Steps

1.  **Clone the Repository:** Open your terminal or command prompt and clone the todolistGO repository.
    ```bash
    git clone https://github.com/ashbrackets/todolistGO.git
    ```
2.  **Navigate to the Directory:** Change your current directory to the newly cloned project folder.
    ```bash
    cd todolistGO
    ```
3.  **Build the Executable:** Compile the Go source code into an executable file.
    * **Windows:**
        ```bash
        go build -o todo.exe
        ```
        This creates `todo.exe` in the current directory.
    * **macOS / Linux:**
        ```bash
        go build -o todo
        ```
        This creates an executable file named `todo` (without an extension) in the current directory.

    *(Optional: To make `todo` globally accessible without `./` or specifying the full path, you can move the compiled executable to a directory included in your system's PATH, like `/usr/local/bin` on macOS/Linux or a custom bin directory on Windows. For example, on macOS/Linux: `sudo mv todo /usr/local/bin/`)*

## How to Use

---

Once `todo` (or `todo.exe` on Windows) is built, you can use the following commands to manage your tasks:

* **Add a Task:**
    ```bash
    ./todo add "Your task description here"
    ```
    Adds a new task with the specified description to your to-do list.

* **Update Task Status:**
    ```bash
    ./todo update <id> <status>
    ```
    Modifies the status of a task.
    * `<id>`: The numerical ID of the task you want to update.
    * `<status>`: Can be `check`, `c`, `1` (to mark as complete) or `uncheck`, `u`, `0` (to mark as incomplete).

* **Delete Tasks:**
    ```bash
    ./todo delete <id>
    ```
    Deletes a specific task by its ID.
    ```bash
    ./todo delete all
    ```
    Deletes all tasks from your list.

* **List Tasks:**
    ```bash
    ./todo list
    ```
    Displays all tasks in your to-do list.
    ```bash
    ./todo list check
    ```
    Shows only completed tasks.
    ```bash
    ./todo list uncheck
    ```
    Shows only incomplete tasks.

* **Get Help:**
    ```bash
    ./todo help
    ```
    Provides information on available commands and their usage.

## Credits
https://roadmap.sh/projects/task-tracker
