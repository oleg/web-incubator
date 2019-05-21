package books.proginsca.ch10.v2

import books.proginsca.ch10.v2.Element.elem
import org.scalatest.{FunSuite, Matchers}

class ElementTest extends FunSuite with Matchers {

  test("widen") {
    elem("00").widen(3).toString shouldEqual " 00"
    elem("111").widen(5).toString shouldEqual " 111 "

    elem("00").widen(1).toString shouldEqual "00"
    elem("2").widen(-10).toString shouldEqual "2"
  }

  test("heighten") {
    elem(Array("00", "11")).heighten(0).toString shouldEqual "00\n11"
    elem(Array("00", "11", "00")).heighten(5).toString shouldEqual "  \n00\n11\n00\n  "
  }

  test("using above with elements of different size") {
    (elem("**") above elem("----")).toString shouldEqual " ** \n----"
    (elem("o0o0o") above elem("00")).toString shouldEqual "o0o0o\n  00 "
    (elem("o0o0o") above elem("0o0")).toString shouldEqual "o0o0o\n 0o0 "
  }

  test("more complex test") {
    (elem("o0o") above elem("o0o0o") above elem("o0o")).toString shouldEqual
      elem(Array(
        " o0o ",
        "o0o0o",
        " o0o ")).toString
  }

  test("besides with different height") {
    val result = elem(Array("2", "2")) beside elem(Array("4", "4", "4", "4"))
    result.toString shouldEqual
      """ 4
        |24
        |24
        | 4""".stripMargin
  }

}
