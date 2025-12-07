
Objective:
Create a Go command-line utility that reads template files and generates output files. 
Primary use cases include generating Dockerfiles and Terraform files.

Functional Requirements:
- Read a template file and produce an output file.
- Optionally substitute values from a configuration file (YAML or
JSON).
- if --config is not specified, then do not read a config file
- Support template engines: Jinja, Mustache, Handlebars, Liquid, Go text/template, and m4.
  - For m4, invoke the system preprocessor.
- use popular and maintained libraries
- Detect template format from file extension or via --format option.
- extensions:
    Jinja: .j2, .jinja2
    Mustache: .mustache
    Handlebars: .hbs
    Liquid: .liquid
    Go text/template: .tmpl, .tpl
    m4: .m4

- Allow multiple input files; process them concurrently using N worker
routines.
- default workers: 1
- Output filename = input filename with template extension removed (e.g., `filename.j2` to `filename`).
- Support stdin to stdout processing mode.
- overwrite output file if --overwrite is specified
- logging
  --quiet: No output at all, except for fatal errors?
  --verbose: Standard level of informational output?
  --debug: Very detailed output, perhaps including the data being passed
  to templates?
  default: log only errors


Command-Line Options:
--debug, --verbose, --quiet, --config <file>, --workers <n>, --stdin,
--format <engine>, --overwrite

Non-Functional Requirements:
- The name of the project is outgen
- The cli command SHALL be called outgen
- Include structured logging and robust error handling.
- Organize code into reusable functions and packages.
- Add comment headers to each source file with metadata:
  Filename:
  Author: Michael Moscovitch
  Assistant: Jules
  Date: 2025/12/06
  Copyright (c) 2025 Michael Moscovitch
  Description:

Deliverables:
- Source code with headers
- README.md (usage, installation, examples)
- agents.md (design notes, architecture, contributors)
- Makefile for building
- Unit tests with clear coverage of core functionality
