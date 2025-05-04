package com.taranttini.cool.gof.builder;

public class CarBuilder {

    String model;
    Integer year;

    Integer doors = 0;
    Integer tires = 0;
    String transmition = "";
    Integer maximumDoor;
    Integer maximumTires;

    public  CarBuilder(String model, Integer year, Integer maximumDoor, Integer maximumTires) {
        this.model = model;
        this.year = year;
        this.maximumDoor = maximumDoor;
        this.maximumTires = maximumTires;
    }

    public CarBuilder AddDoor()
    {
        System.out.println("Adding door...");
        if (this.doors == maximumDoor) {
            System.out.println("ERROR: Cars has maximum doors: " + this.doors);
            return this;    
        }
        this.doors += 1;
        System.out.println("Added Door");
        System.out.println("Car has " + this.doors + " door(s)");

        return this;
    }

    public CarBuilder AddTire()
    {
        System.out.println("Adding tires...");
        if (this.tires == maximumTires) {
            System.out.println("ERROR: Cars has maximum tires: " + this.tires);
            return this;    
        }
        this.tires += 1;
        System.out.println("Added Tire");
        System.out.println("Car has " + this.tires + " tire(s)");

        return this;
    }

    public CarBuilder RemoveDoor()
    {
        System.out.println("Removing doors...");
        if (this.tires == 0) {
            System.out.println("ERROR: Cars no has more doors");
            return this;    
        }
        this.tires -= 1;
        System.out.println("Removed door");
        System.out.println("Car has " + this.doors + " door(s)");

        return this;
    }

    public CarBuilder RemoveTire()
    {
        System.out.println("Removing tires...");
        if (this.tires == 0) {
            System.out.println("ERROR: Cars no has more tires");
            return this;    
        }
        this.tires -= 1;
        System.out.println("Removed tire");
        System.out.println("Car has " + this.tires + " tire(s)");

        return this;
    }

    public CarBuilder SetTransmition(String transmition) {

        this.transmition = transmition;
        System.out.println("Now car has transmition: " + transmition);
        return this;
    }
	public void create() {

		System.out.println("\n== creating car ==");
        System.out.println("Model: " + model);
        System.out.println("Year: " + year);

        if (this.maximumDoor != doors) {
            System.out.println("Doors: ERROR, car doors has only " + doors);    
        } else {
        System.out.println("Doors: " + doors);
        }
        if (this.maximumTires != tires) {
            System.out.println("Tires: ERROR, car tires has only " + tires);    
        } else {

            System.out.println("Tires: " + tires);
        }
        System.out.println("Transmition: " + transmition);
        System.out.println("==================\n");
	}
	
}
