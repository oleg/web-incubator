package books.proginsca.ch10.v2

import books.proginsca.ch10.v2.Element.elem

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

  def above(that: Element): Element = {
    val e1 = this widen that.width
    val e2 = that widen this.width
    elem(e1.contents ++ e2.contents)
  }

  def beside(that: Element): Element = {
    val e1 = this heighten that.height
    val e2 = that heighten this.height
    elem(
      for ((line1, line2) <- e1.contents zip e2.contents)
        yield line1 + line2)
  }

  def heighten(h: Int): Element = {
    if (h <= height)
      return this
    val linesToAdd = h - this.height
    val aboveLines = linesToAdd / 2
    val belowLines = linesToAdd - aboveLines
    elem(' ', width, aboveLines) above this above elem(' ', width, belowLines)
  }

  def widen(w: Int): Element = {
    if (w <= width)
      return this
    val spaceToAdd = w - this.width
    val rightSpan = spaceToAdd / 2
    val leftSpan = spaceToAdd - rightSpan
    elem(' ', leftSpan, 1) beside this beside elem(' ', rightSpan, 1)
  }

  override def toString: String = contents.mkString("\n")
}


