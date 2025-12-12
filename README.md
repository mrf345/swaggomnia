
Generate Swagger Documentation from Insomnium REST Client (maintained fork of [Fyb3roptik/swaggomnia](https://github.com/Fyb3roptik/swaggomnia))

## Install

```
go install github.com/mrf345/swaggomnia@latest
```

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
| -i | Insomnia exported file |
| -c | API Global Configuration file (see [Configuration Format](#configuration-format))|
| -o | Insomnia output format (json or yaml, default yaml)  |


## Example

Let's convert the following Insomnia API documentation to Swagger:


```
$ swaggomnia generate -i examples/watchnow.json -c examples/config.json -o json
```

## Configuration Format

> [!TIP]
> `security` field is case-sensitive and only supports `Bearer`, `Basic Auth`, `API Key`.

```json
{
  "title" : "API Name",
  "version" : "API version",
  "basePath" : "https://api.domain.com/v1",
  "description" : "API description",
  "security": "Bearer"
}
```

### Updating `tmpl/`

1. `go install github.com/go-bindata/go-bindata/...@latest`
2. make some changes to `tmpl/swagger.yaml`
3. `go-bindata -o template.go tmpl/`


### Changes & TODOs

- [x] enable request body string
- [x] enable escaped resource description content (useful for showing code snippets)
- [x] make query params optional by default
- [x] hide empty groups
- [x] add security definitions and options to config
- [x] sort by group names instead of ids
- [x] enable query params description
- [x] enable and replace url path params `{% request 'parameter', 'id', 0 %}` with `{id}`
