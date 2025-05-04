
public class Main {

	public static void main(String[] args) {
		Cliente venilton = new Cliente();
		venilton.setNome("Venilton");
		
		Conta cc1 = new ContaCorrente(venilton);
		Conta poupanca1 = new ContaPoupanca(venilton);

		cc1.depositar(100);
		cc1.transferir(100, poupanca1);

		cc1.imprimirExtrato();
		poupanca1.imprimirExtrato();

		Cliente jaspion = new Cliente();
		jaspion.setNome("Jaspion");
		
		Conta cc2 = new ContaCorrente(jaspion);
		Conta poupanca2 = new ContaPoupanca(jaspion);

		cc2.depositar(500);
		cc2.depositar(99);
		cc2.transferir(150, poupanca2);
		
		cc2.imprimirExtrato();
		poupanca2.imprimirExtrato();
	}

}
