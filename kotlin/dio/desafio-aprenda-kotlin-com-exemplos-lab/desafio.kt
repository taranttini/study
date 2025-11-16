enum class Nivel { BASICO, INTERMEDIARIO, DIFICIL }

class Usuario(var nome: String, var cpf: String, var dataNascimento: String)

data class ConteudoEducacional(var nome: String, val duracao: Int = 60)

data class Formacao(val nome: String, var conteudos: List<ConteudoEducacional>) {

    val inscritos = mutableListOf<Usuario>()
    //val turma = mutableListOf<Pair(String, mutableListOf<Usuario>())>()
    val turma = ArrayList< Pair< String,  ArrayList<Usuario> >>()  // ArrayList<String>
    
    fun printNomeInscritos() {
        for (i in inscritos) println(i.nome)
    }
     
    fun usuarioJaMatriculado(usuario: Usuario, inscritos: List<Usuario>): Boolean {        
        return inscritos.filter { entry -> entry.cpf == usuario.cpf }.any()
    }
    
    fun enturmar(usuario: Usuario, conteudo: ConteudoEducacional ) {        
        
        var result = obterTurma(conteudo)
        if ( usuarioJaMatriculado(usuario, result.second) ) {
            println("Usuário já está enturmado [${usuario.nome}]")
        } else {
        	result.second.add(usuario)    
            println("Usuário enturmado [${usuario.nome}]")
        }
        
        
        //println(usuario.nome)
        //println(result)
        //return result
    }
    
    fun matricular(usuario: Usuario) {        
        //TODO("Utilize o parâmetro $usuario para simular uma matrícula (usar a lista de $inscritos).")
        
        if ( usuarioJaMatriculado(usuario, inscritos) )  {
        	println("Usuario já está matriculado [${usuario.nome}]")
        }
        else {
            inscritos.add(usuario)
            println("Usuario matriculado [${usuario.nome}]")
            //printNomeInscritos()
        } 
    }
    
    fun obterTurma(conteudo: ConteudoEducacional): Pair<String, ArrayList<Usuario>> {
        
        var lista = conteudos.filter { entry -> entry.nome == conteudo.nome }
        if (!lista.any()) {
            throw Exception("conteúdo não cadastrado na DIO")
        }
        
        var result = turma.filter { entry -> entry.first == conteudo.nome }
        if (result.any()) {
            return result.first()
        }
        else {
            var list = ArrayList<Usuario>()
            turma.add(Pair<String, ArrayList<Usuario>>(conteudo.nome, list) )
        }
        result = turma.filter { entry -> entry.first == conteudo.nome }
        return result.first()        
    }
    
    fun exibirAlunosTurma(conteudo: ConteudoEducacional) {
        var result = obterTurma(conteudo)
        if (result.second.any()) {
            
            println("\nExibir alunos da turma [${conteudo.nome}] ")
        
            for(item in result.second) {
                println("\t ${item.nome}")
            }
        }
        else {
         	println("\nTurma sem alunos [${conteudo.nome}]")   
        }
    }
}

fun main() {
    
    //TODO("Analise as classes modeladas para este domínio de aplicação e pense em formas de evoluí-las.")
    //TODO("Simule alguns cenários de teste. Para isso, crie alguns objetos usando as classes em questão.")
    
    var cursos = mutableListOf<ConteudoEducacional>(
        ConteudoEducacional("Programacao Kotlin 1")
        
        
    );
    var c2 = ConteudoEducacional("Programacao Kotlin 2")
    cursos.add(c2)
    
    var c3 = ConteudoEducacional("Programacao Kotlin 3")   
    cursos.add(c3)
    
    var f = Formacao("x", cursos)
    var u1 = Usuario(nome = "Ben", cpf = "01", dataNascimento = "01/01/2017")    
    var u2 = Usuario(nome = "Ben 2", cpf = "02", dataNascimento = "01/01/2017")
    
    println("Exibir usuários matriculados:")
    f.printNomeInscritos()
    
    
    f.matricular(u1)
    // tentar matricular novamente ( vai gerar log de aviso e nao incluir novamente )
    f.matricular(u1)
    
    f.matricular(u2)
    
    println("\nExibir usuários matriculados:")
    f.printNomeInscritos()
    
    
    // pegar turma, ou criar uma nova se nao existir
    var t = f.obterTurma(c3)    
    
    println("")
    // enturmar usuario2 no curso3
    f.enturmar(u2, c3)
    // tentar enturmar usuario2 no curso3 novamente ( vair gerar log de aviso e nao incluir novamente ) 
    f.enturmar(u2, c3)
    
    f.enturmar(u1, c3)
    f.enturmar(u1, c3)
    f.enturmar(u2, c3)
    
    // imprimir alunos da turma 3
    f.exibirAlunosTurma(c3)
    f.exibirAlunosTurma(c2)
    
    println (t)
    var c4 = ConteudoEducacional("Programacao Kotlin 4")   
    f.enturmar(u2, c4)
    
}