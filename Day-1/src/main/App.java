package main;

import java.io.File;
import java.io.FileNotFoundException;
import java.util.Scanner;

public class App {
    public static void main(String[] args) throws FileNotFoundException {
        FuelCounter steve = new FuelCounter();

        int totalFuel = 0;

        Scanner scanner = new Scanner(new File("input"));
        int[] fuel = new int[100];
        int i = 0;
        while (scanner.hasNextInt()) {
            fuel[i++] = scanner.nextInt();
        }

        for (int x : fuel) {
            totalFuel += steve.fuelForModule(x);

        }
        System.out.println("Total fuel: " + totalFuel);
    }
}