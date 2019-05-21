package algoclass;

import sort.IntSorter;
import util.Arr;

public class QuickSort implements IntSorter {

  @Override
  public int[] sort(int[] input) {
    if (input == null) {
      return input;
    }
    return sort(input, 0, input.length);
  }

  int[] sort(int[] input, int from, int to) {
    if (to - from <= 1) {
      return input;
    }

    int position = partition(input, from, to);
    sort(input, from, position);
    sort(input, position + 1, to);
    return input;
  }

  static int partition(int[] input, int from, int to) {
    //swap(input, from, to - 1);
    int pivot = input[from];
    int border = from + 1;
    int next = from + 1;
    while (next < to) {
      if (input[next] < pivot) {
        Arr.swap(input, border, next);
        border++;
      }
      next++;
    }
    Arr.swap(input, from, border - 1);
    return border - 1;
  }

}
