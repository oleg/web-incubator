(ns euler-clojure.e2
  (:use [euler-clojure.core :only (fib-seq)]))

(defn e2 []
  (reduce + (filter even? (take-while #(< % 4000000) fib-seq))))