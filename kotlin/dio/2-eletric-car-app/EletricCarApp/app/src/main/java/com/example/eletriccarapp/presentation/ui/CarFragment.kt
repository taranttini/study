package com.example.eletriccarapp.presentation.ui.adapter

import android.content.Intent

import android.os.Bundle
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.Button
import androidx.fragment.app.Fragment
import androidx.recyclerview.widget.RecyclerView
import com.example.eletriccarapp.R
import com.example.eletriccarapp.presentation.data.CarFactory
import com.example.eletriccarapp.presentation.ui.CalcularAutonomiaActivity
import com.google.android.material.floatingactionbutton.FloatingActionButton

class CarFragment : Fragment(){
    override fun onCreateView(
        inflater: LayoutInflater,
        container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
        return inflater.inflate(R.layout.car_fragment, container, false)
        //super.onCreateView(inflater, container, savedInstanceState)
    }

    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)
        setupView(view)
        setupListerners()
        setupList()
    }

    lateinit var btnRedirect: FloatingActionButton
    lateinit var lista: RecyclerView

    private fun setupView(view: View) {
        view.apply {
            btnRedirect = findViewById(R.id.fab_calcular) //btn_redirect
            lista = findViewById(R.id.rv_carros)
        }

    }

    private fun setupList() {
        val adpater = CarAdapter(CarFactory.list)
        lista.adapter = adpater
    }

    private fun setupListerners() {
        btnRedirect.setOnClickListener {
            startActivity( Intent(context, CalcularAutonomiaActivity::class.java))
        }
    }

}