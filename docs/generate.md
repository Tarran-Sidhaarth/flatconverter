[<-- Back to Main README](../README.md)

# Buffman Generate 🚀🔥

The `generate` command turns your FlatBuffer schemas (`.fbs`) into language-specific source code using the `flatc` compiler. Run it directly from the CLI for one-off jobs or configure every language you need in `buffman.yml` and let Buffman handle the rest.

## 🔧 Quick Command Reference

| Command                                                                                                                                   | Description                                    |
| ----------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------- |
| `buffman`                                                                                                            | Generates code using settings in `buffman.yml` |
| `buffman generate flatbuffers --I ./my-fbs --language go -o ./gen/go --module_options "github.com/me/project/fb"` | Generates code through CLI flags               |

## 🧠 Full Command

```bash
buffman generate flatbuffers -I ./my-fbs --language go -o ./gen/go --module_options "github.com/me/project/fb"
```
If `-o` is omitted Buffman writes generated code to the current working directory.

## 🚀 Usage Modes

### CLI Mode

Best for quick, single-language generation.

```bash
buffman generate flatbuffers --flatbuffers_dir ./my-fbs --language cpp --target_dir ./gen/cpp
```

### Config Mode (Recommended)

Define every language once and generate them all with one command.

```yaml
version: v1
input:
  directory: "./schemas"           # Where your .proto files live
plugins:
  - name: flatbuffers
    out: "./build/fbs"             # Where the .fbs files were written by convert
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

Run everything with:

```bash
buffman
```

To use a different configuration file:

```bash
buffman -f ./path/to/config.yml
```

## 🚩 Flags

| Flag                | Shorthand | Description                                                             | Required |
| ------------------- | --------- | ----------------------------------------------------------------------- | -------- |
| `--flatbuffers_dir` | `-I`      | Directory containing source `.fbs` files                                | Yes      |
| `--language`        | `-l`      | Target language to generate (`go`, `cpp`, `java`, `kotlin`, etc.)       | Yes      |
| `--target_dir`      | `-o`      | Directory to write generated code. Defaults to current directory        | No       |
| `--module_options`  | `-m`      | Language-specific options such as Go module path or Java package prefix | No       |

Happy generating! 💻

[<-- Back to Main README](../README.md)
