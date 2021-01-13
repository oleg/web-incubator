(ns e6)

(defn e6 []
  (let [sum-of-pows (reduce + (for [i (range 101)] (* i i)))
        sum (reduce + (for [i (range 101)] i))]
    (- (* sum sum) sum-of-pows)))
