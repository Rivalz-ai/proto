# Script to build and publish the protobuf TypeScript package
Write-Host "Building and publishing protobuf TypeScript package..." -ForegroundColor Green

# Navigate to the typescript directory
Set-Location $PSScriptRoot

# Install dependencies
Write-Host "Installing dependencies..." -ForegroundColor Yellow
npm install

# Build the package
Write-Host "Building package..." -ForegroundColor Yellow
npm run build

# Check if build was successful
if ($LASTEXITCODE -ne 0) {
    Write-Host "Build failed!" -ForegroundColor Red
    exit 1
}

# Ask for confirmation before publishing
Write-Host "Build successful!" -ForegroundColor Green
Write-Host "Ready to publish package to npm." -ForegroundColor Yellow
$confirm = Read-Host "Do you want to publish to npm? (y/N)"

if ($confirm -eq "y" -or $confirm -eq "Y") {
    Write-Host "Publishing to npm..." -ForegroundColor Yellow
    npm publish
    
    if ($LASTEXITCODE -eq 0) {
        Write-Host "Package published successfully!" -ForegroundColor Green
    } else {
        Write-Host "Failed to publish package!" -ForegroundColor Red
    }
} else {
    Write-Host "Publishing cancelled." -ForegroundColor Yellow
} 