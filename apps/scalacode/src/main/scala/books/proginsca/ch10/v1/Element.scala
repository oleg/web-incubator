package books.proginsca.ch10.v1

import books.proginsca.ch10.v1.Element.elem

object Element {
  def elem(contents: Array[String]): Element = new ArrayElement(contents)

  def elem(chr: Char, width: Int, height: Int): Element = new UniformElement(chr, width, height)

  def elem(line: String): Element = new LineElement(line)

  class ArrayElement(val contents: Array[String]) extends Element

  class LineElement(s: String) extends ArrayElement(Array(s))

  class UniformElement(ch: Char, w: Int, h: Int) extends Element {
    val contents: Array[String] = Array.fill(h)(ch.toString * w)
  }

}

abstract class Element {
  def contents: Array[String]

  def height: Int = contents.length

  def width: Int = if (height == 0) 0 else contents(0).length

  def above(that: Element): Element = elem(this.contents ++ that.contents)

  def besides(that: Element): Element = {
    elem(
      for ((line1, line2) <- this.contents zip that.contents)
        yield line1 + line2)
  }

  override def toString: String = contents.mkString("\n")
}


