package books.algsinjava.ch3

class Scoreboard(capacity: Int) {

  private val board: Array[GameEntry] = new Array[GameEntry](capacity)

  def add(newEntry: GameEntry): Unit = {

    val place = board.indexWhere(x => x == null || x.score < newEntry.score)
    if (place == -1)
    //no place for this entry
      return

    val oldEntry = board(place)
    board(place) = newEntry

    if (oldEntry != null)
      add(oldEntry) //todo not optimal

  }

  //throw index of bound, if index is wrong
  def remove(index: Int): GameEntry = ???


  def display: String =
    board
      .filter(_ != null)
      .zipWithIndex
      .map(e => s"${e._2 + 1}. ${e._1.name}")
      .mkString("\n")

}

case class GameEntry(name: String, score: Int)


