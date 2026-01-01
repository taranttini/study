package br.com.igorbag.githubsearch.ui.adapter

import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.ImageView
import android.widget.TextView
import androidx.recyclerview.widget.RecyclerView
import br.com.igorbag.githubsearch.R
import br.com.igorbag.githubsearch.domain.Repository

class RepositoryAdapter(private val repositories: List<Repository>) :
    RecyclerView.Adapter<RepositoryAdapter.ViewHolder>() {

    var btnRepositoryLinkListner: (Repository) -> Unit = {}
    var btnShareListner: (Repository) -> Unit = {}

    // Cria uma nova view
    override fun onCreateViewHolder(parent: ViewGroup, viewType: Int): ViewHolder {
        val view =
            LayoutInflater.from(parent.context).inflate(R.layout.repository_item, parent, false)
        return ViewHolder(view)
    }

    // Pega o conteudo da view e troca pela informacao de item de uma lista
    override fun onBindViewHolder(holder: ViewHolder, position: Int) {
        //@SOLVED_TODO 8 -  Realizar o bind do viewHolder
        holder.nomeRepositorio.text = repositories[position].name

        holder.itemView.setOnClickListener {
            btnRepositoryLinkListner(repositories[position])
        }

        holder.linkRepositorio.setOnClickListener {
            btnShareListner(repositories[position])
        }
    }

    // Pega a quantidade de repositorios da lista
    //@SOLVED_TODO 9 - realizar a contagem da lista
    override fun getItemCount(): Int = repositories.size

    class ViewHolder(view: View) : RecyclerView.ViewHolder(view) {
        //@SOLVED_TODO 10 - Implementar o ViewHolder para os repositorios
        val nomeRepositorio: TextView
        var linkRepositorio: ImageView

        init {
            view.apply {
                nomeRepositorio = findViewById(R.id.tv_repositorio)
                linkRepositorio = findViewById(R.id.iv_favorite)
            }
        }
    }
}


