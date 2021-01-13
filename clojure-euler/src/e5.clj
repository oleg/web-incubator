(ns e5
  (:use [euler-clojure.core :only (dividers)]))

(defn count-values [m]
  (into {} (for [[k, v] m]
             [k, (count v)])))

(def dividers-up-21
  (for [i (range 1 21)]
    (dividers i)))

(defn lst []
  (for [i dividers-up-21]
    (count-values (group-by identity i))))


(def take-longest (apply merge-with #(if (> % %2) % %2) (lst)))

(defn e5 []
    (int (reduce * (map #(Math/pow (key %) (val %)) take-longest))))
