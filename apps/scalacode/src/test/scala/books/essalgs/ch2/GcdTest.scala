package books.essalgs.ch2

import org.scalatest.{FunSuite, Matchers}

class GcdTest extends FunSuite with Matchers {

  test("recursive gcd") {
    doAllTests(RecursiveGcd.calc)
  }

  test("loop gcd") {
    doAllTests(IterGcd.calc)
  }

  test("loop gcd v2") {
    doAllTests(IterGcdWithLambda.calc)
  }

  private def doAllTests(gcd: (Int, Int) => Int) = {
    gcd(100, 10) should equal(10)
    gcd(21, 14) should equal(7)
    gcd(14, 21) should equal(7)
    gcd(4581, 3003) should equal(3)
    gcd(10, 20) should equal(10)
    gcd(7, 100) should equal(1)

    an[IllegalArgumentException] should be thrownBy gcd(0, 10)
    an[IllegalArgumentException] should be thrownBy gcd(10, 0)
  }
}
