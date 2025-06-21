[<-- Back to Main README](../README.md)

# Buffman Generate 🚀🔥

Ready to turn your schemas into actual, usable code? The `generate` command is here to do the heavy lifting! It takes your schema files (like `.fbs`) and generates language-specific source code so you can start building.

---

## 🛠️ The `generate flatbuffers` Command

This is where the magic happens! The `generate flatbuffers` subcommand invokes the `flatc` compiler to create source files from your FlatBuffer schemas.

```bash
buffman generate flatbuffers [FLAGS]
```

---

## 🚀 Usage Examples

You can generate code for a single language via the CLI or—even better—configure a multi-language generation pipeline in your `buffman.yml`.

### 1. 🖥️ Command-Line Interface (CLI)

Great for generating code for one language on the fly.

**Example:** Generate Go code from `.fbs` files in `my-fbs/`

```bash
buffman generate flatbuffers \
  --flatbuffers_dir=./my-fbs \
  --language=go \
  --target_dir=./gen/go \
  --module_options="github.com/your-org/your-project/gen/go"
```

---

### 2. ⚙️ Using `buffman.yml` (The Power-User Way 💪)

This is where Buffman truly shines. Define all your target languages in `buffman.yml`, and Buffman will generate them all with a single command.

**Example `buffman.yml`:**

```yaml
generate:
  flatbuffers:
    input_dir: ./my-fbs
    languages:
      go:
        output_dir: ./gen/go
        module_options: "github.com/your-org/your-project/gen/go"
      cpp:
        output_dir: ./gen/cpp
      java:
        output_dir: ./gen/java
        module_options: "com.your-org.your-project"
```

Then just run:

```bash
buffman
```

Buffman will read your config and create the Go, C++, and Java files in their respective output directories. How cool is that?

---

## 🚩 Flags

Here are the flags available for `buffman generate flatbuffers`:

| Flag                | Shorthand | Description                                                                 | Required? |
|---------------------|-----------|-----------------------------------------------------------------------------|-----------|
| `--flatbuffers_dir` | `-I`      | The directory where your source `.fbs` schema files are.                    | ✅ Yes     |
| `--language`        | `-l`      | The target language to generate (e.g., `go`, `java`, `cpp`, `kotlin`).      | ✅ Yes     |
| `--target_dir`      | `-o`      | The output directory for the generated code.                                | ❌ No      |
| `--module_options`  | `-m`      | Language-specific options, like a Go module path or Java package prefix.    | ❌ No      |

---

## 🧩 Configuration Keys

Here are the keys you can use in the `generate` section of your `buffman.yml`:

| Key                                         | Type   | Description                                                                                           |
|---------------------------------------------|--------|-------------------------------------------------------------------------------------------------------|
| `generate.flatbuffers.input_dir`            | string | **Required.** The path to your source `.fbs` files.                                                   |
| `generate.flatbuffers.languages`            | object | A container for all target language configurations.                                                   |
| `...languages.<lang>.output_dir`            | string | **Required.** The output directory for this specific language (e.g., `languages.go.output_dir`).      |
| `...languages.<lang>.module_options`        | string | Optional. Language-specific options like a Go package path or Java package name for that language.    |

---

Happy generating! 💻

[<-- Back to Main README](../README.md)
