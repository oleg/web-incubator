package books.essalgs.ch2

import org.scalatest.FunSuite

class RandomNumbersTest extends FunSuite {

  test("should generate same random number sequence for the same seed") {
    val random = new RandomNumbers()

    assert(tenTimes.map(_ => random.next) == List(5, 7, 10, 9, 2, 8, 6, 3, 4, 0))
    assert(tenTimes.map(_ => random.next) == List(5, 7, 10, 9, 2, 8, 6, 3, 4, 0))
  }

  test("generate for different seeds") {
    val m13 = new RandomNumbers(9, 19, 15, 13)
    assert(tenTimes.map(_ => m13.next) == Vector(4, 0, 2, 1, 8, 11, 3, 7, 5, 6))

    val m12312 = new RandomNumbers(0, 20, 14, 12312)
    assert(tenTimes.map(_ => m12312.next) == Vector(14, 294, 5894, 7086, 6302, 2934, 9446, 4254, 11222, 2838))
  }

  test("non interesting cases") {
    val x0 = new RandomNumbers(0, 20, 20, 5)
    assert(tenTimes.map(_ => x0.next) == Vector(0, 0, 0, 0, 0, 0, 0, 0, 0, 0))

    val x1 = new RandomNumbers(1, 20, 20, 5)
    assert(tenTimes.map(_ => x1.next) == Vector(0, 0, 0, 0, 0, 0, 0, 0, 0, 0))

    val x2 = new RandomNumbers(2, 20, 20, 5)
    assert(tenTimes.map(_ => x2.next) == Vector(0, 0, 0, 0, 0, 0, 0, 0, 0, 0))
  }

  private def tenTimes: Range.Inclusive = 1 to 10
}
