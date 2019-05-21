package books.proginsca.ch10.v1

import books.proginsca.ch10.v1.Element.elem
import org.scalatest.{FunSuite, Matchers}

class ElementTest extends FunSuite with Matchers {
  
  test("empty element") {
    val e = elem(Array[String]())

    e.contents shouldEqual Array()
    e.width shouldEqual 0
    e.height shouldEqual 0

  }

  test("should create array element") {
    val e = elem(Array("Hello", "World"))

    e.contents shouldEqual Array("Hello", "World")
    e.width shouldEqual 5
    e.height shouldEqual 2
  }

  test("line element 1") {
    val e = elem("foo")

    e.contents shouldEqual Array("foo")
    e.width shouldEqual 3
    e.height shouldEqual 1
  }

  test("uniform element v1") {
    val e = elem('a', 2, 2)

    e.contents shouldEqual Array("aa", "aa")
    e.width shouldEqual 2
    e.height shouldEqual 2
  }

}
