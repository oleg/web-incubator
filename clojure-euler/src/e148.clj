(ns e148)

(defn next-pascal [l] (map + (cons 0 l) (conj l 0)))

;;(defn pascal [n]
;;  (case n 
;;    2 [1 1]
;;    (next-pascal (vec (pascal (- n 1))))))

(defn divnum [l]
  (count (filter #(not (= 0 (mod % 7))) l)))

(defn pascal [n]
  (pascal7 1 n [1 1] 0))

(defn pascal7 [cr n sd old-n]
  (if (< cr n)
    (let [line (next-pascal (vec sd))
	  new-n (divnum line)]
      (recur (inc cr) n line (+ old-n new-n)))
    old-n))


;;tests
;;(assert (= (pascal 0) []))
;;(assert (= (pascal 1) [1]))
;;(assert (= (pascal 2) [1 1]))
;;(assert (= (pascal 3) [1 2 1]))
;;(assert (= (pascal 4) [1 3 3 1]))
;;(assert (= (pascal 5) [1 4 6 4 1]))