package maze;

import java.util.Random;

public class Main {
    public static void main(String[] args) {
        int[][] maze = generate();
        System.out.println(printMe(maze));
    }

    private static int[][] generate() {
        Random random = new Random();
        int[][] maze = new int[20][20];
        for (int[] rows : maze) {
            for (int i = 0; i < rows.length; i++) {
                rows[i] = random.nextInt(2);
            }
        }

        return maze;
    }

    public static String printMe(int[][] maze) {
        String response = "";
        for (int[] ints : maze) {
            for (int c : ints) {
                if (c == 1) {
                    response += '\u2588' + "" + '\u2588';
                } else {
                    response += "  ";
                }
            }
            response += '\n';
        }
        return response;
    }
}
