package books.essalgs.ch2

import org.scalatest.{FunSuite, Matchers}

class RandomizeArrayTest extends FunSuite with Matchers {

  test("should randomize array with java") {
    val ints = Array(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
    val randomizedInts = RandomizeArray.newRandomizedArrayByJava(ints)

    randomizedInts should contain allElementsOf Array(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
    randomizedInts should not contain inOrderElementsOf(Array(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
  }

  test("should randomize array with scala") {
    val ints = Array(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
    val randomizedInts = RandomizeArray.newRandomizedArrayByScala(Array(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))

    randomizedInts should contain allElementsOf ints
    randomizedInts should not contain inOrderElementsOf(Array(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
  }

  test("should randomize array with custom code") {
    val ints = Array(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)

    val randomizedInts = RandomizeArray.randomizeCustom(ints)

    println(randomizedInts.mkString)

    randomizedInts should contain allElementsOf Array(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
    randomizedInts should not contain inOrderElementsOf(Array(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
  }

}
