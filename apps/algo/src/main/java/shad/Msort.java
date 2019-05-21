package shad;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class Msort {

  public int[] mergesort(int[] array) {
    if (array.length == 1) {
      return array;
    } else {
      int middle = (array.length / 2);
      int[] left = Arrays.copyOfRange(array, 0, middle);
      int[] right = Arrays.copyOfRange(array, middle, array.length);

      left = mergesort(left);
      right = mergesort(right);
      return merge(left, right);
    }
  }

  private int[] merge(int[] left, int[] right) {
    final List<Integer> result = new ArrayList<>(left.length + right.length);
    while (left.length > 0 && right.length > 0) {
      if (left[0] < right[0]) {
        result.add(left[0]);
        left = Arrays.copyOfRange(left, 1, left.length);
      } else {
        result.add(right[0]);
        right = Arrays.copyOfRange(right, 1, right.length);
      }
    }
    if (left.length > 0) {
      append(left, result);

    }
    if (right.length > 0) {
      append(right, result);
    }
    return toArr(result);
  }

  private int[] toArr(List<Integer> result) {
    final int[] res = new int[result.size()];
    for (int i = 0; i < result.size(); i++) {
      res[i] = result.get(i);
    }
    return res;
  }

  private void append(int[] right, List<Integer> result) {
    for (int i : right) {
      result.add(i);
    }
  }
}
