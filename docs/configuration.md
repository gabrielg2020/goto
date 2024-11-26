# Configuration

`goto` allows you to customize its behavior using a configuration file. This file provides options to define default search settings and exclude specific directories from the results.

---

## Config File Location

The configuration file is located at: `~/.goto_config.yaml`

This file is optional. If it doesn't exist, `goto` uses sensible defaults.

---

## Config File Format

The configuration file is written in YAML format. Below is an example configuration:

```yaml
maxDepth: <integer>
excludeDirs:
  - <string>
  - <string>
```

## Available Options

|    Option     | Description                                 | Default                  |
| ------------- | ------------------------------------------- | ------------------------ |
|  `maxDepth`   | Maximum search depth.                       | `5`                      |
| `excludeDirs` | Directories to be excluded from the search. | `".git", "node_modules"` |

---

## Example Configuration 

Here's a complete example of a typical `.goto_config.yaml` file:
```yaml
maxDepth: 5
excludeDirs:
  - build
  - dist
  - .cache
```

---

[Back](index.md)