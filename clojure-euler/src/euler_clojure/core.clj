(ns euler-clojure.core)

(defn divides? [v n]
  (= 0 (mod v n)))

(defn dividers [num]
  (loop [n num
         divider 2
         akk []]
    (cond (= 1 n) akk
          (divides? n divider) (recur (/ n divider) divider (conj akk divider))
          (> divider (Math/sqrt n)) (conj akk n)
          :else (recur n (inc divider) akk))))

(defn max-prime [num]
  (loop [n num
         divider 2]
    (cond (= 1 n) divider
          (divides? n divider) (recur (/ n divider) divider)
          (> divider (Math/sqrt n)) n
          :else (recur n (+ 1 divider)))))

;(defn prime? [n]
;  (-> n dividers count (= 1)))

(defn prime? [n]
  (= n (max-prime n)))


(def fib-seq
     (lazy-cat [1 2]
               (map + fib-seq (rest fib-seq))))

(defn palindrome? [input]
    (let [length (count input)
          [begin end] (split-at (/ length 2) input)]
      (if (odd? length)
        (= (rest (reverse begin)) end)
        (= (reverse begin) end))))

(defn digits [& rest]
  (sort
   (reduce concat (map #(.toString %) rest))))

(def pandigits [\1 \2 \3 \4 \5 \6 \7  \8 \9])

(defn pandigital? [a b]
  (= pandigits (digits a b (* a b))))

(defn p [border]
  (for [i (range 0 border)
        j (range i border)
          :when (pandigital? i j)]
    [i, j]))



(defn shift [s]
  (let [[h & t] s]
    (conj (vec t) h)))

(defn circular [data]
  (let [circular-internal
        (fn [n count akk]
          (if (= 0 count)
            akk
            (let [shifted (shift n)]
              (recur shifted (dec count) (cons shifted akk)))))]
    (circular-internal data (count data) [])))

(defn circular-prime? [n]
  (reduce #(and % %2)
          (map (comp prime? #(Integer/parseInt %) #(apply str %))
               (circular (.toString n)))))

(defn euler35 []
  (-> (filter circular-prime? (range 2 1000000)) count))
