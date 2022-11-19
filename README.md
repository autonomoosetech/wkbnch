# wkbnch

Pronounced workbench, wkbnch is a tool for SchemaCAN translation and code generation.

## Installation

```bash
go get github.com/AutonoMooseTech/wkbnch
```

wkbnch will come to package managers once it reaches v1.0.0

## Use

### In the Command Line

Take manifests from an input directory `my-manifests` and generate C code in the `src` folder.
```bash
wkbnch -in=./manifests -out=./src -lang=c
```

Do nothing other than verify the manifests in the `my-manifests` directory can be parsed 
correctly. Without the `-verify` flag, wkbnch will complain of a lack of output directory.
```bash
wkbnch -in=./manifests -verify
```

If you want to see more on what you can do, run `wkbnch -h` to print out the help text.

### With a config file

(not yet implemented)

Each command line option can be configured in a YAML file. This makes the workflow for 
generating code much easier.

```yaml
# wkbnch.yaml
in: ./schema-cam-manifests/
out: ./include/
lang: c
remotes:
  - github.com/AutonoMooseTech/schema-can-tritium-solar-racing:v1.0.0
  - github.com/AutonoMooseTech/schema-can-elmar-mppt:v1.0.0
```

Running `wkbnch` in the same directory as the above file will automatically pick up the file 
because it is named either `wkbnch.yaml` or `wkbnch.yml`.
