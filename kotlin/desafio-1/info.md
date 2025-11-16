# Desafio

Você trabalha como consultor júnior em uma empresa de TI especializada em soluções Mobile para Android. Seu colega mais experiente lhe passou uma lista de tarefas a serem realizadas para garantir que o projeto de um aplicativo esteja seguindo as melhores práticas. Cada tarefa pode ter sido marcada como "ok" (quando concluída) ou "pendente" (quando ainda falta ser feita). Para facilitar a revisão, você precisa criar uma rotina que gere um relatório simples: liste apenas as tarefas que ainda estão pendentes, uma por linha, mantendo a ordem de recebimento. Se todas as tarefas estiverem marcadas como "ok", o relatório deve informar "Projeto pronto".

Seu trabalho é processar a entrada, composta por pares de tarefa e status, e imprimir as tarefas pendentes ou a mensagem caso não haja nada pendente.

# Entrada

Uma única linha contendo pares alternados de nome de tarefa e seu status, separados por espaço. O status será "ok" ou "pendente".

# Saída

Cada tarefa pendente em uma linha, seguindo a ordem original. Caso não haja tarefas pendentes, exiba apenas a linha "Projeto pronto".

# Exemplos

A tabela abaixo apresenta exemplos de entrada e saída:

| Entrada |	Saída |
| -- | -- |
| Login ok Splash pendente Banco ok API pendente |	Splash API |
| Layout ok Servicos ok Notificacoes ok	| Projeto pronto |
| Push pendente	| Push |