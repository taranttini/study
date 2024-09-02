# INFO

Esse desafio foi complicado, pois até mesmo a aula de telemetria do curso com os exemplos parecem não ter funcionado no meu ambiente com docker, fiz ajustes onde achei que conseguiria e onde parece ter sido possível, infelizmente o zipkin não tive tanto sucesso, perdi alguns dias fazendo, desfazendo, montando e desmontando mas não tive sucesso.

Pediria um auxilio com essa tarefa, ou algo que eu possa fazer para ajudar.

Parece que consegui fazer os log e span no jagger, mas não consigo entender como mostra a informação de log lá ou os dados que foram solicitados.

Tinha feito também um exemplo usando o que o site informa, só que lá só deu certo via console, e ai mesmo fuçando não tive muito sucesso, então acredito que para essa tarefa eu não consegui aprender ou não consegui evoluir como gostaria, não sei se esse ponto que eu fiz atende para conclusão do curso, pois parece ser um item também fundamental, parece que não temos meio termo nesse curso para completar os desafios ou conseguir ser aprovado e assim receber o diploma.

Grato pela atenção o responsável pela correção.

Para rodar o projeto é possível acessar o diretório http e lá tem as chamadas para testar os serviços, eu acabei isolando o serviço de CEP, então ele é auto suficiente, já o serviço de temperatura ele processa com base no que ele receber então ambos funcionam separados e para completar o desafio eles funcionam em conjunto.

para rodar o projeto precisamos fazer o comando `docker-compose up`

E assim ele sobe todos os serviços que parecem ser necessários para fazer alguma telemetria.

Fico no aguardo de dicas ou até mesmo de um meio aceite para que possamos completar a tarefa e o curso.


# Tracing distribuído e span

Objetivo: Desenvolver um sistema em Go que receba um CEP, identifica a cidade e retorna o clima atual (temperatura em graus celsius, fahrenheit e kelvin) juntamente com a cidade. Esse sistema deverá implementar OTEL(Open Telemetry) e Zipkin.

Basedo no cenário conhecido "Sistema de temperatura por CEP" denominado Serviço B, será incluso um novo projeto, denominado Serviço A.



Requisitos - Serviço A (responsável pelo input):

    O sistema deve receber um input de 8 dígitos via POST, através do schema:  { "cep": "29902555" }
    O sistema deve validar se o input é valido (contem 8 dígitos) e é uma STRING
        Caso seja válido, será encaminhado para o Serviço B via HTTP
        Caso não seja válido, deve retornar:
            Código HTTP: 422
            Mensagem: invalid zipcode

Requisitos - Serviço B (responsável pela orquestração):

    O sistema deve receber um CEP válido de 8 digitos
    O sistema deve realizar a pesquisa do CEP e encontrar o nome da localização, a partir disso, deverá retornar as temperaturas e formata-lás em: Celsius, Fahrenheit, Kelvin juntamente com o nome da localização.
    O sistema deve responder adequadamente nos seguintes cenários:
        Em caso de sucesso:
            Código HTTP: 200
            Response Body: { "city: "São Paulo", "temp_C": 28.5, "temp_F": 28.5, "temp_K": 28.5 }
        Em caso de falha, caso o CEP não seja válido (com formato correto):
            Código HTTP: 422
            Mensagem: invalid zipcode
        ​​​Em caso de falha, caso o CEP não seja encontrado:
            Código HTTP: 404
            Mensagem: can not find zipcode

Após a implementação dos serviços, adicione a implementação do OTEL + Zipkin:

    Implementar tracing distribuído entre Serviço A - Serviço B
    Utilizar span para medir o tempo de resposta do serviço de busca de CEP e busca de temperatura

Dicas:

    Utilize a API viaCEP (ou similar) para encontrar a localização que deseja consultar a temperatura: https://viacep.com.br/
    Utilize a API WeatherAPI (ou similar) para consultar as temperaturas desejadas: https://www.weatherapi.com/
    Para realizar a conversão de Celsius para Fahrenheit, utilize a seguinte fórmula: F = C * 1,8 + 32
    Para realizar a conversão de Celsius para Kelvin, utilize a seguinte fórmula: K = C + 273
        Sendo F = Fahrenheit
        Sendo C = Celsius
        Sendo K = Kelvin
    Para dúvidas da implementação do OTEL, você pode clicar aqui
    Para implementação de spans, você pode clicar aqui
    Você precisará utilizar um serviço de collector do OTEL
    Para mais informações sobre Zipkin, você pode clicar aqui

Entrega:

    O código-fonte completo da implementação.
    Documentação explicando como rodar o projeto em ambiente dev.
    Utilize docker/docker-compose para que possamos realizar os testes de sua aplicação.
