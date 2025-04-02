# love2d-hot-reload

A hot-reload tool for LÖVE2D projects implemented in Go (Golang). This tool leverages Go's powerful file monitoring and process management capabilities to create an efficient development environment for LÖVE2D projects.

## Features

- Real-time file monitoring for `.lua` files and resources
- Automatic LÖVE2D process management
- Cross-platform support (primarily tested on macOS)
- Debounced reload mechanism to prevent excessive restarts
- Configurable file ignore patterns

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
go build -o love2d-hot-reload
```

## Usage

1. Place the `love2d-hot-reload` executable in your LÖVE2D project directory
2. Run the tool:
```bash
./love2d-hot-reload
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

The tool can be configured by modifying the following in `main.go`:
- Project directory path
- Ignored file patterns
- Debounce timing
- LÖVE2D executable path

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

- [LÖVE2D](https://love2d.org/) - The game framework this tool is designed for
- [fsnotify](https://github.com/fsnotify/fsnotify) - The file system monitoring library