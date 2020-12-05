package shen;

import util.Num;

import java.math.BigDecimal;

import static java.math.BigDecimal.ONE;
import static util.Num.checkPositive;
import static util.Num.isZero;

public class LoopPower implements Power {
  public static final BigDecimal TWO = BigDecimal.valueOf(2);

  @Override
  public BigDecimal get(BigDecimal base, BigDecimal power) {
    checkPositive(power);

    if (isZero(power)) {
      return ONE;
    }

    if (Num.equal(power, ONE)) {
      return base;
    }

//    final BigDecimal half = power.divideToIntegralValue(TWO);
//    final BigDecimal baseSquared = base.multiply(base);
//
//    final BigDecimal rs = get(baseSquared, half);
//    for() {
//
//    }
//
//
//    return rs.multiply((odd ? base : ONE));
    return null;
  }


}
