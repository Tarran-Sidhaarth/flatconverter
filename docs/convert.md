[<-- Back to Main README](../README.md)

# Buffman Convert 🤖✨

Welcome to the `convert` command. This is your tool for transforming `.proto` schema files into `.fbs` using the power of FlatBuffers. You can run conversions directly through the CLI or automate them using your `buffman.yml`.

---

## 🔧 Quick Command Reference

| Command                                                                     | Description                                           |
| --------------------------------------------------------------------------- | ----------------------------------------------------- |
| `buffman convert flatbuffers`                                               | Converts `.proto` files to `.fbs` using `buffman.yml` |
| `buffman convert flatbuffers -I ./my-protos -o ./my-fbs` | Converts using direct CLI flags                       |

---

## 🧠 Full Command

```bash
buffman convert flatbuffers --proto_dir ./my-protos --output_dir ./my-fbs
```

Or using short flags:

```bash
buffman convert flatbuffers -I ./my-protos -o ./my-fbs
```

This reads all `.proto` files from the specified input directory and writes the converted `.fbs` files to the specified output directory.
If `--output_dir` is not provided, Buffman defaults to the current working directory.

---

## 🚀 Usage Modes

You can run the convert command in two ways — CLI flags or using a configuration file.

### CLI Mode

Use this for quick conversions when you do not want to set up a config file.

```bash
buffman convert flatbuffers --proto_dir ./my-protos --output_dir ./my-fbs
```

### Config Mode (Recommended)

To use `buffman.yml`, define the input directory and plugin configuration like this:

```yaml
version: v1
input:
  directory: "./schemas"
plugins:
  - name: flatbuffers
    out: "./build/fbs"
    languages:
      - language: go
        out: "./generated/go"
        opt: "github.com/example/project/fb"
```

Then run:

```bash
buffman
```

Buffman will use `buffman.yml` from the current directory. To use a custom config path:

```bash
buffman -f ./path/to/config.yml
```

---

## 🚩 Flags

| Flag           | Shorthand | Description                                                                     | Required |
| -------------- | --------- | ------------------------------------------------------------------------------- | -------- |
| `--proto_dir`  | `-I`      | The directory containing `.proto` files                                         | Yes      |
| `--output_dir` | `-o`      | The directory where `.fbs` files will be written. Defaults to current directory | No       |

---

Happy converting! 🎉

[<-- Back to Main README](../README.md)
