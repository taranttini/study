package com.taranttini;

// Para ler e escrever dados em Java, aqui na DIO padronizamos da seguinte forma:
// - new Scanner(System.in): cria um leitor de Entradas, com métodos úteis com prefixo "next";
// - System.out.println:.imprime um texto de Saída (Output) e pulando uma linha.

import java.util.Scanner;

abstract class MainContaCorrentePoupanca_Conta {
    protected double saldo;


    public MainContaCorrentePoupanca_Conta(double saldo) {
        this.saldo = saldo;
    }


    public abstract void sacar(double valor);


    public void exibirSaldo() {
        System.out.printf("Saldo Atual: %.2f%n", saldo);
    }
}


class ContaCorrente extends MainContaCorrentePoupanca_Conta {
    private final double limite;

    public ContaCorrente(double saldo, double limite) {
        super(saldo);
        this.limite = limite;
    }

    @Override
    public void sacar(double valor) {
        // TODO: Implemente a lógica para verificar se o saque é permitido considerando o saldo e o limite:
        // Dica: Se saldo - valor >= -limite, o saque é permitido.
        if (saldo - valor >= -limite) {
            saldo -= valor;
            System.out.printf("Saque realizado: %.2f%n", valor);
        } else {
            System.out.println("Saque invalido: Excede limite");
        }

        exibirSaldo(); // Exibe o saldo atualizado
    }
}


class ContaPoupanca extends MainContaCorrentePoupanca_Conta {

    public ContaPoupanca(double saldo) {
        super(saldo);
    }

    // Implementação do método sacar para Conta Poupança
    @Override
    public void sacar(double valor) {
        // TODO: Implemente a lógica para verificar se o saque é permitido considerando apenas o saldo:
        // Dica: Se saldo >= valor, o saque é permitido.
        if (saldo >= valor) {
            saldo -= valor;
            System.out.printf("Saque realizado: %.2f%n", valor);
        } else {
            System.out.println("Saque invalido: Saldo insuficiente");
        }

        exibirSaldo(); // Exibe o saldo atualizado
    }
}


public class MainContaCorrentePoupanca {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);

        String tipoConta = scanner.next(); // Recebe o tipo de conta (Corrente ou Poupança)
        scanner.nextLine();
        scanner.nextLine();
        scanner.next();
        double saldoInicial = scanner.nextDouble(); // Recebe o saldo inicial

        MainContaCorrentePoupanca_Conta conta;

        // TODO: Implemente a lógica para criar uma instância de ContaCorrente ou ContaPoupanca:
        // Dica: Use um if para verificar o tipo da conta.
        if (tipoConta.equals("Corrente")) {
            double limite = scanner.nextDouble(); // Recebe o limite de cheque especial para Conta Corrente
            conta = new ContaCorrente(saldoInicial, limite);
        } else {
            conta = new ContaPoupanca(saldoInicial);
        }

        while (scanner.hasNextDouble()) {
            double valorSaque = scanner.nextDouble();
            conta.sacar(valorSaque);
        }

        scanner.close();
    }
}