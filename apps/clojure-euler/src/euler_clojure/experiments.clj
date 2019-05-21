(ns euler-clojure.experiments)

(defn lazy-dividers
  ([n] (lazy-dividers n 2))
  ([n divider] (lazy-seq (cond (= 1 n) []
                               (divides? n divider) (cons divider (lazy-dividers(/ n divider) divider))
                               (> divider (Math/sqrt n)) [n]
                               :else (lazy-dividers n (inc divider))))))


(defn lazy-dividers2 [num divider-val]
  (let [step (fn [n divider]
               (cond
                (> divider n) nil
                (divides? n divider) (cons divider (lazy-dividers2 (/ n divider) divider))
                :else (lazy-dividers2 n (+ 1 divider))))]
    (lazy-seq (step num divider-val))))

(defn lazy-dividers-ext [num divider-val where]
  (let [changer (if (= :up where) inc dec)
        stop? (if (= :up where) #(> % %2) (fn [a b] (= a 0)))
        step (fn [n divider]
               (cond
                (stop? divider n) nil
                (divides? n divider) (cons divider (lazy-dividers-ext (/ n divider) divider where))
                :else (lazy-dividers-ext n (changer divider) where)))]
    (lazy-seq (step num divider-val))))

(defn dividers [n]
  (let [dividersInternal
        (fn [n divider akk]
          (cond
           (> divider (Math/sqrt n))
           (cons n akk)
           (divides? n divider)
           (recur (/ n divider) divider (cons divider akk))
           :else
           (recur n (inc divider) akk)))]
    (dividersInternal n 2 [])))

(defn prime? [n]
  (-> n dividers count (= 1)))


(defn primes [n]
  (loop [a 2 num n akk []]
    (cond (= 1 num) akk
          (= 0 (mod num a)) (recur a (/ num a) (conj akk a))
          (not (= 0 (mod num a))) (recur (+ 1 a) num akk))))


(defn fast-is-prime [n]
  (loop [a 2
         num n
         akk []]
    (cond (= 1 num) akk
          (> (count akk) 1) akk
          (= 0 (mod num a)) (recur a (/ num a) (conj akk a))
          :else (recur (+ 1 a) num akk))))
