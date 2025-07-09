# Build script for DevStation CLI

Write-Host "Building DevStation CLI..." -ForegroundColor Green

# Clean up previous builds
if (Test-Path "devstation.exe") {
    Remove-Item "devstation.exe" -Force
    Write-Host "Removed previous build" -ForegroundColor Yellow
}

# Build the application
Write-Host "Running go build..." -ForegroundColor Cyan
go build -o devstation.exe

if ($LASTEXITCODE -eq 0) {
    Write-Host "✓ Build successful!" -ForegroundColor Green
    Write-Host "Executable created: devstation.exe" -ForegroundColor Green
    
    # Show basic info about the build
    $fileInfo = Get-Item "devstation.exe"
    Write-Host "File size: $([math]::Round($fileInfo.Length / 1MB, 2)) MB" -ForegroundColor Cyan
    Write-Host "Build time: $(Get-Date)" -ForegroundColor Cyan
    
    # Test basic functionality
    Write-Host "`nTesting basic functionality..." -ForegroundColor Cyan
    ./devstation.exe --help | Select-Object -First 5
    
    Write-Host "`n✓ DevStation CLI is ready to use!" -ForegroundColor Green
    Write-Host "Run './devstation.exe --help' for usage information" -ForegroundColor Yellow
} else {
    Write-Host "✗ Build failed!" -ForegroundColor Red
    exit 1
}
