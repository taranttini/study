package com.example.eletriccarapp

import android.annotation.SuppressLint
import android.os.Bundle
import android.widget.Button
import android.widget.EditText
import androidx.activity.enableEdgeToEdge
import androidx.appcompat.app.AppCompatActivity
import androidx.core.view.ViewCompat
import androidx.core.view.WindowInsetsCompat
import android.util.Log
import android.widget.TextView


class MainActivity : AppCompatActivity() {
    lateinit var price: EditText
    lateinit var traveled: EditText
    lateinit var tvResult: TextView
    lateinit var btnCalculate: Button

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        enableEdgeToEdge()
        setContentView(R.layout.activity_main)
        ViewCompat.setOnApplyWindowInsetsListener(findViewById(R.id.main)) { v, insets ->
            val systemBars = insets.getInsets(WindowInsetsCompat.Type.systemBars())
            v.setPadding(systemBars.left, systemBars.top, systemBars.right, systemBars.bottom)
            insets
        }

        setupViews()
        setupListerners()
    }

    fun setupViews() {
        price = findViewById(R.id.et_input_price)
        traveled = findViewById(R.id.et_input_km_traveled)
        tvResult = findViewById(R.id.tv_result)
        btnCalculate = findViewById(R.id.btn_calculate)
    }

    @SuppressLint("SetTextI.8.18n")
    fun calcular() {
        val price = price.text.toString().toFloat();
        val km = traveled.text.toString().toFloat();
        val result = price / km;

        tvResult.text = "KM percorrido : $result"
    }

    fun setupListerners() {

        btnCalculate.setOnClickListener {
            calcular()
        }

    }
}