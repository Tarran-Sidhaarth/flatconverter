[<-- Back to Main README](../README.md)

# Buffman Convert 🤖✨

Welcome to the `convert` command. This is your tool for transforming `.proto` schema files into `.fbs` using the power of FlatBuffers. You can run conversions directly through the CLI or automate them using your `buffman.yml`.

`flatbuffers` is a subcommand of `buffman convert`. In future, other formats (like `nanobuffers`) will also be supported.

## 🔧 Quick Command Reference

| Command                                                   | Description                                           |
| --------------------------------------------------------- | ----------------------------------------------------- |
| `buffman convert flatbuffers`                             | Converts `.proto` files to `.fbs` using `buffman.yml` |
| `buffman convert flatbuffers -I ./my-protos -o ./my-fbs`  | Converts using direct CLI flags                       |
| `docker run -v $(pwd):/buffman -w /buffman ghcr.io/tarran-sidhaarth/buffman convert flatbuffers --proto_dir /buffman/protos --output_dir /buffman/flatbuffers` | Docker-based conversion |

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

## 🚀 Usage Modes

You can run the convert command in two ways — CLI flags or using a configuration file.

### CLI Mode

Use this for quick conversions when you do not want to set up a config file.

```bash
buffman convert flatbuffers --proto_dir ./my-protos --output_dir ./my-fbs
```

**Docker CLI Mode**

If you're using Buffman via Docker:

```bash
docker run --rm \
    -v $(pwd):/buffman \
    -w /buffman \
    ghcr.io/tarran-sidhaarth/buffman convert flatbuffers \
    --proto_dir /buffman/protos \
    --output_dir /buffman/flatbuffers
```

> 📌 Make sure paths like `/buffman/protos` and `/buffman/flatbuffers` exist inside your mounted directory. Paths must be **relative to `/buffman`** when using Docker.

### Config Mode (Recommended)

To use `buffman.yml`, define the input and plugin configuration.  
Then run:

```bash
buffman convert flatbuffers -f ./buffman.yml
```

Or via Docker:

```bash
docker run --rm \
    -v $(pwd):/buffman \
    -w /buffman \
    ghcr.io/tarran-sidhaarth/buffman convert flatbuffers \
    -f /buffman/buffman.yml
```

## 🚩 Flags

| Flag           | Shorthand | Description                                                                     | Required |
| -------------- | --------- | ------------------------------------------------------------------------------- | -------- |
| `--proto_dir`  | `-I`      | The directory containing `.proto` files                                         | Yes      |
| `--output_dir` | `-o`      | The directory where `.fbs` files will be written. Defaults to current directory | No       |
| `-f`           | —         | Path to `buffman.yml` for config-based conversion                               | Only for config mode |

[<-- Back to Main README](../README.md)
