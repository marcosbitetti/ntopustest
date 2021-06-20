# NTopus test

## Abstract

This code was generated for a job teste, not intended to use in production models.

You can  use as reference to study, I was started  to study Go too, on this project, PRs are welcome.


## Configuration

Copy and rename .env file:

     cp .env.example .env

## Run

**step 1**

    docker-compose up

**step 2**
    
Open your open brownser on http://localhost to monitoring events.

**step 3**
    
Import *examplo.postman* to you postman

**step 4**

Run **populate.sh** file to start some registers

**NOTE** if you wish to change port of monitoring service *(queue_app)* you need to rebuild project in *rabbitmq-monitor-test*. With ***npm run build*** inside *rabbitmq-monitor* folder.


## Project's structure

    /
        /application -> Main project and test subject
        /rabbitmq -> data directory for RabbitMQ
        /rabbitmq-monitor -> [helper] Vue project to monitoring db events
        /queue-app -> websocket server to repass RabitMQ changes to "rabbitmq-monitor"
        /mongodb
            /config -> mongodb start
            /data -> mongodb volatile data

## Docker Stack

Generate a environment with the services:

**application-maintest**

**mongodb_ntopus**

**rabbitmq_ntopus**

**rabbitmq-monitor-test**

**queue_app**


**NOTE1** queue_app always restart when rabbitmq is not done


## Report *[pt-BR]*

Comecei o projeto lendo um pouco sobre Go, em seguida parti para montar o ambiente.

Como estava com pouca experiência nos pacotes, talvez não tenha usado a melhor estrutura para o projeto.

Alguns conceitos de Go eu já havia visto em outros lugares, mas acabei reaprendendo para poder me adaptar.

Como parâmetros de boas práticas, tentei seguir os tutoriais e algumas dicas pelo StackOverflow e GitHub.

Montei uma estrutura padrão de um serviço web, consegui seguir o padrão, separando em camadas com abstração, facilitando até troca do banco.

Optei por usar o Gin, parecia ser bem cotado. Nome bonito, bastante estrelas no GitHub.

Deixei o setup da rota em cada controller (resource), no entanto não sei se foi a melhor estratégia.

Outra nota sobre o projeto, ao executar o monitor, usei uma implementação de websocket bem simplista, então esta havendo problemas se apertar F5 para recarregar o monitor, se isso ocorrer é preciso reiniciar o "queue_app", como não éra o meu foco não senti muita necessidade de melhorar isso no momento.

Os testes foram funcinais até o primeiro momento, mas a partir do momento q criei mais de um testte no arquivo começou a travar. Embora eu estivesse seguindo todos os manuais, a razão não ficou clara de primeira instância, se houver algum tempo viável entre a data deste post e a de apresentação posso verificar melhor isso.
