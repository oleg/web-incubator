(cons 'chicken 'cat)
(cons 'pork '(beef chicken))

(car '(a b c d e f))
;; A
(cdr '(a b c d e f))
;;(B C D E F)
(cadddr '(a b c d e f))
;; D
(cdddr '(a b c d e f))
;; (D E F)
(list 'a 'b 'c)
;; (A B C)
