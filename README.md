# itemGenerator
A simple item generator with fake data that can be configured

# Building Locally
```
./bin/dev/build.sh
```

# Usage
```
./build/generate --rows=100 --template=templates/simple.json
```

# Templates
Templates are used to configure the generator.
```
"columns": [
    {"name": "id", "type": "integer"},
    {"name": "name", "type": "string"},
    {"name": "price", "type": "float"}
]
```
**output_name**: The output file, minus the extensions. Saved to `./generated/`

**compression**: File compression
```
none|zip|gzip
```

**delimiter** - The column delimiter
```
comma|tab
```

**columns** - A map of columns
    : *name* - The column name used in the file
    : *type* - The value type that is generated (refer to internal/generator/generator.go)


