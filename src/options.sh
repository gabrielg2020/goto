#!/bin/bash

# Option parsing for `goto`

function parse_options() {
  local OPTIND
  OPTIND=1
  while getopts ":d:" opt; do
    case $opt in
      d)
        if [[ ! "$OPTARG" =~ ^[0-9]+$ ]]; then # If argument is not a number
          echo "Invalid depth: $OPTARG" >&2
          return 1
        elif [[ "$OPTARG" == "0" || "$OPTARG" == "unlimited" ]]; then
          DEPTH=() # Unlimited depth
        else
          DEPTH=(-maxdepth "$OPTARG")
        fi
        ;;
      \?)
        echo "Invalid option: -$OPTARG" >&2
        print_usage
        return 1
        ;;
      :)
        echo "Option -$OPTARG requires an argument" >&2
        print_usage
        return 1
        ;;
    esac
  done
  shift $((OPTIND - 1))

  # Assign the remaining argument as the pattern
  PATTERN="$1"
}

function validate_args() {
  if [[ -z "$PATTERN" ]]; then
    echo "Missing directory name pattern" >&2
    print_usage
    return 1
  fi
}