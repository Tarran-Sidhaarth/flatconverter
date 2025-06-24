# Buffman

<p align="center">
  <img src="docs/buffman.png" alt="Buffman Logo" width="400" />
</p>

**Buffman** is a CLI tool that wraps around the `flatc` compiler. It simplifies converting `.proto` files to `.fbs`, and generates code in multiple languages using a declarative YAML config (`buffman.yml`).

> [!NOTE]
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

**Note:** Add the binary to your `PATH` for convenient use from anywhere.

## Quickstart

Make sure a file named `buffman.yml` is present in your current directory. Here's a minimal example:

```yaml
version: v1
input:
  directory: "./protos"
plugins:
  - name: flatbuffers
    out: "./fbs"
    languages:
      - language: go
        out: "./generated/go"
        opt: "github.com/username/project/fb"
```

Then run:

```bash
buffman
```

Buffman will automatically detect `buffman.yml` in the current directory. To use a custom path, use the `-f` flag:

```bash
buffman -f ./path/to/config.yml
```

## Commands

| Command    | Description                                                                                                                      |
| ---------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `buffman`  | The root command. Executes conversion and generation as defined in `buffman.yml`. Use `-f` to specify a custom config file path. |
| `convert`  | Converts `.proto` files to `.fbs` files using your `buffman.yml` settings. [Learn more](docs/convert.md)                         |
| `generate` | Generates code in multiple languages from `.fbs` files as defined in `buffman.yml`. [Learn more](docs/generate.md)               |

## Configuration

Buffman uses a YAML configuration file named `buffman.yml` to define your input directories, output locations, plugins, and language targets. Here’s the complete structure:

```yaml
version: v1

input:
  directory: "./proto"

plugins:
  - name: flatbuffers
    out: "./generated/fbs"  # Directory where .fbs files converted from .proto will be saved
    languages:
      - language: cpp
        out: "./generated/cpp"
        opt: ""
      - language: go
        out: "./generated/go"
        opt: "github.com/username/project/fb"
      - language: java
        out: "./generated/java"
        opt: "com.fb"
      - language: kotlin
        out: "./generated/kotlin"
        opt: ""
      - language: php
        out: "./generated/php"
        opt: ""
      - language: swift
        out: "./generated/swift"
        opt: ""
      - language: dart
        out: "./generated/dart"
        opt: ""
      - language: csharp
        out: "./generated/csharp"
        opt: ""
      - language: python
        out: "./generated/python"
        opt: ""
      - language: rust
        out: "./generated/rust"
        opt: ""
      - language: ts
        out: "./generated/ts"
        opt: ""
```

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
input:
  directory: "./proto"
plugins:
  - name: flatbuffers
    out: "./fbs"
    languages:
      - language: go
        out: "./generated/go"
        opt: "github.com/username/project/fb"
```

### Multi-language production example

```yaml
version: v1
input:
  directory: "./schemas"
plugins:
  - name: flatbuffers
    out: "./build/fbs"
    languages:
      - language: go
        out: "./services/go/generated"
        opt: "github.com/company/project/fb"
      - language: cpp
        out: "./native/cpp/generated"
        opt: ""
      - language: java
        out: "./services/java/generated"
        opt: "com.company.project.fb"
      - language: ts
        out: "./web/src/generated"
        opt: ""
      - language: python
        out: "./analytics/generated"
        opt: ""
```

## License

Buffman is open source under the MIT License. See [`LICENSE`](LICENSE) for full details.