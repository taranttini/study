package com.example.eletriccarapp.presentation.ui

import android.os.Bundle
import android.widget.Button
import android.widget.EditText
import android.widget.ImageView
import android.widget.TextView
import androidx.appcompat.app.AppCompatActivity
import com.example.eletriccarapp.R

class CalcularAutonomiaActivity : AppCompatActivity() {

    lateinit var price: EditText
    lateinit var traveled: EditText
    lateinit var tvResult: TextView
    lateinit var btnCalcular: Button
    lateinit var btnClose: ImageView

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_calcular_autonomia)
        setupViews()
        setupListerners()
    }

    private fun setupViews() {
        price = findViewById(R.id.et_input_price)
        traveled = findViewById(R.id.et_input_km_traveled)
        tvResult = findViewById(R.id.tv_result)
        btnCalcular = findViewById(R.id.btn_calculate)
        btnClose = findViewById(R.id.iv_close)
    }

    fun setupListerners() {
        btnCalcular.setOnClickListener {
            calcular()
        }
        btnClose.setOnClickListener {
            finish()
        }
    }

    fun calcular() {
        val price = price.text.toString().toFloat()
        val km = traveled.text.toString().toFloat()
        val result = price / km

        tvResult.text = "KM percorrido : $result"
    }
}