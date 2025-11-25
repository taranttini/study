package com.example.eletriccarapp.presentation.ui

import android.net.Uri
import android.view.LayoutInflater
import androidx.recyclerview.widget.RecyclerView
import com.example.eletriccarapp.R
import android.view.View;
import android.view.ViewGroup
import android.widget.ImageView
import android.widget.TextView;
import com.example.eletriccarapp.presentation.domain.Carro


class CarAdapter(private val carros: List<Carro>) : RecyclerView.Adapter<CarAdapter.ViewCar>() {


    override fun onCreateViewHolder(parent: ViewGroup, viewType: Int): ViewCar {
        val view = LayoutInflater.from(parent.context).inflate(R.layout.car_item, parent, false)

        return ViewCar(view)
    }

    override fun onBindViewHolder(holder: ViewCar, position: Int) {

        holder.preco.text = carros[position].preco
        holder.recarga.text = carros[position].recarga
        holder.bateria.text = carros[position].bateria
        holder.potencia.text = carros[position].potencia
        //holder.urlPhoto.setImageURI(Uri.parse(carros[position].urlPhoto) )

    }

    // pega o conteudo da view e troca pela informacao de uma lista
    override fun getItemCount(): Int = carros.size


    class ViewCar(view: View) : RecyclerView.ViewHolder(view) {
        val preco: TextView
        val bateria: TextView
        val potencia: TextView
        val recarga: TextView
        val urlPhoto: ImageView


        init {
            view.apply {
                preco = findViewById(R.id.tv_car1_preco_value)
                bateria = findViewById(R.id.tv_car1_bateria_value)
                potencia = findViewById(R.id.tv_car1_potency_value)
                recarga = findViewById(R.id.tv_car1_recharger_value)
                urlPhoto = findViewById(R.id.iv_car1_img)
            }
        }
    }

}



