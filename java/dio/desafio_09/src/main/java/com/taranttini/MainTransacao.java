package com.taranttini;

import java.util.Scanner;

public class MainTransacao {
    public static void main(String[] args) {

        Scanner scanner = new Scanner(System.in);

        Double[] transacoes = new Double[3];

        for (int i = 0; i < transacoes.length; i++) {
            transacoes[i] = scanner.nextDouble();
        }

        Double transacaoComparar = scanner.nextDouble();

        int contador = 0;

        for (int i = 0; i < transacoes.length; i++) {
            if (transacoes[i] > transacaoComparar) {
                contador++;
            }
        }

        System.out.println(contador);
    }
}

/*

Atuo com soluções no desenvolvimento software desde o ano 2004,
desde então venho atuando durante algumas etapas da evolução tecnológica
dentro das grandes corporações governamentais ou financeiras, atuei com MVC, WCF, SOAP,
Microsserviços, 3Tier, Código Limpo, DDD, Arquitetura Hexagonal, Arquitetura Distribuída,
Filas/Mensageria, Stream, Eventos, Bancos Relacionais, NoSql, Batch, Arquivos, SideCard,
ApiGateway, Teste Unitários/Integrados, CI/CD, Docker, Kubernetes, Log, OpenTelemetry,
JWT, Git, Jira, Azure, Aws, Gcp, Firebase. Ajudei na integração de sistemas, conversões
de legados monolíticos para modelos modular ou microsserviço, melhoria de performance no
processamento de lotes de dados ou migração de grandes volumes de dados, atuação em
Sistemas como a Bolsa Eletrônica de Compras do estado de São Paulo, ou captação de
1 milhão de clientes Mei, e integração com os fluxos internos de clientes devido a
fusão da empresa. Processamento de boletos eletrônicos, e processamento de ETL para
filtrar clientes com potencial de adesão a cartela de produtos, processamento de
filas e tópicos e notificações via push para os clientes ou avisos de boletos a vencer,
construção de dashboards ou painel de onboard para validar o crescimento das contas,
utilização de IA ou serviços de reconhecimento para validação de documentos ou
biometria dos usuários.

*/
