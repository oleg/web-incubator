import maze.Main;
import org.hyperskill.hstest.v4.stage.MainMethodTest;
import org.hyperskill.hstest.v4.testcase.CheckResult;
import org.hyperskill.hstest.v4.testcase.TestCase;

import java.util.ArrayList;
import java.util.List;


class Clue {
    int height;
    int width;
    Clue(int h, int w) {
        height = h;
        width = w;
    }
}

public class MazeRunnerTest extends MainMethodTest<Clue> {

    public MazeRunnerTest() throws Exception {
        super(Main.class);
    }

    @Override
    public List<TestCase<Clue>> generateTestCases() {
        return List.of(
            new TestCase<Clue>()
                .setInput("6 8")
                .setAttach(new Clue(6, 8)),

            new TestCase<Clue>()
                .setInput("15 65")
                .setAttach(new Clue(15, 65))
        );
    }

    private List<String> getMaze(String reply) {

        List<String> maze = new ArrayList<>();
        String[] rows = reply.split("\n");

        for (String row : rows) {
            boolean possibleMazeRow = true;
            boolean haveSpecialSymbol = false;
            for (char c : row.toCharArray()) {
                if (c == '█') {
                    haveSpecialSymbol = true;
                }
                if (c != '█' && c != ' ') {
                    possibleMazeRow = false;
                    break;
                }
            }
            if (haveSpecialSymbol && possibleMazeRow) {
                maze.add(row);
            }
        }

        return maze;
    }

    @Override
    public CheckResult check(String reply, Clue clue) {

        List<String> maze = getMaze(reply);

        if (maze.size() != clue.height) {
            return new CheckResult(false,
                "Number of rows in the maze is incorrect");
        }

        int columnsLength = maze.get(0).length();

        if (columnsLength / 2 != clue.width) {
            return new CheckResult(false,
                "Number of columns in the maze is incorrect");
        }

        for (String row : maze) {
            int columnLength = row.length();
            if (columnLength != columnsLength) {
                return new CheckResult(false,
                    "Number of columns " +
                        "should be the same on every row");
            }
        }

        return CheckResult.TRUE;
    }
}
