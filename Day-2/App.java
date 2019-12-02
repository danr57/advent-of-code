import java.io.File;
import java.io.FileNotFoundException;
import java.util.ArrayList;
import java.util.Scanner;

public class App {
    public static void main(String[] args) throws FileNotFoundException {

        Scanner sc = new Scanner(new File("input")).useDelimiter(",");

        ArrayList<Integer> intcode = new ArrayList<Integer>();

        while (sc.hasNextInt()) {
            intcode.add(sc.nextInt());
        }
        sc.close();

        new MagicComputer(intcode);
    }
}