#!/bin/bash

# Main entry point for `goto`
function goto() {
    SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

    # Source dependencies
    source "$SCRIPT_DIR/src/config.sh"

    # Load config
    load_config

    echo "Depth: ${DEPTH[@]}"
    echo "Excluded paths: ${EXCLUDE_PATHS[@]}"
}
