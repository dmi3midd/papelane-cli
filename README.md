<p align="center">
  <img src="assets/logo.png" alt="macauth logo" width="300">
</p>

# Papelane CLI Documentation

> [!IMPORTANT]
> **Attention:** You must have **Docker** installed and running on your machine for this utility to work properly. Papelane CLI uses a local Telegram Bot API server running in a Docker container to bypass the standard 50 MB file upload limit imposed by Telegram bots.

**Papelane CLI** is a high-performance command-line utility that transforms your Telegram bot into structured cloud storage. Thanks to a local SQLite database, it provides instant navigation through a Virtual File System (VFS).

---

## 📦 Installation

To install Papelane CLI, simply use the `go install` command:

```bash
go install github.com/dmi3midd/papelane-cli@latest
```
*(Make sure your `$GOPATH/bin` or `~/go/bin` is added to your system's `PATH` so you can run the `papelane` command globally.)*

---

## 🚀 Initialization and Checking

### `init`
Initializes the application, creates configuration files (in your `os.UserConfigDir()`), and sets up Telegram connection parameters.
- **Required flags:**
  - `--apid` — Your `TELEGRAM_API_ID`.
  - `--apih` — Your `TELEGRAM_API_HASH`.
  - `--token` — Your Telegram bot token.
  - `--cid` — Your Chat ID (where files will be stored).
- **Optional flags:**
  - `--port` — Port for the local Telegram Bot API Docker container (default is `8081`).
  - `--sa` — Always stop the container (true/false).
*Example:* `papelane init --apid 123456 --apih abcdef... --token 123:abc... --cid 987654321`

### `check`
Checks if Docker is installed, if the Telegram Bot API image is downloaded, and if the corresponding container is running. If the container doesn't exist, this command will automatically create and start it.
*Example:* `papelane check`

### `ping`
Checks the availability of the local Telegram Bot API server (makes a request to the `/getMe` endpoint).
*Example:* `papelane ping`

---

## 📁 Virtual File System (VFS) Navigation

### `curr`
Prints the current virtual directory (path) you are in.
*Example:* `papelane curr`

### `cd <path>`
Changes the current virtual directory. Supports navigating deeper (e.g., `folder1/folder2`) and going one level up (`..`).
*Example:* `papelane cd myfolder` or `papelane cd ..`

### `root`
Quickly returns you to the root directory (`root`) of the virtual file system.
*Example:* `papelane root`

### `ls`
Lists files and folders in the current virtual directory.
- **Flags:**
  - `-f`, `--files` — Show only files.
  - `-d`, `--dirs` — Show only directories.
*Example:* `papelane ls -f`

---

## 📝 File and Folder Management

### `mkdir <folder_name>`
Creates a new folder in the current virtual directory.
- **Flags:**
  - `-d`, `--cd` — Automatically switch (`cd`) to the newly created folder.
*Example:* `papelane mkdir photos -d`

### `upload <local_file_path>`
Uploads a file from your computer to the current virtual directory in your Telegram cloud storage.
*Example:* `papelane upload ./my-document.pdf`

### `download <storage_file_name>`
Downloads a file from the cloud storage to your computer. By default, files are saved to the `~/Downloads` folder.
- **Flags:**
  - `--out` — Specify a custom local path to save the file.
*Example:* `papelane download my-document.pdf --out ./saved-docs`

### `rmd <folder_name>`
Deletes the specified virtual directory from the current folder.
*Example:* `papelane rmd photos`

### `rmf <file_name>`
Removes the specified file from the virtual file system and **physically deletes it from Telegram**.
*Example:* `papelane rmf my-document.pdf`

### `cleanc`
Cleans the file cache of the local Telegram Bot API server (cleans the Docker Volume) to free up disk space on your computer.
*Example:* `papelane cleanc`
