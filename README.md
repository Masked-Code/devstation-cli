# DevStation CLI

A Go CLI application that assists with spinning up a new development station for Python and C on Windows machines.

## Features

- **Automatic Environment Setup**: Installs Python, C compilers, and development tools
- **Package Manager Support**: Works with both Chocolatey and winget
- **Project Scaffolding**: Creates complete project structures for Python and C
- **Development Tools**: Installs essential development tools like Git, VS Code, CMake, etc.
- **Status Checking**: Verify what's installed on your system

## Installation

### Prerequisites

- Windows 10/11
- PowerShell 5.0 or later
- Go 1.19 or later (for building from source)

### Building from Source

1. Clone or download the project
2. Navigate to the project directory
3. Install dependencies:
   ```bash
   go mod tidy
   ```
4. Build the application:
   ```bash
   go build -o devstation.exe
   ```

## Usage

### Setup Development Environment

Set up Python development environment:
```bash
devstation setup python
```

Set up C development environment:
```bash
devstation setup c
```

Set up both Python and C environments:
```bash
devstation setup all
```

### Create New Projects

Create a new Python project:
```bash
devstation new python my-python-project
```

Create a new C project:
```bash
devstation new c my-c-project
```

### Check Environment Status

Check what's installed on your system:
```bash
devstation status
```

### Get Help

```bash
devstation --help
devstation setup --help
devstation new --help
```

## What Gets Installed

### Python Environment
- Python (latest stable version)
- pip (Python package manager)
- Essential packages: virtualenv, pip-tools, black, flake8, pytest, requests, numpy, pandas, jupyter
- Development tools: Git, VS Code

### C Environment
- C Compiler (MinGW-w64 or Visual Studio Build Tools)
- Development tools: CMake, Make, Git, VS Code, GDB, clang-format

### Package Managers
- Automatically detects and uses winget (Windows Package Manager) or Chocolatey
- Installs Chocolatey if neither is available

## Project Structure

When you create a new project, DevStation CLI creates a complete project structure:

### Python Project Structure
```
my-python-project/
├── src/
│   └── __init__.py
├── tests/
│   └── __init__.py
├── docs/
├── venv/                 # Virtual environment
├── requirements.txt      # Dependencies
├── setup.py             # Package setup
├── README.md            # Project documentation
└── .gitignore          # Git ignore rules
```

### C Project Structure
```
my-c-project/
├── src/
│   ├── main.c           # Main application
│   └── my-c-project.c   # Library source
├── include/
│   └── my-c-project.h   # Header files
├── tests/
│   └── test_main.c      # Test files
├── build/               # Build artifacts
├── docs/                # Documentation
├── CMakeLists.txt       # CMake configuration
├── Makefile            # Make configuration
├── README.md           # Project documentation
└── .gitignore         # Git ignore rules
```

## Examples

### Setting up a complete development environment:
```bash
# Install everything
devstation setup all

# Check what's installed
devstation status

# Create a new Python project
devstation new python my-web-app

# Create a new C project
devstation new c my-game-engine
```

### Building a C project:
```bash
devstation new c calculator
cd calculator
mkdir build && cd build
cmake ..
make
./calculator
```

### Setting up a Python project:
```bash
devstation new python data-analysis
cd data-analysis
venv\Scripts\activate
pip install -r requirements.txt
```

## Requirements

- Windows 10/11
- PowerShell execution policy allowing script execution
- Administrator privileges (for installing development tools)
- Internet connection (for downloading packages)

## Troubleshooting

### Package Manager Issues
If you encounter issues with package managers:
1. Run PowerShell as Administrator
2. Set execution policy: `Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope CurrentUser`
3. Try installing Chocolatey manually if winget isn't available

### Compiler Issues
If C compilation fails:
1. Ensure MinGW is in your PATH
2. Try installing Visual Studio Build Tools as an alternative
3. Restart your terminal after installation

### Python Issues
If Python packages fail to install:
1. Ensure pip is up to date: `python -m pip install --upgrade pip`
2. Try installing packages in a virtual environment
3. Check for network connectivity issues

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test thoroughly on Windows
5. Submit a pull request

## License

This project is open source and available under the MIT License.
