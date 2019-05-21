package algoclass;

public class Inversions {

  public static long count(int... ints) {
    final int size = ints.length;
    long inversions = 0;
    for (int i = 0; i < size; i++) {
      for (int j = i + 1; j < size; j++) {
        if (ints[i] > ints[j]) {
          inversions++;
        }
      }
    }
    return inversions;
  }
}

