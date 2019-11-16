# ProjetoFinalDistribuida

Projeto final da disciplina INE410130 - Computação Distribuída

Funcionamento geral: Requests são feitos no localhost:8080/exec

Nesse endpoint o circuit breaker é chamado para a execução remota de uma função remota.

Configurações do CB pode ser feita no arquivo cb.json, definindo threshold para falha e tempo para transição para o estado half-open

Configurações da simulação no servidor que implementa a RPC:
Alterar o config.json que encontra-se na pasta remotecall: definir dois momentos para falhar e tempo para o serviço ficar fora.





