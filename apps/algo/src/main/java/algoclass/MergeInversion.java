package algoclass;

import java.util.Arrays;

public class MergeInversion {

  public static long count(int... input) {
    return countInvAndSort(input).l;
  }

  public static Pair<Long, int[]> countInvAndSort(int[] input) {
    if (input == null || input.length <= 1) {
      return Pair.from(0L, input);
    }

    final Pair<Long, int[]> left = countInvAndSort(Arrays.copyOfRange(input, 0, input.length / 2));
    final Pair<Long, int[]> right = countInvAndSort(Arrays.copyOfRange(input, input.length / 2, input.length));
    final Pair<Long, int[]> between = countInvAndMerge(left.r, right.r);
    return Pair.from(between.l + left.l + right.l, between.r);
  }

  private static Pair<Long, int[]> countInvAndMerge(int[] left, int[] right) {
    final int[] result = new int[left.length + right.length];
    long invCount = 0;
    int li = 0;
    int ri = 0;
    for (int i = 0; i < result.length; i++) {
      if (li == left.length) {
        result[i] = right[ri++];
      }//
      else if (ri == right.length) {
        result[i] = left[li++];
      }//
      else if (left[li] < right[ri]) {
        result[i] = left[li++];
      }//
      else {
        result[i] = right[ri++];
        invCount += left.length - li;
      }
    }
    return Pair.from(invCount, result);
  }

}
