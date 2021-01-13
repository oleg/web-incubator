(ns e1
  (:use [core :only (divides?)]))

(defn e1 []
  (reduce + (for [i (range 999 0 -1)
                  :when (or (divides? i 3)
                            (divides? i 5))]
              i)))
