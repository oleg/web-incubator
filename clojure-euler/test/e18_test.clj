(ns e18-test
  (:use [e18] :reload-all)
  (:use [clojure.test]))

(deftest solve-test
  (is (= 15 (count data))))

(def test-data1 [[:a]
		 [:b1 :b2]])

(def test-data2 [[:a]
		 [:b1 :b2]
		 [:c1 :c2 :c3]])

(def test-data3 [[:b1 :b2]
		  [:c1 :c2 :c3]])

(def test-data4 [[[[:b1 :a]] [[:b2 :a]]]
		  [:c1 :c2 :c3]])

(def test-data5 [[:a] [:b1 :b2] [:c1 :c2 :c3] [:d1 :d2 :d3 :d4]])

(deftest solve-2-test
  (is (= [[[:b1 :a]] [[:b2 :a]]]
	   (solve (first test-data1) (rest test-data1))))
  
  (is (= []
	   (solve (first test-data2) (rest test-data2))))
  
  (is (= [[[:c1 :b1]] [[:c2 :b1] [:c2 :b2]] [[:c3 :b2]]]
	   (solve (first test-data3) (rest test-data3))))
  
  (is (= []
	   (solve (first test-data4) (rest test-data4))))
  
  (is (= []
	   (solve (first test-data5) (rest test-data5)))))

(deftest append-to-each-test
  (is (= [[:h 1 2] [:h 4 5]] (append-to-each :h [[1 2] [4 5]])))
  (is (= [ [[:h 1] [:h 5]] [[:h 6] [:h 8]] ] (append-to-each :h [ [[1] [5]] [[6] [8]] ])))
  (is (= [[:h 1] [:h 2] [:h 3]] (append-to-each :h [[1] [2] [3]]))))

(deftest ll-test
  (is (ll? [[1]]))
  (is (ll? [[1 2] [2]]))
  (is (not (ll? [[[1]]])))
  (is (not (ll? [1 3])))
  (is (not (ll? 1))))
    
(deftest adding-test
  (is (= [[[:c1 :b1]] [[:c2 :b1] [:c2 :b2]] [[:c3 :b2]]] (adding [:b1 :b2] [:c1 :c2 :c3])))
  (is (= [[[:b1 :a1]] [[:b2 :a1]]] (adding [:a1] [:b1 :b2]))))