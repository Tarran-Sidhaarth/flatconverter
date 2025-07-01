# Buffman

<p align="center">
  <img src="docs/buffman.png" alt="Buffman Logo" width="400" />
</p>

**Buffman** is a CLI tool that wraps around the `flatc` compiler. It simplifies converting `.proto` files to `.fbs`, and generates code in multiple languages using a declarative YAML config (`buffman.yml`).

> [!IMPORTANT]  
> This project is under active development. APIs, configurations, and features may change without notice. Use with caution in production environments.

- [Buffman](#buffman)
   * [Installation](#installation)
   * [Quickstart](#quickstart)
   * [Commands](#commands)
   * [Configuration](#configuration)
   * [Supported Languages](#supported-languages)
   * [Examples](#examples)
      + [Minimal example](#minimal-example)
      + [Multi-language production example](#multi-language-production-example)
   * [License](#license)

## Installation

You can install Buffman in two ways:

1. **Download Precompiled Binary**  
   Visit the [Releases page](releases/) and download the binary for your OS.

   ```bash
   export BUFFMAN_VERSION="1.0.0" && \
   curl -L "https://github.com/machanirobotics/buffman/releases/download/v$BUFFMAN_VERSION/buffman-linux-x86-64-$BUFFMAN_VERSION" -o buffman && \
   sudo mv buffman /usr/local/bin/ && \
   sudo chmod +x /usr/local/bin/buffman
   ```

2. **Build from Source**

   ```bash
   git clone https://github.com/your-org/buffman.git
   cd buffman
   go build -o buffman main.go
   ```
> [!TIP]
> Add the binary to your `PATH` for convenient use from anywhere.


## Quickstart

Buffman requires a YAML configuration file and **does not** auto-detect it.  
You **must specify the file explicitly** using the `-f` flag.

Here's a minimal example config (`buffman.yml`):

```yaml
version: v1

inputs:
  - name: source
    path: "./proto"

  - name: googleprotobuf
    remote: https://github.com/protocolbuffers/protobuf
    commit: <commit-hash>

plugins:
  - name: flatbuffers
    out: "./fbs"
    languages:
      - language: go
        out: "./generated/go"
        opt:
          - go_package=github.com/username/project/fb
```

Then run:

```bash
buffman generate -f ./buffman.yml
```

You can use any filename and location for the config—just update the path with `-f`.


## Commands

| Command             | Description                                                                                               |
|---------------------|-----------------------------------------------------------------------------------------------------------|
| `buffman generate`  | Generates code as defined in your config file. Use the `-f` flag to specify the config path.              |
| `buffman convert`   | Converts `.proto` files to `.fbs` files using your config. [Learn more](docs/convert.md)                  |



## Configuration

Buffman uses a YAML configuration file (`buffman.yml`) to define your input sources, output directories, plugins, and language-specific options.

### Structure

```yaml
version: v1

inputs:
  - name: source
    path: "./proto"

  # Optional external repositories
  # - name: googleprotobuf
  #   remote: https://github.com/protocolbuffers/protobuf
  #   commit: <commit-hash>

plugins:
  - name: flatbuffers
    out: "./fbs"
    languages:
      - language: cpp
        out: "./generated/cpp"

      - language: go
        out: "./generated/go"
        opt:
          - go_package=github.com/username/project/fb

      - language: java
        out: "./generated/java"
        opt:
          - java_package_prefix=com.fb

      - language: kotlin
        out: "./generated/kotlin"

      - language: php
        out: "./generated/php"

      - language: swift
        out: "./generated/swift"

      - language: dart
        out: "./generated/dart"

      - language: csharp
        out: "./generated/csharp"

      - language: python
        out: "./generated/python"

      - language: rust
        out: "./generated/rust"

      - language: ts
        out: "./generated/ts"
```

- `inputs` define your schema sources.
- `plugins` define how `.proto` files are converted and which language targets to generate.
- `opt` is required only for `go` (`go_package`) and `java` (`java_package_prefix`).

## Supported Languages

The following languages are currently supported for code generation via FlatBuffers:

* `cpp`
* `go`
* `java`
* `kotlin`
* `php`
* `swift`
* `dart`
* `csharp`
* `python`
* `rust`
* `ts`

## Examples

### Minimal example

```yaml
version: v1
inputs:
  - name: source
    path: "./proto"

plugins:
  - name: flatbuffers
    out: "./fbs"
    languages:
      - language: go
        out: "./generated/go"
        opt:
          - go_package=github.com/username/project/fb
```

### Multi-language production example

```yaml
version: v1
inputs:
  - name: source
    path: "./schemas"

plugins:
  - name: flatbuffers
    out: "./build/fbs"
    languages:
      - language: go
        out: "./services/go/generated"
        opt:
          - go_package=github.com/company/project/fb

      - language: cpp
        out: "./native/cpp/generated"

      - language: java
        out: "./services/java/generated"
        opt:
          - java_package_prefix=com.company.project.fb

      - language: ts
        out: "./web/src/generated"

      - language: python
        out: "./analytics/generated"
```

## License
Buffman is open source under the MIT License. See `LICENSE` for full details.

> [!NOTE]
> For full documentation and advanced usage, [read the DOCS](./docs/).
