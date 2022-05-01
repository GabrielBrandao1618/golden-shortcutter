# golden-shortcutter

Uma API de encurtador de links simples

## Endpoints

[/createLink](#createlink)

[/getUrl](#geturl)

### /createLink

Método: `POST`

#### Parâmetros

| nome | tipo   |
| ---- | ------ |
| ref  | string |
| name | string |

É utilizado para insetir um novo link encurtado, essa rota recebe dois parâmetros através de um objeto JSON: ref e name. Exemplo à seguir:

```json
{
  "ref": "image.google.com",
  "name": "googleImage"
}
```

Exemplo de resposta do servidor

```json
{
  "Sucess": true,
  "Msg": "Link created successfully!",
  "ShortedLink": {
    "ref": "image.google.com",
    "name": "googleImage"
  }
}
```

Caso o link já tenha sido criado porém com outro nome customizado(ou até o mesmo), o servidor apenas responderá com o link já criado

### /getUrl

Método: `GET`

É utilizado para retornar o link criado, para isso ele recebe o nome customizado como um parâmetro de rota, por exemplo:

```bash
curl /getUrl/googleImage
```

Exemplo de resposta do servidor

```json
{
  "ref": "image.google.com",
  "name": "googleImage"
}
```

Caso não exista um registro com o nome fornecido, a propriedade name terá um valor de string vazia (`" "`)
