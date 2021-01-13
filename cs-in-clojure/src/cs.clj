(ns cs)

(defn fib-recur
  "naive recursive implementation"
  [n]
  (if (< n 2)
    n
    (+
      (fib-recur (- n 1))
      (fib-recur (- n 2)))))

(defn fib-memo
  "calculates recursively with memoization"
  ([n]
   (fib-memo n (atom {})))

  ([n memo]
   (let [calced-val (@memo n)]
     (if calced-val
       calced-val
       (let [result (if (< n 2)
                      n
                      (+
                        (fib-memo (- n 1) memo)
                        (fib-memo (- n 2) memo)))]
         (swap! memo #(assoc % n result))
         result)))))

(defn fib-iter
  "iterative version"
  [n]
  (if (= n 0)
    0
    (loop [prev 0
           last 1
           counter (dec n)]
      (if (= counter 0)
        last
        (recur last
               (+ prev last)
               (dec counter))))))
