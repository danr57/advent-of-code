import java.util.ArrayList;

public class MagicComputer {

    int currentIndex;
    ArrayList<Integer> intcode;
    ArrayList<Integer> originalIntcode;

    public MagicComputer(ArrayList<Integer> originalIntcode) {
        this.originalIntcode = new ArrayList<Integer>(originalIntcode);
    }

    public ArrayList<Integer> getIntcode() {
        return this.intcode;
    }

    public void resetIntcode() {
        this.intcode = new ArrayList<Integer>(originalIntcode);
    }

    public ArrayList<Integer> calculation() {
        currentIndex = 0;
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
        return intcode;
    }

    public int findNounVerb() {
        ArrayList<Integer> result;
        this.resetIntcode();
        for (int pos1 = 0; pos1 < 100; pos1++) {

            for (int pos2 = 0; pos2 < 100; pos2++) {

                this.intcode.set(1, pos1);
                this.intcode.set(2, pos2);

                result = new ArrayList<>(this.calculation());

                if (result.get(0) == 19690720) {
                    return (100 * intcode.get(1) + intcode.get(2));
                } else {
                    this.resetIntcode();
                }

            }
        }
        return 404;
    }
}