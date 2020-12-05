(ns euler-clojure.e52)

(defn same-chars? [a b]
  (= (sort (.toString a))
     (sort (.toString b))))

(defn multiply-each [num seq]
  (for [i seq] (* i num)))

(defn same-digits? [input]
  (let [pairs (drop-last (partition 2 1 nil input))]
    (every? #(apply same-chars? %) pairs)))

(defn e52 []
  (let [mul-range (range 2 7)]
    (take 1 (filter #(same-digits? (multiply-each % mul-range)) (iterate inc 1)))))