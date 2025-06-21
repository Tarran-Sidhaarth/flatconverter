[<-- Back to Main README](../README.md)

# Buffman Convert 🤖✨

Welcome to the `convert` command's headquarters! This is your go-to tool for transforming schema files from one format to another. Right now, its superpower is turning Protocol Buffers (`.proto`) into FlatBuffers (`.fbs`).

---

## 🧠 The `convert flatbuffers` Command

At its core, `convert` has one super-useful subcommand: `flatbuffers`.

```bash
buffman convert flatbuffers [FLAGS]
```

This command reads all the `.proto` files from a source directory and spits out shiny new `.fbs` files in an output directory.

---

## 🚀 Usage Examples

You can run conversions in two ways: directly from the command line or using your `buffman.yml` for automated runs.

### 1. 🖥️ Command-Line Interface (CLI)

Perfect for quick conversions without touching your config file.

**Example:** Convert protos in `my-protos/` to fbs files in `my-fbs/`

```bash
buffman convert flatbuffers --proto_dir ./my-protos --output_dir ./my-fbs
```

---

### 2. ⚙️ Using `buffman.yml`

This is the recommended way to manage your project's workflow. Define your conversion path in `buffman.yml`, and then just run the main `buffman` command.

**Example `buffman.yml`:**

```yaml
convert:
  proto_dir: ./my-protos
  flatbuffers:
    output_dir: ./my-fbs
```

Then, simply run:

```bash
buffman
```

Voila! Buffman reads your config and handles the rest.

---

## 🚩 Flags

Here are the flags you can use with `buffman convert flatbuffers`:

| Flag             | Shorthand | Description                                             | Required? |
|------------------|-----------|---------------------------------------------------------|-----------|
| `--proto_dir`    | `-I`      | The directory where your source `.proto` files live.    | ✅ Yes     |
| `--output_dir`   | `-o`      | The directory where the generated `.fbs` files will go. | ❌ No      |

---

## 🛠️ Configuration Keys

You can set up conversion paths using the following keys inside your `buffman.yml`:

| Key                              | Type   | Description                                                   |
|----------------------------------|--------|---------------------------------------------------------------|
| `convert.proto_dir`              | string | **Required.** The path to your directory of `.proto` files.   |
| `convert.flatbuffers.output_dir` | string | Output directory for `.fbs` files. Defaults to `./`.          |

---

Happy converting! 🎉

[<-- Back to Main README](../README.md)
