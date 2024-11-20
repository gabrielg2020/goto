#!/bin/bash

function goto() {
    # Initialize default variables
    local depth="-maxdepth 5" # Default depth
    local select_option="" # Default select option
    
    # Check if pattern is given
    if [ -z "$1" ]; then
        echo "Usage: goto <directory_name_pattern>"
        return 1
    fi

    # Setup select option based on pattern length
    local pattern="$1"
    if [ ${#pattern} -ge 3 ]; then
        select_option="--select-1"
    fi

    # Find directory and pipe to fzf
    local dir
    dir=$(find . -type d -iname '*' -print 2>/dev/null \
        | fzf --query="$pattern" $select_option --height 40% --reverse --prompt="Goto> ")
    if [ -n "$dir" ]; then
        cd "$dir" || echo "Error: Cannot change directory to $dir"
    else
        echo "No directory selected."
    fi
}
