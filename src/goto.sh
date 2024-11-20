#!/bin/bash

function goto() {
    local OPTIND=1  # Reset OPTIND
    # Initialize default variables
    local depth=(-maxdepth 5) # Default depth
    local select_option="" # Default select option

    # Parse options
    while getopts ":d:" opt; do 
        case $opt in
            d)
                if [ "$OPTARG" = "0" ] || [ "$OPTARG" = "unlimited" ]; then
                    depth=() # Set depth to unlimited
                else
                    depth=(-maxdepth $OPTARG) # Set depth to given value
                fi
                ;;
        esac
    done
    # Remove parsed options
    shift $((OPTIND -1))
    
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

    # Setup prompt based on depth
    if [ ${#depth[@]} -eq 0 ]; then
        local prompt="Goto (Unlimited)> "
    else 
        local prompt="Goto (Depth: ${depth[@]})> "
    fi


    # Find directory and pipe to fzf
    local dir
    dir=$(find . "${depth[@]}" -type d -iname '*' -print 2>/dev/null \
        | fzf --query="$pattern" $select_option --height 40% --reverse --prompt=$prompt)
    if [ -n "$dir" ]; then
        cd "$dir" || echo "Error: Cannot change directory to $dir"
    else
        echo "No directory selected."
    fi
}
