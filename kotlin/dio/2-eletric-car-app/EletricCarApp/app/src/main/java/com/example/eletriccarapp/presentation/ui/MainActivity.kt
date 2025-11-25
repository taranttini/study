package com.example.eletriccarapp.presentation.ui

import android.content.Intent
import android.os.Bundle
import android.util.Log
import android.widget.Button
import android.widget.LinearLayout
import androidx.appcompat.app.AppCompatActivity
import androidx.recyclerview.widget.RecyclerView
import androidx.viewpager.widget.ViewPager
import androidx.viewpager2.widget.ViewPager2
import com.example.eletriccarapp.R
import com.example.eletriccarapp.presentation.data.CarFactory
import com.example.eletriccarapp.presentation.ui.adapter.CarAdapter
import com.google.android.material.tabs.TabLayout

class MainActivity : AppCompatActivity() {
    lateinit var btnRedirect: Button
    lateinit var lista: RecyclerView
    lateinit var tabLayout: TabLayout
    lateinit var viewPager: ViewPager2

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        Log.d("Lifecycle", "CREATE")
        setContentView(R.layout.activity_main)
        setupView()
        setupListerners()
        setupList()
        setupTabs()
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
        tabLayout = findViewById(R.id.tab_layout)
        btnRedirect = findViewById(R.id.btn_redirect) //btn_redirect
        lista = findViewById(R.id.rv_carros)
        viewPager = findViewById(R.id.view_pager)
    }

    private fun setupList() {
        val adpater = CarAdapter(CarFactory.list)
        lista.adapter = adpater
    }

    private fun setupListerners() {
        btnRedirect.setOnClickListener {
            startActivity( Intent(this, CalcularAutonomiaActivity::class.java))
        }
        tabLayout.addOnTabSelectedListener(object : TabLayout.OnTabSelectedListener{
            override fun onTabSelected(tab: TabLayout.Tab?) {
                tab?.let {
                    viewPager.currentItem = it.position
                }
            }

            override fun onTabUnselected(tab: TabLayout.Tab?) {
                //TODO("Not yet implemented")
            }

            override fun onTabReselected(tab: TabLayout.Tab?) {
                //TODO("Not yet implemented")
            }


        })

    }

    private  fun setupTabs() {
        val tabsAdapter = TabAdapter(this)
        viewPager.adapter = tabsAdapter
        //tabLayout
    }
}