package sort;

import static util.Arr.swap;

//2.2-2
public class SelectionSort implements IntSorter {

  @Override
  public int[] sort(int[] input) {
    for (int current = 0; current < input.length; current++) {
      int smallest = indexOfSmallest(input, current, input.length);
      swap(input, current, smallest);
    }
    return input;
  }

  private int indexOfSmallest(int[] input, int rangeFrom, int rangeTo) {
    int minIndex = rangeFrom;
    int min = input[minIndex];
    for (int i = rangeFrom; i < rangeTo; i++) {
      if (input[i] < min) {
        min = input[i];
        minIndex = i;
      }
    }
    return minIndex;
  }

}