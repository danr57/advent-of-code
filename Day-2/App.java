import java.io.File;
import java.io.FileNotFoundException;
import java.util.ArrayList;
import java.util.Scanner;

public class App {

    public static void main(String[] args) throws FileNotFoundException {

        Scanner sc = new Scanner(new File("Day-2/input")).useDelimiter(",");

        ArrayList<Integer> originalIntcode = new ArrayList<Integer>();

        while (sc.hasNextInt()) {
            originalIntcode.add(sc.nextInt());
        }
        sc.close();

        MagicComputer steve = new MagicComputer(originalIntcode);

        System.out.println("NounVerb: " + steve.findNounVerb());
    }
}