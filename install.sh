#!/bin/bash

# Crypto Price CLI Installation Script

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Configuration
BINARY_NAME="crypto"
INSTALL_DIR="/usr/local/bin"

echo -e "${BLUE}🚀 Installing Crypto Price CLI...${NC}"
echo

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo -e "${RED}❌ Go is not installed. Please install Go first.${NC}"
    echo -e "${YELLOW}💡 Visit: https://golang.org/doc/install${NC}"
    exit 1
fi

echo -e "${GREEN}✅ Go is installed${NC}"

# Check Go version
GO_VERSION=$(go version | cut -d' ' -f3 | tr -d 'go')
REQUIRED_VERSION="1.19"

if ! printf '%s\n%s\n' "$REQUIRED_VERSION" "$GO_VERSION" | sort -V -C; then
    echo -e "${RED}❌ Go version $GO_VERSION is too old. Minimum required: $REQUIRED_VERSION${NC}"
    exit 1
fi

echo -e "${GREEN}✅ Go version $GO_VERSION is compatible${NC}"

# Download dependencies
echo -e "${YELLOW}📦 Downloading dependencies...${NC}"
go mod tidy

# Build the binary
echo -e "${YELLOW}🔨 Building binary...${NC}"
go build -o $BINARY_NAME -ldflags="-s -w" .

if [ $? -eq 0 ]; then
    echo -e "${GREEN}✅ Binary built successfully${NC}"
else
    echo -e "${RED}❌ Failed to build binary${NC}"
    exit 1
fi

# Check if install directory exists and is writable
if [ ! -d "$INSTALL_DIR" ]; then
    echo -e "${YELLOW}📁 Creating install directory: $INSTALL_DIR${NC}"
    sudo mkdir -p "$INSTALL_DIR"
fi

# Install the binary
echo -e "${YELLOW}📝 Installing to $INSTALL_DIR...${NC}"
sudo mv $BINARY_NAME $INSTALL_DIR/

# Make it executable
sudo chmod +x $INSTALL_DIR/$BINARY_NAME

# Verify installation
if command -v $BINARY_NAME &> /dev/null; then
    echo
    echo -e "${GREEN}🎉 Installation completed successfully!${NC}"
    echo
    echo -e "${BLUE}Usage examples:${NC}"
    echo "  $BINARY_NAME price bitcoin"
    echo "  $BINARY_NAME price btc eur"
    echo "  $BINARY_NAME list"
    echo "  $BINARY_NAME list --limit 20"
    echo
    echo -e "${YELLOW}💡 Run '$BINARY_NAME --help' for more information${NC}"
else
    echo -e "${RED}❌ Installation failed. Binary not found in PATH${NC}"
    exit 1
fi 