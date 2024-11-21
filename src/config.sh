#!/bin/bash

# Configuration management for `goto`

function load_default_config() {
  DEPTH=(-maxdepth 5) # Default depth
  EXCLUDE_PATHS=( # Default exclude directories
    "node_modules"
    ".git"
    ".github"
  )
}

function load_config() {
  # Load default config
  load_default_config

  local config_file="$HOME/.goto.conf"
    if [ -f "$config_file" ]; then
        source "$config_file"
    fi
}