#!/bin/bash

# Main entry point for `goto`
function goto() {
    SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

    # Source dependencies
    source "$SCRIPT_DIR/src/utils.sh"
    source "$SCRIPT_DIR/src/config.sh"
    source "$SCRIPT_DIR/src/options.sh"

    # Load config
    load_config

    # Parse options and validate arguments
    parse_options "$@" || return 1
    validate_args || return 1

    echo "Depth: ${DEPTH[@]}"
    echo "Excluded paths: ${EXCLUDE_PATHS[@]}"
}
