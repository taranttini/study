package main

// condicionais

func main() {
	valor_1 := 1
	valor_2 := 2

	// so tem if {} [ou] if {} else {}
	if valor_1 > valor_2 {
		println(valor_1)
	} else {
		println(valor_2)
	}

	if 1 < 2 {
		println("1 < 2 \n  SIM 1 é menor que 2")
	}

	if 1 == 1 {
		println("1 == 1 \n  SIM 1 é igual a 1")
	}

	if 2 > 1 {
		println("2 > 1 \n  SIM 2 é maior que 1")
	}

	if 1 > 2 || 2 > 1 {
		println("1 > 2 || 2 > 1 \n  1>2 ? NAO ' OU ' 2>1 ? SIM \n  temos uma condicao verdadeira")
	}

	if 2 > 1 || 1 > 1 {
		println("2 > 1 || 1 > 1 \n  2>1 ? SIM ' OU ' 1>2 ? NAO \n  temos uma condicao verdadeira")
	}

	if 1 == 1 && 2 == 2 {
		println("1 == 1 && 2 == 2 \n  1==1 ? SIM ' E ' 2==2 ? SIM \n  temos duas condicoes verdadeiras")
	}

	if 1 == 1 && 2 == 1 {
		println("algo esta errado")
	} else {
		println("1 == 1 && 2 == 1 \n  1==1 ? SIM ' E ' 2==1 ? NAO \n  nao temos duas condicoes verdadeiras")
	}

	// switch

	valor_1 = 1
	valor_2 = 2
	valor_3 := 3

	// brincar de mudar os valores
	switch valor_3 {
	case 1:
		println("valor_1")
	case 2:
		println("valor_2")
	case 3:
		println("valor_3")
	default:
		println("valor nao encontrado")
	}
}
