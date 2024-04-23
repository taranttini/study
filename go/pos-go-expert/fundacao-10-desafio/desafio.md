Neste desafio você terá que usar o que aprendemos com Multithreading e APIs para buscar o resultado mais rápido entre duas APIs distintas.

As duas requisições serão feitas simultaneamente para as seguintes APIs:

https://brasilapi.com.br/api/cep/v1/ + cep

http://viacep.com.br/ws/ + cep + /json/

Os requisitos para este desafio são:

- Acatar a API que entregar a resposta mais rápida e descartar a resposta mais lenta.

- O resultado da request deverá ser exibido no command line com os dados do endereço, bem como qual API a enviou.

- Limitar o tempo de resposta em 1 segundo. Caso contrário, o erro de timeout deve ser exibido.



Saída ViaBrasil

```json
{
  "cep": "02765000",
  "state": "SP",
  "city": "São Paulo",
  "neighborhood": "Vila Hebe",
  "street": "Rua Daniel de Toledo",
  "service": "widenet"
}
```

Saída ViaCep

```json
{
  "cep": "02765-000",
  "logradouro": "Rua Daniel de Toledo",
  "complemento": "",
  "bairro": "Vila Hebe",
  "localidade": "São Paulo",
  "uf": "SP",
  "ibge": "3550308",
  "gia": "1004",
  "ddd": "11",
  "siafi": "7107"
}
```