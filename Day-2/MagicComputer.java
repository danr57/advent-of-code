import java.util.ArrayList;

public class MagicComputer {

    int currentIndex;
    ArrayList<Integer> intcode;

    public MagicComputer(ArrayList<Integer> intcode) {
        this.intcode = intcode;
        System.out.println("Array before manip: " + this.printSnip());

        while (currentIndex < intcode.size()) {

            if (intcode.get(currentIndex) == 1) {
                intcode.set(intcode.get(currentIndex + 3),
                        (intcode.get(intcode.get(currentIndex + 1)) + intcode.get(intcode.get(currentIndex + 2))));
                this.currentIndex += 4;

            } else if (intcode.get(currentIndex) == 2) {
                intcode.set(intcode.get(currentIndex + 3),
                        (intcode.get(intcode.get(currentIndex + 1)) * intcode.get(intcode.get(currentIndex + 2))));
                this.currentIndex += 4;

            } else if (intcode.get(currentIndex) == 99) {
                break;
            } else {
                System.out.println("And i oop");
                break;
            }
        }
        System.out.println("Array after manip : " + this.printSnip());
    }

    public String printSnip() {
        String result = "";
        for (int i = 0; i < intcode.size(); i++) {
            result += intcode.get(i) + " ";
        }
        return result + "\n";
    }
}