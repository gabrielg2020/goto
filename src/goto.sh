#!/bin/bash

# Main entry point for `goto`
function goto() {
    SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

    # Source dependencies
    source "$SCRIPT_DIR/src/utils.sh"
    source "$SCRIPT_DIR/src/config.sh"
    source "$SCRIPT_DIR/src/options.sh"
    source "$SCRIPT_DIR/src/search.sh"
    source "$SCRIPT_DIR/src/fzf_integration.sh"

    # Load config
    load_default_config
    #load_config

    # Parse options and validate arguments
    parse_options "$@" || return 1
    validate_args || return 1

    # Build and execute the find command
    build_find_command
    execute_find_command

    # Configure fzf and launch it
    configure_fzf
    launch_fzf

    # Change directory to the selected directory
    if [ -n "$SELECTED_DIR" ]; then
        cd "$SELECTED_DIR" || return 1
    else
        echo "No directory selected" 
    fi
}
