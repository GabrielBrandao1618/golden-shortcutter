# golden-shortcutter

Uma API de encurtador de links simples

## Endpoints

[/generatrUrl](#generateurl)

[/getUrl](#geturl)

### /generateUrl

#### Parâmetros

| nome | tipo   |
| ---- | ------ |
| ref  | string |
| name | string |

É utilizado para insetir um novo link encurtado, essa rota recebe dois parâmetros através de um objeto JSON: ref e name. Exemplo à seguir:

```json
{
  "ref": "github.com",
  "name": "github"
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

#### Parâmetros

| nome | tipo   |
| ---- | ------ |
| name | string |

É utilizado para retornar o link criado, para isso ele recebe o nome customizado como parâmetro. Exemplo de requisição à seguir:

```json
{
  "name": "googleImage"
}
```

Exemplo de resposta do servidor

```json
{
  "ref": "image.google.com",
  "name": "googleImage"
}
```

Caso não exista um registro com o nome fornecido, a propriedade name terá um valor de string vazia (`" "`)
