package com.example.eletriccarapp.presentation

import android.content.Intent
import android.os.Bundle
import android.util.Log
import android.widget.Button
import androidx.appcompat.app.AppCompatActivity
import androidx.recyclerview.widget.RecyclerView
import com.example.eletriccarapp.R
import com.example.eletriccarapp.presentation.data.CarFactory
import com.example.eletriccarapp.presentation.ui.CarAdapter

class MainActivity : AppCompatActivity() {
    lateinit var btnRedirect: Button
    lateinit var lista: RecyclerView

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        Log.d("Lifecycle", "CREATE")
        setContentView(R.layout.activity_main)
        setupView()
        setupListerners()
        setupList()
    }

    override fun onResume() {
        super.onResume()
        Log.d("Lifecycle", "RESUME")
    }

    override fun onPause() {
        super.onPause()
        Log.d("Lifecycle", "PAUSE")
    }

    override fun onStart() {
        super.onStart()
        Log.d("Lifecycle", "START")
    }

    override fun onDestroy() {
        super.onDestroy()

        Log.d("Lifecycle", "DESTROY")
    }

    private fun setupView() {
        btnRedirect = findViewById(R.id.btn_redirect) //btn_redirect
        lista = findViewById(R.id.rv_carros)
    }

    private fun setupList() {
        val adpater = CarAdapter(CarFactory.list)
        lista.adapter = adpater
    }

    private fun setupListerners() {
        btnRedirect.setOnClickListener {
            startActivity( Intent(this, CalcularAutonomiaActivity::class.java))
        }
    }
}