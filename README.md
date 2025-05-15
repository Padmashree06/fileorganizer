# Go & Organize 
**File and Folder Organizer Using Go**

A simple command-line tool written in **Go** that organizes files in a given directory based on their file types (e.g., `.jpg`, `.png`, `.pdf`, etc.). The program will scan the directory, classify files according to their extensions, and move them to corresponding folders based on their types.

## Features

- **Automatic Organization**: Organizes files based on their extensions into respective folders (e.g., `Images`, `Documents`, `Audio`, etc.).
- **No Memory Wastage**: The program creates directories only for the file types present in the input folder, ensuring efficient resource usage.
- **Cross-platform**: Works on Windows, macOS, and Linux.

## Prerequisites
- **Go**: Ensure you have [Go](https://golang.org/dl/) installed on your system. You can check if Go is installed by running:
  ```bash
  go version
