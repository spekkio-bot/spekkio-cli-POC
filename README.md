# Spekkio CLI
[![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/spekkio-bot/spekkio-cli?label=latest)](https://github.com/spekkio-bot/spekkio-cli/tags)
[![Go Report Card](https://goreportcard.com/badge/github.com/spekkio-bot/spekkio-cli)](https://goreportcard.com/report/github.com/spekkio-bot/spekkio-cli)

This repository contains source code for the command line tool that accesses the core Spekkio service.

## Related Repositories

Listed below are the related repositories for Spekkio:
- [Spekkio Core](https://github.com/spekkio-bot/spekkio)
- [Database Schema](https://github.com/spekkio-bot/spekkio-dbschema)

## Build and Install

Follow these steps to build and install Spekkio CLI:

1. Run `./build.sh p` to build the Spekkio CLI binary to `/usr/local/bin`.
2. Run `spekkio` on your command line!

## Available Scripts

### `./build.sh`

Running `build.sh` with no arguments will build a binary to the `bin/` directory. You can then navigate to `bin/` and run `./spekkio`.

### `./build.sh prod`, `./build.sh p`

This will build a binary to `/usr/local/bin`. You will then be able to run Spekkio globally with `spekkio`.

## License

Licensed by the GNU General Public License, version 3.
