package com.example.eletriccarapp.presentation.data

import com.example.eletriccarapp.presentation.domain.Carro

object CarFactory {
    val list = listOf(
        Carro(
            id=1,
            nome = "Testa Rossa 2026",
            preco = "R$ 7.000.000,00",
            bateria = "300kWh",
            potencia="800cv",
            recarga="30 min",
            urlPhoto = "@drawable/ferrari_849_testarossa_2026.jpg"
        ),
        Carro(
            id=2,
            nome = "Amalfi 2026",
            preco = "R$ 3.000.000,00",
            bateria = "200kWh",
            potencia="200cv",
            recarga="40 min",
            urlPhoto = "@drawable/ferrari_amalfi_2026.jpg"
        ),
        Carro(
            id=3,
            nome = "F80 2025",
            preco = "R$ 2.000.000,00",
            bateria = "200kWh",
            potencia="200cv",
            recarga="40 min",
            urlPhoto = "@drawable/ferrari_f80_2025.jpg"
        )
    )
}