(ns euler-clojure.e36
  (:use [euler-clojure.core :only (palindrome?)]))

(defn bipals [top]
  (filter #(and (palindrome? (Integer/toString %))
                (palindrome? (Integer/toBinaryString %)))
          (range top)))

(defn e36 []
  (reduce + (bipals 1000000)))