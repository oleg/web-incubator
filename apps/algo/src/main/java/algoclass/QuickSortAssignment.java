package algoclass;

import sort.IntSorter;
import util.Arr;

import static java.lang.Math.max;

public class QuickSortAssignment implements IntSorter {
  private long comparisonCount;
  private final PivotStrategy pivotStrategy;

  public QuickSortAssignment(PivotStrategy pivotStrategy) {
    this.pivotStrategy = pivotStrategy;
  }

  @Override
  public int[] sort(int[] input) {
    if (input == null) {
      return input;
    }
    sort(input, 0, input.length);
    return input;
  }

  void sort(int[] input, int from, int to) {
    if (to - from <= 1) {
      return;
    }

    int position = partition(input, from, to);
    sort(input, from, position);
    sort(input, position + 1, to);
  }

  int partition(int[] input, int from, int to) {
    comparisonCount += to - from - 1;

    int pivot = pivotStrategy.pivot(input, from, to);

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

  public long getComparisonCount() {
    return comparisonCount;
  }

  public enum PivotStrategy implements PivotStrategyService {
    FIRST {
      public int pivot(int[] input, int from, int to) {
        return input[from];
      }
    },
    FINAL {
      public int pivot(int[] input, int from, int to) {
        Arr.swap(input, from, to - 1);
        return input[from];
      }
    },
    MEDIAN {
      public int pivot(int[] input, int from, int to) {
        int f = from;
        int t = to - 1;
        final int m = (f + t) / 2;

        int fv = input[f];
        int tv = input[t];
        int mv = input[m];

        if (fv > max(mv, tv)) {
          if (mv > tv) {
            Arr.swap(input, f, m);
            return mv;
          }
          Arr.swap(input, f, t);
          return tv;
        }
        if (mv > max(fv, tv)) {
          if (fv > tv) {
//            Arr.swap(input, f, f);
            return fv;
          }
          Arr.swap(input, f, t);
          return tv;
        }
        if (fv > mv) {
//          Arr.swap(input, f, f);
          return fv;
        }
        Arr.swap(input, f, m);
        return mv;
      }
    }
  }

  private interface PivotStrategyService {
    int pivot(int[] input, int from, int to);
  }
}
