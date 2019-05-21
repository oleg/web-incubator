package shen;

import util.Num;

import java.math.BigDecimal;

import static java.math.BigDecimal.ONE;
import static java.math.BigDecimal.ZERO;
import static util.Num.checkPositive;

public class LinearPower implements Power {

  public BigDecimal get(BigDecimal base, BigDecimal power) {
    checkPositive(power);

    BigDecimal result = ONE;
    BigDecimal i = ZERO;
    while (Num.less(i, power)) {
      i = i.add(ONE);
      result = result.multiply(base);
    }
    return result;
  }
}