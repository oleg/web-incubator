package books.essalgs.ch2

class RandomNumbers(private var x: Int = 0,
                    a: Int = 7,
                    b: Int = 5,
                    m: Int = 11) {

  def next: Int = {
    x = calc
    x
  }

  private def calc = (x * a + b) % m

  override def toString: String = s"PRNG {$x, $a, $b, $m}"
}
