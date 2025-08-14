# teerotate

A simple command-line utility that reads from `stdin` and writes to a rotating log file. It uses the [lumberjack](https://github.com/natefinch/lumberjack) library to handle log rotation automatically.

`teerotate` is designed to be a universal adapter for any application that outputs logs to standard output, allowing you to easily add robust log rotation capabilities without modifying the original application.

## Features

- **Pipeable**: Works seamlessly with standard Unix pipes.
- **Configurable Rotation**: Control log rotation based on file size, age, and number of backups.
- **Automatic Compression**: Optionally compress rotated log files to save space.
- **Simple & Standalone**: A single, statically-linked binary with no external dependencies after building.

## Building

To build the tool from source, you need to have Go installed. 

1.  **Fetch Dependencies**:
    ```bash
    go mod tidy
    ```

2.  **Build the Binary**:
    ```bash
    go build -o teerotate
    ```
    This will create an executable file named `teerotate` in the current directory.

## Usage

The tool is designed to be placed at the end of a command pipeline. It reads every line from its standard input and writes it to the specified log file.

### Synopsis

```bash
# Generic usage
<your_command> | ./teerotate [OPTIONS]
```

### Options

All log rotation parameters are configurable via command-line flags:

| Flag          | Description                                                  | Default Value |
|---------------|--------------------------------------------------------------|---------------|
| `-filename`     | The log file name.                                           | `app.log`     |
| `-max-size`     | Maximum size in megabytes of the log file before rotation.   | `100`         |
| `-max-age`      | Maximum number of days to retain old log files.              | `28`          |
| `-max-backups`  | Maximum number of old log files to retain.                   | `3`           |
| `-local-time`   | Use local time for timestamps in backup files.               | `false`       |
| `-compress`     | Compress rolled-over files using gzip.                       | `false`       |

### Example

Imagine you have a script or application that continuously generates output, and you want to capture it in a log file that rotates once it reaches 5MB.

```bash
# This script simulates a log-producing application
./my_app.sh | ./teerotate \
  -filename="/var/log/my_app.log" \
  -max-size=5 \
  -max-backups=10 \
  -compress=true
```

In this example, `teerotate` will:
- Write all output from `my_app.sh` to `/var/log/my_app.log`.
- When the file exceeds 5 MB, it will be rotated.
- Up to 10 old log files will be kept.
- The rotated files will be compressed with gzip.

```