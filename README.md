# outgen

`outgen` is a command-line utility that reads template files and generates output files. It's designed to be a flexible tool for code generation, supporting various template engines and value substitution from configuration files.

## Installation

To install `outgen`, you'll need to have Go installed on your system. You can then build it from source:

```bash
git clone https://github.com/infologicmgmt/outgen.git
cd outgen
make build
```

This will create an `outgen` binary in the project's root directory.

## Usage

```
outgen [flags] [input_file1] [input_file2] ...
```

### Flags

- `--config, -c <file>`: Configuration file (YAML or JSON) for value substitution.
- `--format, -f <engine>`: Template format (e.g., `jinja`, `mustache`, `go`). If not specified, it's detected from the file extension.
- `--workers, -w <n>`: Number of concurrent workers (default: 1).
- `--stdin`: Read from stdin and write to stdout.
- `--overwrite`: Overwrite existing output files.
- `--debug`: Enable debug logging.
- `--verbose`: Enable verbose logging.
- `--quiet`: Suppress all output except for fatal errors.

### Examples

**Basic Usage**

Given a template file `Dockerfile.j2`:

```jinja
FROM {{ base_image }}
RUN echo "Hello, {{ name }}!"
```

And a configuration file `config.yaml`:

```yaml
base_image: alpine:latest
name: outgen
```

You can generate the `Dockerfile` with:

```bash
outgen --config config.yaml Dockerfile.j2
```

**Using stdin and stdout**

```bash
echo "Hello, {{ .name }}" | outgen --stdin --format go --config '{"name": "World"}'
```

## Supported Template Engines

- Jinja (`.j2`, `.jinja2`)
- Mustache (`.mustache`)
- Handlebars (`.hbs`)
- Liquid (`.liquid`)
- Go `text/template` (`.tmpl`, `.tpl`)
- m4 (`.m4`)
