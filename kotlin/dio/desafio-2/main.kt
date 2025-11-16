import java.util.*
 
fun main() {
    // Lê a letra inicial desejada (primeiro caractere, remove espaços e normaliza para minúsculo)
    val letraDesejada = readLine()!!.trim().lowercase().firstOrNull()
 
    // Lê a linha com os nomes separados por espaço
    val nomesEntrada = readLine() ?: ""
    val nomesClientes = nomesEntrada.split(" ").filter { it.isNotEmpty() }
 
    // TODO: Filtre somente os nomes que começam com a letra desejada (sem diferenciar maiúsculas/minúsculas)
    // Dica: Use filter e compare apenas o primeiro caractere de cada nome com a letra informada
    val nomesFiltrados = nomesClientes.filter { it.lowercase().firstOrNull() == letraDesejada }
 
    // Impressão conforme especificação
    if (nomesFiltrados.isEmpty()) {
        println("Nenhum cliente encontrado")
    } else {
        println(nomesFiltrados.joinToString(" "))
    }
}