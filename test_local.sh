#!/bin/bash
set -euo pipefail

# first add your user to docker group with:
# sudo usermod -aG docker $USER
# newgrp docker

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"

# Paths to p2p/proxy repos
PROXY_DIR="$ROOT_DIR/optimum-proxy"
P2P_DIR="$ROOT_DIR/optimum-p2p"

echo ">>> Building local Docker images from your branches..."

# Build proxy image
sudo docker build -t getoptimum/proxy:test \
  -f "$PROXY_DIR/deploy/Dockerfile.proxy" "$PROXY_DIR"

# Build p2pnode image
sudo docker build -t getoptimum/p2pnode:test \
  -f "$PROXY_DIR/deploy/Dockerfile.p2pnode" "$PROXY_DIR"

echo ">>> Starting docker-compose with local images..."
sudo docker-compose -f "$ROOT_DIR/optimum-dev-setup-guide/docker-compose.yml" \
               -f "$ROOT_DIR/optimum-dev-setup-guide/docker-compose.override.yml" \
               up --build "$@"

