# Desafio 2

## Solução

Foi feito o básico pedido mais esses extras de melhoramento: 

1. Foi usado Context para cancelar e liberar o recurso da requisição que fosse mais lenta após a primeira concluir; 


1. Tem um For no Select para esperar pelo menos 2 resultados, pois o primeiro resultado poderia ser de falha;
1. Foi tratado o problema caso dê  erro nas duas requisições.

## Requisitos

 Neste desafio você terá que usar o que aprendemos com Multithreading e APIs para buscar o resultado mais rápido entre duas APIs distintas.

As duas requisições serão feitas simultaneamente para as seguintes APIs:

https://brasilapi.com.br/api/cep/v1/01153000 + cep

http://viacep.com.br/ws/" + cep + "/json/

Os requisitos para este desafio são:

- Acatar a API que entregar a resposta mais rápida e descartar a resposta mais lenta.

- O resultado da request deverá ser exibido no command line com os dados do endereço, bem como qual API a enviou.

- Limitar o tempo de resposta em 1 segundo. Caso contrário, o erro de timeout deve ser exibido.
