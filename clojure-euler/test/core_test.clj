(ns core-test
  (:use [core] :reload-all)
  (:use [clojure.test]))

;; todo fix me
;(deftest by-pair-test
;  (testing "base cases"
;    (is (= [] (by-pair [])))
;    (is (= [[0 nil]] (by-pair [0])))
;    (is (= [[0 1]] (by-pair [0 1])))
;    (is (= [[0 1] [1 2]] (by-pair [0 1 2])))))
;
;125874
;(deftest multiplied-by?-test
;  (is (= [5 10 15 20 25] (multiplied-by? 5 5)))
;  (is (= [12 24 36] (multiplied-by? 12 3))))
;
;(deftest same-digit-test
;  (is (= false (same-digits? 12 22)))
;  (is (= true (same-digits? 152 125))))
;
;(deftest same-digit-test
;  (is (= [125874] (like52 2))))
  
;(deftest lazy-dividers-test
;  (is (= [2 2 3] (lazy-dividers 12 2))))
;
;(deftest lazy-dividers-ext-test
;  (is (= [2 2 3] (lazy-dividers-ext 12 2 :up))))
;  ; (is (= [2 2 3] (lazy-dividers-ext 12 2 :down))))
;
;
;
