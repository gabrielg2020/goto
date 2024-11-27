# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/), and this project adheres to [Semantic Versioning](https://semver.org/).

---
## [1.0.0] - 2024-11-27
### Added

- Search directories for a specific `<pattern>`.
- Define max search depth with `-d` flag.
- Display usage/help with `goto -h`.
- Support for a configuration file called `.goto_config.yaml`.
  - Change default max depth with key `maxDepth`.
  - Change excluded directories with key `excludeDirs`.
- Documentation with GitHub pages: [Documentation](gabrielg2020.github.io/goto).
- CI/CD Workflows to test, build, attach build artifacts to releases.

---

[Back](index.md)