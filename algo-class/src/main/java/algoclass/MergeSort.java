package algoclass;

import java.util.Arrays;

public class MergeSort {

  public int[] sort(int[] input) {
    if (input == null || input.length <= 1) {
      return input;
    }

    int[] left = sort(Arrays.copyOfRange(input, 0, input.length / 2));
    int[] right = sort(Arrays.copyOfRange(input, input.length / 2, input.length));
    return merge(left, right);
  }

  private int[] merge(int[] left, int[] right) {
    final int[] result = new int[left.length + right.length];
    int li = 0;
    int ri = 0;
    for (int i = 0; i < result.length; i++) {
      if (li == left.length) {
        result[i] = right[ri++];
      } else if (ri == right.length) {
        result[i] = left[li++];
      } else if (left[li] < right[ri]) {
        result[i] = left[li++];
      } else {
        result[i] = right[ri++];
      }
    }
    return result;
  }


}
