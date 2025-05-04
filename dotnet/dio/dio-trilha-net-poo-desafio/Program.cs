using DesafioPOO.Models;

// TODO: Realizar os testes com as classes Nokia e Iphone
Console.WriteLine("Criando Smartphone Nokia");
Smartphone nokia = new Nokia("nokia 2000", "NL2020", "123456789012345", 16);
nokia.Ligar();
nokia.InstalarAplicativo("Jogo da cobrinha");
nokia.ReceberLigacao();
nokia.InstalarAplicativo("Jogo Sudoku");

Console.WriteLine("----");

Console.WriteLine("Criando Smartphone Iphone");
Smartphone iphone = new Iphone("iphone 6", "6", "123456789012346", 64);
iphone.Ligar();
iphone.InstalarAplicativo("Jogo Subway");
iphone.ReceberLigacao();
iphone.InstalarAplicativo("Jogo Sudoku");


