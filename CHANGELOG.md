# Changelog
The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.2.0] - 2022-07-31
### Added
- `ExecCommandWithStdin`
### Changes
- The `Mount` function now has a password parameter.
### Security
- `Mount` now uses `ExecCommandWithStdin` to protect against password leaks. See PR [#2](https://github.com/mark-hartmann/vera/pull/2) for a detailed writeup of the problem.
## [0.1.0] - 2021-10-07
Initial release
### Added
- Basic functions and structs for communicating with the veracrypt cli.
- Advanced functions that simplify the use of the most common commands:
  - `List`
  - `DismountAll`
  - `DismountSlot`
  - `DismountVolume`
  - `PropertiesSlot`
  - `PropertiesVolume`
  - `Installed`
  - `Mount`
- Basic tests, which cover basic use cases. This is definitely something that needs and will improve in further
  versions.

[Unreleased]: https://github.com/mark-hartmann/vera/compare/v0.2.0...HEAD
[0.2.0]: https://github.com/mark-hartmann/vera/releases/tag/v0.2.0
[0.1.0]: https://github.com/mark-hartmann/vera/releases/tag/v0.1.0