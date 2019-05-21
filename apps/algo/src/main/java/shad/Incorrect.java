package shad;

import java.util.Arrays;

import static util.Arr.swap;

public class Incorrect {
  //каково минимальное n, что существует такая перестановка чисел от 0 до n-1, на которой алгоритм работает бесконечно?
  public static void invoke(Integer[] array) {
    int[] rs = new int[array.length];
    for (int i = 0; i < array.length; i++) {
      rs[i] = array[i];
    }
    invoke(rs);
  }

  public static void invoke(int[] array) {
    int i = 0;
    while (i < array.length) {
      System.out.println(Arrays.toString(array));
      if (array[i] > i) {
        int j = i;
        while ((j < array.length) && (array[j] >= j)) {
          j++;
        }
        swap(array, i, j);
        i = 0;
      } else {
        i++;
      }
    }

  }
}
