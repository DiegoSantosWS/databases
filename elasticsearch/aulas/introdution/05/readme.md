# Analyzers comuns




| TIPOS                      | DESC                                                                                                                  |
| :------------------------: | --------------------------------------------------------------------------------------------------------------------: |
| Espaço em branco           | SOGNORA ESPAÇOS                                                                                                       |
| Padrão                     | QUEBRA OS ESPAÇOS, IGNORA ",","." e etc... mas armazena algarizmos.                                                   |
| Simples                    | Parece com padrão, mas não coloca algarizmo                                                                           |
| Idioma(portuguese,englesh) | Faz o mesmo que simples, mas remove acento, reconhece plural e armazena no singular, mas o elastic retorna no plural. |
| :------------------------: | --------------------------------------------------------------------------------------------------------------------: |

# Analyzer phrase

```kibana
GET catalogo/_analyze
{
  "text": "Eu nasci há 10 mil (sim, isso mesmo, 10 mil) anos atrás"
}
```

### Analyzer Simple

```kibana
GET catalogo/_analyze
{
  "analyzer": "simple", 
  "text": "Eu nasci há 10 mil (sim, isso mesmo, 10 mil) anos atrás"
}
```

### Analyzer whitespace

```kibana
GET catalogo/_analyze
{
  "analyzer": "whitespace", 
  "text": "Eu nasci há 10 mil (sim, isso mesmo, 10 mil) anos atrás"
}
```

### Analyzer idioma

```kibana
GET catalogo/_analyze
{
  "analyzer": "portuguese", 
  "text": "Eu nasci há 10 mil (sim, isso mesmo, 10 mil) anos atrás"
}
```