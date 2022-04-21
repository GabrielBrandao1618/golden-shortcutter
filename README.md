# golden-shortcutter
Uma API de encurtador de links simples

O funcionamento da API se baseia em dois endpoints: generateUrl e getUrl

O primeiro é utilizado para insetir um novo link encurtado, assim que inserido, é retornado um uuid(hashCode) que faz referência ao link inserido

O segundo é utilizado para retornar o link inserido, para isso ele deve receber o uuid(hashCode) que faz referência ao link inserido
