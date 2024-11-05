
public class Main {

	public static void main(String[] args) {
		Cliente jaspion = new Cliente();
		jaspion.setNome("Jaspion");
		
		Conta cc = new ContaCorrente(jaspion);
		Conta poupanca = new ContaPoupanca(jaspion);

		cc.depositar(500);
		cc.depositar(99);
		cc.transferir(150, poupanca);
		
		cc.imprimirExtrato();
		poupanca.imprimirExtrato();
	}

}
