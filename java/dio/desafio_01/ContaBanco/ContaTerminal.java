package study.java.dio.desafio_01.ContaBanco;

import java.util.Locale;
import java.util.Scanner;

public class ContaTerminal {
    public static void main(String[] args) {
        // leitura dos dados via input
        // criando o objeto scanner

        String nomeCliente = "";// [Nome Cliente]";
        String agencia = "";// "[Agência]";
        Integer numero = 0;// "[Número]";
        Float saldo = (float) 0.0; // [Saldo]

        try (Scanner scanner = new Scanner(System.in).useLocale(Locale.US)) {
            try {
                System.out.println("Por favor, digite o nome do cliente !");
                nomeCliente = scanner.next();
            } catch (Exception e) {
                System.out.println("Erro ao obter o nome do cliente!");
                return;
            }

            try {
                System.out.println("Por favor, digite o código da Agência !");
                agencia = scanner.next();
            } catch (Exception e) {
                System.out.println("Erro ao obter o código da agência!");
                return;
            }

            try {
                System.out.println("Por favor, digite o número da Conta !");
                numero = scanner.nextInt();
            } catch (Exception e) {
                System.out.println("Erro ao obter o número da Conta!");
                return;
            }

            try {
                System.out.println("Por favor, digite o valor do saldo !");
                saldo = scanner.nextFloat();
            } catch (Exception e) {
                System.out.println("Erro ao obter o valor do saldo");
                return;
            }
        } catch (Exception ex) {
            System.out.println("Erro ao processar leitura" + ex.getMessage());
        }

        String aviso = "Olá " + nomeCliente +
                ", obrigado por criar uma conta em nosso banco" +
                ", sua agência é " + agencia +
                ", conta " + numero +
                " e seu saldo " + saldo +
                " já está disponível para saque.";

        System.out.println(aviso);
    }
}
