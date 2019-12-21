require "./add"
require "./number"
require "./multiply"

require './test_setup'


class AddTest < Test::Unit::TestCase

  def test_add_to_s
    assert_equal "7 + 9", Add.new(Number.new(7), Number.new(9)).to_s
  end

  def test_add_inspect
    assert_equal "«33 + 44»", Add.new(Number.new(33), Number.new(44)).inspect
  end

  def test_add
    a = Add.new(Number.new(1), Number.new(2))
    assert_equal 1, a.left.value
    assert_equal 2, a.right.value
  end
  
  def test_add_multiply
    a = Add.new(Multiply.new(Number.new(1), Number.new(2)),
                Multiply.new(Number.new(3), Number.new(4)))
    assert_equal 1, a.left.left.value
    assert_equal 4, a.right.right.value
  end
  
  def test_add_reducible
    assert_true Add.new(Number.new(100), Number.new(1)).reducible?
  end

  def test_add_reduce_no
    add = Add.new(Number.new(1),
                  Number.new(2))
    assert_equal Number.new(3), add.reduce({})
  end

  def test_add_reduce_left
    add = Add.new(Add.new(Number.new(4), Number.new(2)),
                  Number.new(1))
    
    add_1 = add.reduce({})
    assert_equal Add.new(Number.new(6), Number.new(1)), add_1
    
    add_2 = add_1.reduce({})
    assert_equal Number.new(7), add_2
  end
  
  def test_add_reduce_right
    add = Add.new(Number.new(7),
                  Add.new(Number.new(3), Number.new(1)))

    add_1 = add.reduce({})
    assert_equal Add.new(Number.new(7), Number.new(4)), add_1

    add_2 = add_1.reduce({})
    assert_equal Number.new(11), add_2
  end
  
  def test_add_reduce_left_and_right
    add = Add.new(Add.new(Number.new(1), Number.new(2)),
                  Add.new(Number.new(3), Number.new(4)))

    add_1 = add.reduce({})
    assert_equal Add.new(Number.new(3), Add.new(Number.new(3), Number.new(4))), add_1

    add_2 = add_1.reduce({})
    assert_equal Add.new(Number.new(3), Number.new(7)), add_2

    add_3 = add_2.reduce({})
    assert_equal Number.new(10), add_3
  end

  def test_add_mixed_reduce
    add = Add.new(Add.new(Number.new(1), Number.new(2)),
                  Multiply.new(Number.new(3), Number.new(4)))
    
    add_1 = add.reduce({})
    assert_equal Add.new(Number.new(3), Multiply.new(Number.new(3), Number.new(4))), add_1
    
    add_2 = add_1.reduce({})
    assert_equal Add.new(Number.new(3), Number.new(12)), add_2

    add_3 = add_2.reduce({})
    assert_equal Number.new(15), add_3
  end  
end
