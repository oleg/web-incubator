package books.essalgs.ch2

object RaiseToPowerJavaMath {
  def power(num: Double, power: Int): Double = math.pow(num, power)
}

object RaiseToPowerMulti {
  //todo check power >= 1
  def power(num: Double, power: Int): Double = {
    var result = num
    for (_ <- 1 until power) { //O(n)
      result *= num
    }
    result
  }
}

object RaiseToPowerMultiFast {
  //todo check args
  def power(num: Double, power: Int): Double = {
    //2, 3, 4, 5, 6, 7
    ???
    var result = num
    for (_ <- 1 until power) {
      result *= num
    }
    result
  }
}

