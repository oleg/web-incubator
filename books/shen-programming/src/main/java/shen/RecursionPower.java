package shen;


import java.math.BigDecimal;

import static java.math.BigDecimal.ONE;
import static util.Num.checkPositive;
import static util.Num.isZero;

public class RecursionPower implements Power {

  public static final BigDecimal TWO = BigDecimal.valueOf(2);

  @Override
  public BigDecimal get(BigDecimal base, BigDecimal power) {
    final BigDecimal remainder = power.remainder(TWO);
    return recursion(base, power, remainder.compareTo(ONE) == 0);
  }

  private BigDecimal recursion(BigDecimal base, BigDecimal power, boolean odd) {
    checkPositive(power);

    if (isZero(power)) {
      return ONE;
    }

    final BigDecimal half = power.divideToIntegralValue(TWO);
    final BigDecimal baseSquared = base.multiply(base);

    final BigDecimal rs = get(baseSquared, half);

    return rs.multiply((odd ? base : ONE));

  }


}
