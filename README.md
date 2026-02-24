# Grafana Ops

![License](https://img.shields.io/github/license/aristidesneto/grafana-ops)

A Go tool for backing up Grafana components (dashboards, folders, data sources,
etc.).

> ⚠️ Project under development – use with caution in production.

## Installation

The easiest way to install is to download the latest release from the GitHub
Releases page and extract the `gops` binary:

```sh
# Linux/x86_64 example
curl -LO https://github.com/aristidesneto/grafana-ops/releases/latest/download/gops_Linux_x86_64.tar.gz
tar -xzf gops_Linux_x86_64.tar.gz
chmod +x gops
mv gops /usr/local/bin/        # or another directory on your PATH
```

For other platforms, see [Releases](https://github.com/aristidesneto/grafana-ops/releases).

> 🧰 **Development alternative**: if you prefer to build locally, the
> repository includes a `Makefile` that uses `goreleaser`. Run `make deps &&
> make build` and the binary will be placed in `dist/`.

## Configuration and usage

The utility accepts options from multiple sources, in the following order of
precedence:
1. Command‑line flags
2. Environment variables
3. YAML configuration file

### Quick example – flags

```sh
gops save \
  --grafana-url https://grafana.example.com \
  --grafana-token "mytoken" \
  --output ./backup \
  --loglevel debug
```

### Using environment variables

You can configure Grafana Ops using environment variables instead of flags. All variables must be prefixed with `GO_`.

```sh
export GO_GRAFANA_URL=https://grafana.example.com
export GO_GRAFANA_TOKEN=mytoken
export GO_OUTPUT=./backup
export GO_LOGLEVEL=info
# run without flags
gops save
```

| Flag | Environment Variable | Default |
| --- | --- | --- |
| `--grafana-url` | `GO_GRAFANA_URL` | |
| `--grafana-token` | `GO_GRAFANA_TOKEN` | |
| `--output` | `GO_OUTPUT` | ./_output |
| `--loglevel` | `GO_LOGLEVEL` | info |

### Configuration file

The file is YAML and can be passed with `--config` or `-c`. Example `config.yaml`:

```yaml
# config.yaml
grafana-url: https://grafana.example.com
grafana-token: "mytoken"
output: ./backup
loglevel: info
```

```sh
gops save --config /path/to/config.yaml
```

By default the program looks for `config.yaml` in `./` and
`~/.gops` if you don’t specify a file.

## Usage examples

- Full backup using flags:
```sh
gops save --grafana-url https://grafana.example.com \
   --grafana-token mytoken \
   --output ./backup
```

## Contributing

Feel free to open pull requests, report issues or suggest improvements.

## License

This project is licensed under the **Apache License 2.0** – see the
[LICENSE](LICENSE) file for details.
