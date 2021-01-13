(ns e4
  (:use [euler-clojure.core :only (palindrome?)]))

(defn max-palindrome [num]
  (apply max (filter #(palindrome? (str %))
                     (for [i (range num)
                           j (range num)]
                       (* i j)))))

(defn e4 [] (max-palindrome 999))