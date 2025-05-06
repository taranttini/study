package com.taranttini;
// Para ler e escrever dados em Java, aqui na DIO padronizamos da seguinte forma:
// - new Scanner(System.in): cria um leitor de Entradas, com métodos úteis com prefixo "next";
// - System.out.println:.imprime um texto de Saída (Output) e pulando uma linha.


import java.util.Scanner;

class MainContaSimples_Conta {

    // TODO: Implemente os atributos privados para garantir que só podem ser acessados dentro da classe:
    private String titular;
    private int numeroConta;
    private double saldo;

    // TODO: Crie o construtor para inicializar os atributos da conta:
    public MainContaSimples_Conta(String titular, int numeroConta, double saldoInicial) {
        this.titular = titular;
        this.numeroConta = numeroConta;
        this.saldo = saldoInicial;
    }

    // TODO: Adicione o método público para consultar o saldo:
    public double consultarSaldo() {
        return this.saldo;
    }

    // TODO: Adicione o método público para realizar saque:
    public Boolean sacar(double valorSaque) {
        if (valorSaque > this.saldo) {
            return false;
        }

        this.saldo -= valorSaque;

        return true;
    }

    // TODO: Crie o método público para realizar depósito
    public void depositar(double valorDeposito) {
        this.saldo += valorDeposito;
    }

}

public class MainContaSimples {

    public static void main(String[] args) {


        Scanner sc = new Scanner(System.in);

        String titular = sc.nextLine();
        int numeroConta = sc.nextInt();
        double saldoInicial = sc.nextDouble();

        MainContaSimples_Conta conta = new MainContaSimples_Conta(titular, numeroConta, saldoInicial);

        // Operações de saque e depósito
        if (sc.hasNextDouble()) {
            double valorSaque = sc.nextDouble();
            if (!conta.sacar(valorSaque)) {
                System.out.println("Saque invalido: Saldo insuficiente");
                System.out.println("Saldo Atual: " + String.format("%.2f", conta.consultarSaldo()));
                return;
            }
        }

        if (sc.hasNextDouble()) {
            double valorDeposito = sc.nextDouble();
            conta.depositar(valorDeposito);
        }

        // Imprime o saldo final atualizado
        System.out.println("Saldo Atualizado: " + String.format("%.2f", conta.consultarSaldo()));

        sc.close(); // Fecha o scanner
    }
}