#!/bin/sh
set -e
echo "📦 Pulling model mistral..."
ollama pull mistral || true
echo "🚀 Starting ollama..."
exec ollama serve
