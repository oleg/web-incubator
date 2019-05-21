package books.essalgs.ch2

import org.scalatest.{FunSuite, Matchers}

class RaiseToPowerTest extends FunSuite with Matchers {

  test("java version") {
    RaiseToPowerJavaMath.power(2, 3) shouldEqual 8
  }

  test("naive version") {
    RaiseToPowerMulti.power(2, 3) shouldEqual 8
  }

  test("faster version") {
    RaiseToPowerMultiFast.power(2, 3)
  }

}
