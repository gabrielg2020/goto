#!/bin/bash

# fzf integration for `goto`

function configure_fzf() {
  FZF_OPTIONS=()
  
  # If pattern is >= 3, then select the first match
  if [ ${#PATTERN} -ge 3 ]; then
    FZF_OPTIONS+=("--select-1")
  fi
  FZF_OPTIONS+=("--height" "40%" "--reverse")

  # Setup prompt based on depth
  if [ ${#DEPTH[@]} -eq 0 ]; then
    FZF_PROMPT="Goto(Depth: unlimited)> "
  else
    FZF_PROMPT="Goto(Depth: ${DEPTH[2]})> "
  fi
  FZF_OPTIONS+=("--prompt" "${FZF_PROMPT}")
}

function launch_fzf() {
  # Launch fzf with the find results
  SELECTED_DIR=$(printf '%s\n' "${FIND_RESULTS[@]}" \
        | fzf --query="$PATTERN" "${FZF_OPTIONS[@]}")
}