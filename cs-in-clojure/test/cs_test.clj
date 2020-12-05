(ns cs-test
  (:require [clojure.test :refer :all]
            [cs :refer :all]))

(deftest a-test
  (testing "naive recursive implementation."
    (is (= (fib-recur 10) 55))
    (is (= (take 11 (map #(fib-recur %) (range))) [0 1 1 2 3 5 8 13 21 34 55]))))

(deftest b-test
  (testing "implementation with memoization."
    (is (= (fib-memo 10) 55))
    (is (= (take 11 (map #(fib-memo %) (range))) [0 1 1 2 3 5 8 13 21 34 55]))))

(deftest c-test
  (testing "iterative implementation."
    (is (= (fib-iter 10) 55))
    (is (= (take 11 (map #(fib-iter %) (range))) [0 1 1 2 3 5 8 13 21 34 55]))))
