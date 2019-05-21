package books.essalgs.ch2

object RecursiveGcd {
  def calc(a: Int, b: Int): Int = {
    if (a == 0 || b == 0) throw new IllegalArgumentException
    gcdRecursive(a, b)
  }

  private def gcdRecursive(a: Int, b: Int): Int = {
    val reminder = a % b
    if (reminder == 0) b
    else gcdRecursive(b, reminder)
  }
}

object IterGcd {
  def calc(ax: Int, bx: Int): Int = {
    if (ax == 0 || bx == 0) throw new IllegalArgumentException
    var (a, b) = (ax, bx)
    while (b != 0) {
      val rem = a % b
      a = b
      b = rem
    }
    a
  }
}

object IterGcdWithLambda {
  def calc(ax: Int, bx: Int): Int = {
    if (ax == 0 || bx == 0) throw new IllegalArgumentException
    var (a, b) = (ax, bx)
    val update = (na: Int, nb: Int) => {
      a = na; b = nb
    }
    while (b != 0) {
      update(b, a % b)
    }
    a
  }
}
