package shad;

import java.util.ArrayList;
import java.util.List;

import static java.util.Arrays.asList;

public class Permutations {

  public static <T> List<List<T>> get(T... args) {
    final boolean loopVersion = false;
    if (loopVersion) {
      return loopGet(args);
    } else {
      return recursiveGet(args);
    }
  }


  public static <T> List<List<T>> recursiveGet(T[] args) {
    final List<T> list = new ArrayList<>();
    list.add(args[0]);

    List<List<T>> result = new ArrayList<>();
    result.add(list);

    for (int i = 1; i < args.length; i++) {
      result = permute(args[i], result);
    }
    return result;
  }

  private static <T> List<List<T>> permute(T other, List<List<T>> original) {
    final List<List<T>> result = new ArrayList<>();
    for (List<T> previous : original) {
      final int size = original.get(0).size();

      for (int i = 0; i <= size; i++) {
        final List<T> permuted = new ArrayList<>(previous);
        permuted.add(i, other);
        result.add(permuted);
      }

    }
    return result;
  }

  //TODO doesn't work
  public static <T> List<List<T>> loopGet(T... args) {
    if (args.length == 1) {
      return asList(asList(args));
    }
    final List<List<T>> result = new ArrayList<>();

    ArrayList<T> base = new ArrayList<>(asList(args));

    for (int j = 0; j < args.length; j++) {
      for (int i = 0; i < args.length - 1; i++) {
        result.add(base);
        base = new ArrayList<>(base);
        forward(base, i);
      }
    }
    return result;
  }

  private static <T> void forward(ArrayList<T> base, int i) {
    final T a = base.get(i);
    final T b = base.get(i + 1);

    base.set(i, b);
    base.set(i + 1, a);
  }

}
