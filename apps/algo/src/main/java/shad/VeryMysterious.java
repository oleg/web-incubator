package shad;

public class VeryMysterious {

  public static int[] invoke(String strArray) {
    final String[] split = strArray.split(" ");
    final int[] array = new int[split.length];
    for (int i = 0; i < array.length; i++) {
      array[i] = Integer.parseInt(split[i]);
    }
    return invoke(array);
  }

  public static int[] invoke(int[] array) {
    int[] result = new int[array.length];

    for (int i = 0; i < array.length; i++) {
      result[array[i]] = i;
    }
    return result;
  }

}
