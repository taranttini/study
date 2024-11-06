package com.taranttini.cool;

import com.taranttini.cool.gof.builder.CarBuilder;

public class Cool {
    public static void main(String[] args) {
        System.out.println("Inicio");
        CarBuilder cb = new CarBuilder("vectra", 2008, 4, 4);
        cb.AddDoor()
        .AddDoor()
        .AddDoor()
        .AddDoor()
        .AddDoor()
        .AddTire()
        .AddTire()
        .AddTire()
        .RemoveDoor()
        .AddDoor()
        .AddTire()
        .SetTransmition("manual")
        .SetTransmition("automatic")
        .create();

        CarBuilder cb2 = new CarBuilder("celta", 2010, 2, 4);
        cb2.AddDoor().AddDoor()
        .AddTire().AddTire().AddTire().AddTire().AddTire()
        .SetTransmition("manual")
        .create();

    }
}