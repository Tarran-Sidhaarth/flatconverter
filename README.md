# Buffman

<p align="center">
  <img src="docs/buffman.png" alt="Buffman Logo" width="200" />
</p>

> [!WARNING]
> **Heads up! Buffman is evolving fast 🚀**  
> This project is under active development. APIs, configurations, and features are still being polished and may change without notice. Use with caution in production environments and stay tuned for updates!

---

## Table of Contents

- [What is Buffman?](#what-is-buffman)
- [Getting Started](#getting-started)
- [Buffman Commands](#buffman-commands)
  - [buffman (root command)](#buffman-root-command)
  - [convert](#convert)
  - [generate](#generate)
- [Configuration](#configuration)
- [Contributing](#contributing)
- [License](#license)

---

## What is Buffman?

Buffman is your friendly neighborhood CLI tool that wraps around the mighty `flatc` compiler. It makes converting Protocol Buffer (`.proto`) files into FlatBuffers (`.fbs`) a breeze — no more wrestling with complicated commands! Plus, with a powerful YAML config (`buffman.yml`), you can automate conversions and code generation effortlessly.

And guess what? We’re just getting started! Soon, Buffman will support more buffer formats like Nanobuffers and beyond. Stay tuned!

---

## Getting Started

Install Buffman, create your `buffman.yml` config, and you're ready to roll. Whether you want to convert proto files or generate language-specific code, Buffman’s got your back.

---

## Buffman Commands

### buffman (root command)

The heart of Buffman! Run this command with your config file (default: `buffman.yml`) and it will execute all the conversions and generations you’ve defined. No need to run multiple commands manually — just one command to rule them all!


### convert

Want to transform your `.proto` files into `.fbs`? The `convert` command is your go-to. It supports converting proto files to FlatBuffers schemas with ease.

For a deep dive into usage, flags, and examples, check out the detailed guide:  
[docs/convert.md](docs/convert.md)

### generate

Ready to turn your `.fbs` schemas into language-specific source files? The `generate` command lets you generate code in Go, C++, Java, Kotlin, and more — all configurable via your YAML file.

Explore all the nitty-gritty details here:  
[docs/generate.md](docs/generate.md)

---

## Configuration

Buffman uses a YAML configuration file (`buffman.yml`) to define your conversion and generation workflows. This makes it easy to automate complex pipelines without breaking a sweat.

---

## Contributing

We ❤️ contributions! Whether it’s bug reports, feature requests, or code improvements, your help makes Buffman better. Check out [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

---

## License

Buffman is open source under the MIT License. See [LICENSE](LICENSE) for details.

---

Happy buffering! 🎉
