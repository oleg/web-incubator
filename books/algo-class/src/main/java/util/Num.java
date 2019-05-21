package util;

import java.math.BigDecimal;

import static java.math.BigDecimal.ZERO;

public class Num {
  public static boolean less(BigDecimal a, BigDecimal b) {
    return a.compareTo(b) < 0;
  }

  public static boolean lessOrEqual(BigDecimal a, BigDecimal b) {
    return a.compareTo(b) <= 0;
  }

  public static boolean equal(BigDecimal a, BigDecimal b) {
    return a.compareTo(b) == 0;
  }

  public static boolean greater(BigDecimal a, BigDecimal b) {
    return a.compareTo(b) > 0;
  }

  public static boolean greaterOrEqual(BigDecimal a, BigDecimal b) {
    return a.compareTo(b) >= 0;
  }

  public static boolean notEqual(BigDecimal a, BigDecimal b) {
    return a.compareTo(b) != 0;
  }

  public static boolean isZero(BigDecimal a) {
    return equal(a, BigDecimal.ZERO);
  }

  public static boolean isNotZero(BigDecimal a) {
    return notEqual(a, BigDecimal.ZERO);
  }


  public static void checkPositive(BigDecimal power) {
    if (Num.less(power, ZERO)) {
      throw new IllegalArgumentException();
    }
  }

}
