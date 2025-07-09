# DevStation CLI Installation Guide

## Quick Start

1. **Download or Build**
   - Download the latest `devstation.exe` from releases, or
   - Build from source using the instructions below

2. **Run the CLI**
   ```powershell
   .\devstation.exe --help
   ```

3. **Check Your Environment**
   ```powershell
   .\devstation.exe status
   ```

4. **Set Up Development Environment**
   ```powershell
   # For Python development
   .\devstation.exe setup python
   
   # For C development
   .\devstation.exe setup c
   
   # For both
   .\devstation.exe setup all
   ```

## Building from Source

### Prerequisites
- Windows 10/11
- Go 1.19 or later
- PowerShell 5.0 or later

### Build Steps

1. **Clone or Download the Project**
   ```powershell
   # If you have Git installed
   git clone https://github.com/yourusername/devstation-cli.git
   cd devstation-cli
   
   # Or download and extract the ZIP file
   ```

2. **Install Dependencies**
   ```powershell
   go mod tidy
   ```

3. **Build the Application**
   ```powershell
   # Using the build script (recommended)
   .\build.ps1
   
   # Or manually
   go build -o devstation.exe
   ```

4. **Test the Build**
   ```powershell
   .\devstation.exe --help
   ```

## Installation Options

### Option 1: Standalone Executable (Recommended)
- Copy `devstation.exe` to any folder
- Run it from that location
- No additional installation required

### Option 2: Add to PATH
1. Copy `devstation.exe` to a permanent location (e.g., `C:\Tools\devstation.exe`)
2. Add that directory to your PATH environment variable
3. Open a new PowerShell window
4. Run `devstation --help` from anywhere

### Option 3: PowerShell Profile
Add an alias to your PowerShell profile:
```powershell
# Add this to your PowerShell profile
Set-Alias -Name devstation -Value "C:\Path\To\devstation.exe"
```

## First Run

After installation, run these commands to get started:

```powershell
# Check what's already installed
devstation status

# Set up Python development environment
devstation setup python

# Set up C development environment  
devstation setup c

# Create a new Python project
devstation new python my-project

# Create a new C project
devstation new c my-c-project
```

## Common Issues

### PowerShell Execution Policy
If you get execution policy errors:
```powershell
Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope CurrentUser
```

### Missing Package Managers
The CLI will automatically install Chocolatey if neither winget nor Chocolatey is available.

### Administrator Rights
Some operations (like installing development tools) require administrator privileges. Run PowerShell as Administrator when using setup commands.

## System Requirements

- **OS**: Windows 10/11
- **PowerShell**: 5.0 or later
- **Disk Space**: ~50MB for the CLI, additional space for development tools
- **Network**: Internet connection required for downloading packages
- **Permissions**: Administrator rights for installing development tools

## Uninstallation

To remove DevStation CLI:
1. Delete the `devstation.exe` file
2. Remove any PATH entries (if added)
3. Remove any PowerShell aliases (if added)

Note: This won't uninstall the development tools that were installed. Use your package manager (winget/chocolatey) to remove those if needed.

## Getting Help

```powershell
# General help
devstation --help

# Command-specific help
devstation setup --help
devstation new --help
devstation status --help
```

## Next Steps

After installation, check out the main README.md for usage examples and detailed documentation.
