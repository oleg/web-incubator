package books.essalgs.ch2

import java.util

import scala.collection.JavaConverters._
import scala.util.Random

object RandomizeArray {

  //returns new array
  def newRandomizedArrayByJava(array: Array[Int]): Array[Int] = {
    val javaList = array.toBuffer.asJava
    util.Collections.shuffle(javaList)
    javaList.asScala.toArray
  }

  //returns new array
  def newRandomizedArrayByScala(array: Array[Int]): Array[Int] = {
    Random.shuffle(array.toBuffer).toArray
  }

  //returns same array
  def randomizeCustom(array: Array[Int]): Array[Int] = {
    val maxIndex = array.length

    for (moveFrom <- array.indices) {
      val moveTo = moveFrom + Random.nextInt(maxIndex - moveFrom)

      val tmp = array(moveFrom)
      array(moveFrom) = array(moveTo)
      array(moveTo) = tmp
    }
    array
  }

}
