(ns euler-clojure.e7
  (:use [euler-clojure.core :only (prime?)]))

(defn e7 []
  (take 1 (drop 10000 (filter prime? (iterate inc 1)))))