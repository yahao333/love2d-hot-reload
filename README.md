# love2d-hot-reload

A hot-reload tool for LÖVE2D projects implemented in Go (Golang). This tool leverages Go's powerful file monitoring and process management capabilities to create an efficient development environment for LÖVE2D projects.

## Features

- Real-time file monitoring for `.lua` files and resources
- Automatic LÖVE2D process management
- Cross-platform support (primarily tested on macOS)
- Debounced reload mechanism to prevent excessive restarts
- Configurable file ignore patterns
- Environment variable configuration support

## Prerequisites

- Go 1.22 or later
- LÖVE2D installed on your system
- Basic understanding of Go and LÖVE2D development

## Installation

1. Clone the repository:
```bash
git clone https://github.com/yahao333/love2d-hot-reload.git
cd love2d-hot-reload
```

2. Install dependencies:
```bash
go mod download
```

3. Build the project:
```bash
make
```
or manually:
```bash
go build -o build/hot-reload.bin src/main.go
```

## Usage

1. Place the `build/hot-reload.bin` executable in your LÖVE2D project directory
2. (Optional) Set environment variables for configuration:
```bash
export LOVE2D_DEBOUNCE_TIME=1000  # Set debounce time in milliseconds (default: 1000ms)
```
3. Run the tool:
```bash
./hot-reload.bin
```

The tool will:
- Start your LÖVE2D project
- Monitor all relevant files for changes
- Automatically restart LÖVE2D when changes are detected

## How It Works

- **File Monitoring**: Uses `fsnotify` to detect file system changes
- **Process Management**: Controls the LÖVE2D process, restarting it when changes are detected
- **Debouncing**: Implements a debounce mechanism to prevent rapid consecutive restarts
- **File Filtering**: Ignores temporary files and non-relevant changes

## Configuration

The tool can be configured in two ways:

1. Environment Variables:
   - `LOVE2D_DEBOUNCE_TIME`: Set debounce time in milliseconds (default: 1000ms, minimum: 100ms)

2. Source Code Configuration:
   Modify the following in `src/main.go`:
   - Project directory path
   - Ignored file patterns
   - LÖVE2D executable path

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

- [LÖVE2D](https://love2d.org/) - The game framework this tool is designed for
- [fsnotify](https://github.com/fsnotify/fsnotify) - The file system monitoring library