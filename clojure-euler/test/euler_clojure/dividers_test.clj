(ns euler-clojure.dividers-test
  (:use [euler-clojure.core] :reload-all)
  (:use [clojure.test]))

(deftest dividers-test
  (is (= [] (dividers 1)))
  (is (= [2] (dividers 2)))
  (is (= [3]  (dividers 3)))
  (is (= [2 2] (dividers 4)))
  (is (= [2 11] (dividers 22)))
  (is (= [37] (dividers 37)))
  (is (= [2 2 3 3] (dividers 36)))
  (is (= [3 3 11] (dividers 99))))
  
(deftest lazy-dividers-test
  (is (= [] (lazy-dividers 1)))
  (is (= [2] (lazy-dividers 2)))
  (is (= [3]  (lazy-dividers 3)))
  (is (= [2 2] (lazy-dividers 4)))
  (is (= [2 11] (lazy-dividers 22)))
  (is (= [37] (lazy-dividers 37)))
  (is (= [2 2 3 3] (lazy-dividers 36)))
  (is (= [3 3 11] (lazy-dividers 99))))
  
(deftest dividable?-test 
  (is (dividable? 4 2))
  (is (not (dividable? 3 2)))
  (is (dividable? 22 11)))

(deftest prime?-test 
  (is (prime? 3))
  (is (not (dividable? 3 2)))
  (is (dividable? 22 11)))  
  
