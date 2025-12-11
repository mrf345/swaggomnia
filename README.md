
Generate Swagger Documentation from Insomnium REST Client (maintained fork of [Fyb3roptik/swaggomnia](https://github.com/Fyb3roptik/swaggomnia))

## Install

```
go install github.com/mrf345/swaggymnia@latest
```

## Changes

- Enables request body string
- Enables unescaped resource description (useful for showing code snippets)

## How to use it

See usage with:

```
$ swaggomnia --help
```

Generate Swagger documentation:

```
$ swaggomnia generate -insomnia INSOMNIA_EXPORTED_FILE -config CONFIG_FILE -output FORMAT
```

| Option | Description |
| ------ | ----------- |
| -insomnia | Insomnia exported file |
| -config | API Global Configuration file (see [Configuration Format](#configuration-format))|
| -output | Insomnia output format (json or yaml, default yaml)  |


## Example

Let's convert the following Insomnia API documentation to Swagger:


```
$ swaggomnia generate -i examples/watchnow.json -c examples/config.json -o json
```

## Configuration Format

```
{
  "title" : "API Name",
  "version" : "API version",
  "basePath" : "https://api.domain.com/v1",
  "description" : "API description"
}
```
