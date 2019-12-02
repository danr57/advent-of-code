package main;

class FuelCounter {

    int fuelForModule(int i) {
        int fuel = ((i / 3) - 2);
        if (fuel > 0) {
            return fuel + fuelForModule(fuel);
        }
        return 0;
    }
}