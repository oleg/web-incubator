package books.proginsca.ch10.v1

import books.proginsca.ch10.v1.Element.elem
import org.scalatest.{FunSuite, Matchers}

class ElementLayoutTest extends FunSuite with Matchers {

  test("toString") {
    val element = elem(Array("00000", "00000"))
    element.toString shouldEqual
      """00000
        |00000""".stripMargin
  }

  test("above method ") {
    val element0 = elem(Array("00000", "00000"))
    val element1 = elem(Array("11111", "11111"))

    val combined = element0 above element1
    combined.toString shouldEqual
      """00000
        |00000
        |11111
        |11111""".stripMargin
  }

  test("above method different length") {
    val element0 = elem(Array("55555"))
    val element1 = elem(Array("333"))

    val combined = element0 above element1
    combined.toString shouldEqual
      """55555
        |333""".stripMargin
  }

  test("besides") {
    val element0 = elem(Array.fill(2)(">" * 5))
    val element1 = elem(Array.fill(2)("<" * 5))

    val combined = element0 besides element1
    combined.toString shouldEqual
      """>>>>><<<<<
        |>>>>><<<<<""".stripMargin
  }

}
