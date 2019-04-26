# go-api-docs
GoLang API Documentation Generator

# Instructions

Fork it or download this project as a zip file.

## Installation

```sh
. ./scripts/setpath.sh
```

## Create Your Docs

Modify files found in src/service/apidoc to configure the documentation.

## Build The Server

You may need to call ./scripts/setpath.sh to set your paths, but you should only have to do this once per shell session.

```sh
. ./scripts/setpath.sh
. ./scripts/generate.sh
```

This will generate a "service" executable.  By default, running this executable will make a server available at http://localhost:8080.
