package books.algsinjava.ch3

import org.scalatest.FunSuite

class ScoreboardTest extends FunSuite {

  test("empty scoreboard should display nothing") {
    val scoreboard = new Scoreboard(10)
    assert(scoreboard.display == "")
  }

  test("first player should take first place") {
    val scoreboard = new Scoreboard(4)

    scoreboard.add(GameEntry("Till", 100))

    assert(scoreboard.display == "1. Till")
  }

  test("second player should take second place if his score is less") {
    val scoreboard = new Scoreboard(4)

    scoreboard.add(GameEntry("Till", 100))
    scoreboard.add(GameEntry("Richard", 40))

    assert(scoreboard.display == "1. Till\n2. Richard")
  }

  test("second player should take first place if his score is bigger") {
    val scoreboard = new Scoreboard(4)

    scoreboard.add(GameEntry("Till", 100))
    scoreboard.add(GameEntry("Richard", 200))

    assert(scoreboard.display == "1. Richard\n2. Till")
  }

  test("should not add element if board full and score is too small") {
    val scoreboard = new Scoreboard(2)

    scoreboard.add(GameEntry("Till", 100))
    scoreboard.add(GameEntry("Richard", 200))
    scoreboard.add(GameEntry("Oliver", 50))

    assert(scoreboard.display == "1. Richard\n2. Till")
  }

  test("should add element if board full and score is higher") {
    val scoreboard = new Scoreboard(2)

    scoreboard.add(GameEntry("Till", 100))
    scoreboard.add(GameEntry("Richard", 200))
    scoreboard.add(GameEntry("Oliver", 250))

    assert(scoreboard.display == "1. Oliver\n2. Richard")
  }


}
