require './multiply'
require './add'
require "./number"

require "test/unit"
require "test/unit/assertions"
Test::Unit::Assertions.use_pp = false


class MultiplyTest < Test::Unit::TestCase
  
  def test_multiply
    a = Multiply.new(Number.new(1), Number.new(2))
    assert_equal 1, a.left.value
    assert_equal 2, a.right.value
  end
  
  def test_multiply_to_s
    assert_equal "5 * 2", Multiply.new(Number.new(5), Number.new(2)).to_s
  end

  def test_multiply_inspect
    assert_equal "«12 * 7»", Multiply.new(Number.new(12), Number.new(7)).inspect
  end
  
  def test_multiply_reducible
    assert_true Multiply.new(Number.new(5), Number.new(6)).reducible?
  end

  def test_multiply_reduce_no
    multiply = Multiply.new(Number.new(1), Number.new(2))
    
    assert_equal Number.new(2), multiply.reduce({})
  end

  def test_multiply_reduce_left
    multiply = Multiply.new(Multiply.new(Number.new(4), Number.new(2)),
                            Number.new(1))

    multiply_1 = multiply.reduce({})
    assert_equal Multiply.new(Number.new(8), Number.new(1)), multiply_1

    multiply_2 = multiply_1.reduce({})
    assert_equal Number.new(8), multiply_2
  end
  
  def test_multiply_reduce_right
    multiply = Multiply.new(Number.new(7),
                            Multiply.new(Number.new(3), Number.new(1)))

    multiply_1 = multiply.reduce({})
    assert_equal Multiply.new(Number.new(7), Number.new(3)), multiply_1

    multiply_2 = multiply_1.reduce({})
    assert_equal Number.new(21), multiply_2
  end

  def test_multiply_reduce_left_and_right
    multiply = Multiply.new(Multiply.new(Number.new(1), Number.new(2)),
                            Multiply.new(Number.new(3), Number.new(4)))

    multiply_1 = multiply.reduce({})
    assert_equal Multiply.new(Number.new(2), Multiply.new(Number.new(3), Number.new(4))), multiply_1

    multiply_2 = multiply_1.reduce({})
    assert_equal Multiply.new(Number.new(2), Number.new(12)), multiply_2

    multiply_3 = multiply_2.reduce({})
    assert_equal Number.new(24), multiply_3
  end

  def test_multiply_mixed_reduce
    multiply = Multiply.new(Add.new(Number.new(1), Number.new(2)),
                            Multiply.new(Number.new(3), Number.new(4)))

    multiply_1 = multiply.reduce({})
    assert_equal Multiply.new(Number.new(3),
                              Multiply.new(Number.new(3), Number.new(4))), multiply_1

    multiply_2 = multiply_1.reduce({})
    assert_equal Multiply.new(Number.new(3), Number.new(12)), multiply_2
    
    multiply_3 = multiply_2.reduce({})
    assert_equal Number.new(36), multiply_3
  end
  
end
