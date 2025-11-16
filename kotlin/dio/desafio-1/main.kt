import java.util.Scanner

fun main() { 

  val input = readLine()?.trim().orEmpty()
  
  val tokens = input.split(" ")
  val pendentes = mutableListOf<String>()
  
  // TODO: Percorra os pares (tarefa, status) e adicione à lista 'pendentes' apenas as tarefas com status "pendente"
  // Dica: Avance dois em dois e verifique o status correspondente
  var idx = 0
  for (token in tokens) {
   
    if (tokens[idx] == "pendente") {
      pendentes.add(tokens[idx-1])
    }
    idx += 1 
  }
  // Exibe as tarefas pendentes, mantendo a ordem, ou a mensagem padrão caso não haja nenhuma
  if (pendentes.isNotEmpty()) {
      for (tarefa in pendentes) {
          println(tarefa)
      }
  } else {
      println("Projeto pronto")
  }
}