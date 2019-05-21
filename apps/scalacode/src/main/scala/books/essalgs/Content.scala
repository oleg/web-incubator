package books.essalgs

import books.essalgs.ch2.{RandomNumbers, RandomizeArray}
import meta.{Book, Chapter}

class Content {

  new Book(
    title = "Essential Algorithms: A Practical Approach to Computer Algorithms",
    authors = List("Rod Stephens"),
    chapters = List(
      new Chapter(2, "Numerical Algorithms",
        exercises = List(
          classOf[RandomNumbers]
//          ,classOf[RandomizeArray]
        ))
    )
  )

}