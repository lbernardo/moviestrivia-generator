# moviestrivia generator

## Configure api

Create a API TOKEN in https://developer.themoviedb.org/docs/getting-started and create `config.json`

```json
{
  "base_url": "https://api.themoviedb.org/3",
  "token": "XXXXXXX"
}
```

## Generate

```shell
go run main.go
```

Is generated the file: `result.json` and you need put in https://github.com/lbernardo/moviestrivia/blob/main/lib/getItems.ts