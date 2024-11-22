#!/bin/bash

# Search function for `goto`

function build_find_command() {
  FIND_EXCLUDED=()
  if [ ${#EXCLUDE_PATHS[@]} -ne 0 ]; then
    for excluded_path in "${EXCLUDE_PATHS[@]}"; do
      FIND_EXCLUDED+=(-path "./$excluded_path" -prune -o)
    done
  fi

  # Construct the find command
  FIND_COMMAND=(find . "${DEPTH[@]}" "${FIND_EXCLUDED[@]}" -type d -iname '*' -print)
}

function execute_find_command() {
  # Execute the find command
  FIND_RESULTS=()
  while IFS= read -r line; do
    FIND_RESULTS+=("$line")
  done < <("${FIND_COMMAND[@]}" 2>/dev/null)
}