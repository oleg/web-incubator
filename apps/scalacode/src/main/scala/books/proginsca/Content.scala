package books.proginsca

import books.essalgs.ch2.RandomNumbers
import meta.{Book, Chapter}

class Content {

  new Book(
    title = "Programming in Scala, Third Edition",
    authors = "Lex Spoon" :: "Bill Venners" :: "Martin Odersky" :: Nil,
    chapters = List(
      new Chapter(10, "Composition and Inheritance",
        exercises = List(
          classOf[RandomNumbers]
        ))
    )
  )

}