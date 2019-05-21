package books.algsinjava

import books.algsinjava.ch3.Scoreboard
import meta.{Book, Chapter}

class Content {

  new Book(
    title = "Data Structures and Algorithms in Java, 6th Edition",
    authors = List("Roberto Tamassia", "Michael T. Goodrich"),
    chapters = List(

      new Chapter(3, "Fundamental Data Structures", exercises = List({

        ("3.1.1 STORING GAME ENTRIES IN AN ARRAY", classOf[Scoreboard])

      }))
    ))
}

