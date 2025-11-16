# Desafio

Numa consultoria de TI especializada em projetos de apps Android, a equipe utiliza uma agenda digital exclusiva para registrar os clientes e reuniões do dia. Sempre que uma nova reunião é marcada, o nome do cliente é adicionado a uma lista. Porém, para manter o foco e evitar distrações, é necessário identificar rapidamente quais clientes devem receber uma notificação sobre o horário de sua reunião. Seu papel é criar um recurso que, dado o nome de todos os clientes agendados em um dia, informe quais têm nomes que começam com uma mesma letra específica escolhida pelo consultor. Assim, será possível notificar apenas esse grupo de clientes.

Implemente um programa que leia uma letra e, em seguida, uma lista com os nomes dos clientes agendados, separados por espaço. Sua tarefa é listar todos os nomes que começam com a letra indicada, mantendo a ordem original. Se nenhum nome corresponder, exiba a mensagem "Nenhum cliente encontrado". Utilize estruturas de controle de fluxo e coleções conforme o esperado em apps Android.

# Entrada

A primeira linha consiste de uma letra (maiuscula ou minuscula), indicando a letra inicial desejada.
A segunda linha consiste de uma sequencia de nomes (palavras sem espacos internos), separados por um unico espaco.

# Saída

Uma unica linha contendo todos os nomes que comecam com a letra informada, separados por um espaco, na ordem original. Caso nenhum nome seja correspondente, exiba exatamente: Nenhum cliente encontrado

# Exemplos

A tabela abaixo apresenta exemplos de entrada e saída:

| Entrada |	Saída |
| -- | -- |
| m <br> marina marcelo paulo miguel	| marina marcelo miguel |
| P <br> patricia clara carlos paulo |	patricia paulo |
| t <br> aline beatriz carlos |	Nenhum cliente encontrado |
