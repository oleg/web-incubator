package meta

class Meta {

}

class Chapter(number: Int, title: String, exercises: List[Unit] = List())

class Book(title: String, authors: List[String], chapters: List[Chapter])