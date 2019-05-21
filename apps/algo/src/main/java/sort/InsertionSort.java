package sort;

public class InsertionSort implements BiDirectionalIntSorter {

  @Override
  public int[] sort(int[] input) {
    for (int currId = 1; currId < input.length; currId++) {
      int element = input[currId];
      int previousId = currId - 1;

      while (previousId >= 0 && input[previousId] > element) {
        input[previousId + 1] = input[previousId];
        previousId--;
      }
      input[previousId + 1] = element;
    }
    return input;
  }

  @Override
  public int[] sortDesc(int[] input) {
    for (int current = 1; current < input.length; current++) {
      int element = input[current];
      int previous = current - 1;

      while (previous >= 0 && input[previous] < element) {
        input[previous + 1] = input[previous];
        previous--;
      }
      input[previous + 1] = element;
    }
    return input;
  }

}
