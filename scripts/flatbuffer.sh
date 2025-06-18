#!/bin/sh

# Ensure the script is run as root (required for install)
if [ "$(id -u)" -ne 0 ]; then
  echo "Please run this script with sudo:"
  echo "  sudo $0"
  exit 1
fi

# Check if flatc exists and works
if command -v flatc >/dev/null 2>&1; then
  echo "✅ flatc is already installed at $(command -v flatc)"
  exit 0
fi

# Check for cmake and make
for cmd in cmake make; do
  if ! command -v "$cmd" >/dev/null 2>&1; then
    echo "❌ '$cmd' not found. Please install it before running this script."
    exit 1
  fi
done

# Create a temp directory for the build
TMPDIR=$(mktemp -d)
echo "🔧 Building FlatBuffers in $TMPDIR"
cd "$TMPDIR" || exit 1

# Clone and build FlatBuffers
git clone https://github.com/google/flatbuffers.git || exit 1
cd flatbuffers || exit 1

cmake -G "Unix Makefiles" . || exit 1
make || exit 1

# Run quick test
echo "🧪 Running tests..."
if ./flattests | grep -q "ALL TESTS PASSED"; then
  echo "✅ Tests passed"
else
  echo "❌ Tests failed"
  exit 1
fi

# Install
make install || exit 1

# Copy flatc to /usr/local/bin and make it executable
if [ -f flatc ]; then
  cp flatc /usr/local/bin/flatc
  chmod +x /usr/local/bin/flatc
  echo "✅ flatc installed to /usr/local/bin/flatc"
else
  echo "❌ flatc binary not found"
  exit 1
fi

# Cleanup
echo "🧹 Cleaning up..."
rm -rf "$TMPDIR"

echo "🎉 flatc installation complete!"
