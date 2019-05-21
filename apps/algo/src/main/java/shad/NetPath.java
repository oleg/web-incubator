package shad;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashSet;
import java.util.List;

public class NetPath {
  private final List<List<Direction>> permutations;

  public NetPath(int width, int height) {
    Direction[] arr = new Direction[width + height];
    Arrays.fill(arr, 0, width, Direction.LEFT);
    Arrays.fill(arr, width, width + height, Direction.DOWN);
    final List<List<Direction>> lists = Permutations.get(arr);
    permutations = new ArrayList<>(new HashSet<>(lists));
  }

  public List<List<Direction>> directions() {
    return permutations;
  }

  public int count() {
    return permutations.size();
  }

  public static enum Direction {
    LEFT, DOWN
  }


}
