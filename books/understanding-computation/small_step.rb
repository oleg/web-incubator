class Number < Struct.new(:value)
  
  def to_s
    value.to_s
  end
  
  def inspect
    "«#{self}»"
  end

  def reducible?
    false
  end
  
end

class Boolean < Struct.new(:value)
  
  def to_s
    value.to_s
  end
  
  def inspect
    "«#{self}»"
  end

  def reducible?
    false
  end
  
end

class Add < Struct.new(:left, :right)
  
  def to_s
    "#{left} + #{right}"
  end
  
  def inspect
    "«#{self}»"
  end

  def reducible?
    true
  end
  
  def reduce
    if left.reducible?
      Add.new(left.reduce, right)
    elsif right.reducible?
      Add.new(left, right.reduce)
    else
      Number.new(left.value + right.value)
    end  
  end

end


class Multiply < Struct.new(:left, :right)
  
  def to_s
    "#{left} * #{right}"
  end
  
  def inspect
    "«#{self}»"
  end

  def reducible?
    true
  end
  
  def reduce
    if left.reducible?
      Multiply.new(left.reduce, right)
    elsif right.reducible?
      Multiply.new(left, right.reduce)
    else
      Number.new(left.value * right.value)
    end
  end

end

class LessThan < Struct.new(:left, :right)
  
  def to_s
    "#{left} < #{right}"
  end
  
  def inspect
    "«#{self}»"
  end

  def reducible?
    true
  end
  
  def reduce
    if left.reducible?
      LessThan.new(left.reduce, right)
    elsif right.reducible?
      LessThan.new(left, right.reduce)
    else
      Boolean.new(left.value < right.value)
    end
  end

end


class Machine < Struct.new(:expression)

  def run
    while expression.reducible?
      puts expression
      step
    end
    puts expression
  end

  def step
    self.expression = expression.reduce
  end

end



require "test/unit"


class NumberTest < Test::Unit::TestCase
  
  def test_new
    n = Number.new 100
    assert_equal 100, n.value
  end

  def test_to_s
    assert_equal "99", Number.new(99).to_s
  end

  def test_inspect
    assert_equal "«10»", Number.new(10).inspect
  end
  
  def test_number
    assert_false Number.new(100).reducible?
  end
  
end

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
    assert_equal Number.new(3), add.reduce
  end

  def test_add_reduce_left
    add = Add.new(Add.new(Number.new(4), Number.new(2)),
                  Number.new(1))
    
    add_1 = add.reduce
    assert_equal Add.new(Number.new(6), Number.new(1)), add_1
    
    add_2 = add_1.reduce
    assert_equal Number.new(7), add_2
  end
  
  def test_add_reduce_right
    add = Add.new(Number.new(7),
                  Add.new(Number.new(3), Number.new(1)))

    add_1 = add.reduce
    assert_equal Add.new(Number.new(7), Number.new(4)), add_1

    add_2 = add_1.reduce
    assert_equal Number.new(11), add_2
  end
  
  def test_add_reduce_left_and_right
    add = Add.new(Add.new(Number.new(1), Number.new(2)),
                  Add.new(Number.new(3), Number.new(4)))

    add_1 = add.reduce
    assert_equal Add.new(Number.new(3), Add.new(Number.new(3), Number.new(4))), add_1

    add_2 = add_1.reduce
    assert_equal Add.new(Number.new(3), Number.new(7)), add_2

    add_3 = add_2.reduce
    assert_equal Number.new(10), add_3
  end

  def test_add_mixed_reduce
    add = Add.new(Add.new(Number.new(1), Number.new(2)),
                  Multiply.new(Number.new(3), Number.new(4)))
    
    add_1 = add.reduce
    assert_equal Add.new(Number.new(3), Multiply.new(Number.new(3), Number.new(4))), add_1
    
    add_2 = add_1.reduce
    assert_equal Add.new(Number.new(3), Number.new(12)), add_2

    add_3 = add_2.reduce
    assert_equal Number.new(15), add_3
  end  
end


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
    
    assert_equal Number.new(2), multiply.reduce
  end

  def test_multiply_reduce_left
    multiply = Multiply.new(Multiply.new(Number.new(4), Number.new(2)),
                            Number.new(1))

    multiply_1 = multiply.reduce
    assert_equal Multiply.new(Number.new(8), Number.new(1)), multiply_1

    multiply_2 = multiply_1.reduce
    assert_equal Number.new(8), multiply_2
  end
  
  def test_multiply_reduce_right
    multiply = Multiply.new(Number.new(7),
                            Multiply.new(Number.new(3), Number.new(1)))

    multiply_1 = multiply.reduce
    assert_equal Multiply.new(Number.new(7), Number.new(3)), multiply_1

    multiply_2 = multiply_1.reduce
    assert_equal Number.new(21), multiply_2
  end

  def test_multiply_reduce_left_and_right
    multiply = Multiply.new(Multiply.new(Number.new(1), Number.new(2)),
                            Multiply.new(Number.new(3), Number.new(4)))

    multiply_1 = multiply.reduce
    assert_equal Multiply.new(Number.new(2), Multiply.new(Number.new(3), Number.new(4))), multiply_1

    multiply_2 = multiply_1.reduce
    assert_equal Multiply.new(Number.new(2), Number.new(12)), multiply_2

    multiply_3 = multiply_2.reduce    
    assert_equal Number.new(24), multiply_3
  end

  def test_multiply_mixed_reduce
    multiply = Multiply.new(Add.new(Number.new(1), Number.new(2)),
                            Multiply.new(Number.new(3), Number.new(4)))

    multiply_1 = multiply.reduce
    assert_equal Multiply.new(Number.new(3),
                              Multiply.new(Number.new(3), Number.new(4))), multiply_1

    multiply_2 = multiply_1.reduce
    assert_equal Multiply.new(Number.new(3), Number.new(12)), multiply_2
    
    multiply_3 = multiply_2.reduce
    assert_equal Number.new(36), multiply_3
  end
  
end

class MachineTest < Test::Unit::TestCase

  def test_machine_add
    machine = Machine.new(Add.new(Number.new(1),
                                  Add.new(Number.new(2),
                                          Add.new(Number.new(3),
                                                  Number.new(4)))))

    out = capture_output { machine.run }[0]
    expected = <<-eos
1 + 2 + 3 + 4
1 + 2 + 7
1 + 9
10
eos
    assert_equal expected, out
    assert_equal Number.new(10), machine.expression
  end

  def test_machine_less_than
    machine = Machine.new(LessThan.new(Add.new(Number.new(1), Number.new(2)),
                                       Add.new(Number.new(3), Number.new(4))))

    out = capture_output { machine.run }[0]
    expected = <<-eos
1 + 2 < 3 + 4
3 < 3 + 4
3 < 7
true
eos
    assert_equal expected, out
    assert_equal Boolean.new(true), machine.expression
  end
  
end


class BooleanTest < Test::Unit::TestCase
 
  def test_new_true
    assert_equal true, Boolean.new(true).value
  end
  
  def test_new_false
    assert_equal false, Boolean.new(false).value
  end

  def test_to_s
    assert_equal "true", Boolean.new(true).to_s
    assert_equal "false", Boolean.new(false).to_s
  end

  def test_inspect
    assert_equal "«true»", Boolean.new(true).inspect
    assert_equal "«false»", Boolean.new(false).inspect    
  end
  
  def test_number
    assert_false Boolean.new(true).reducible?
    assert_false Boolean.new(false).reducible?
  end
  
end


class LessThanTest < Test::Unit::TestCase

  def test_less_than_new
    lt = LessThan.new(Number.new(7), Number.new(9))
    assert_equal Number.new(7), lt.left
    assert_equal Number.new(9), lt.right
  end

  def test_to_s
    assert_equal "2 < 4", LessThan.new(Number.new(2), Number.new(4)).to_s
    assert_equal "3 < 1", LessThan.new(Number.new(3), Number.new(1)).to_s
  end

  def test_inspect
    assert_equal "«2 < 4»", LessThan.new(Number.new(2), Number.new(4)).inspect
    assert_equal "«3 < 1»", LessThan.new(Number.new(3), Number.new(1)).inspect
  end

  def test_reducible
    lt = LessThan.new(Add.new(Number.new(1), Number.new(1)),
                      Add.new(Number.new(2), Number.new(2)))
    
    assert_true lt.reducible?
  end

  def test_reduce_true
    lt = LessThan.new(Add.new(Number.new(1), Number.new(1)),
                      Add.new(Number.new(2), Number.new(2)))

    lt_1 = lt.reduce
    assert_equal LessThan.new(Number.new(2),
                             Add.new(Number.new(2), Number.new(2))), lt_1

    lt_2 = lt_1.reduce
    assert_equal LessThan.new(Number.new(2), Number.new(4)), lt_2

    lt_3 = lt_2.reduce
    assert_equal Boolean.new(true), lt_3
  end

  def test_reduce_false
    assert_equal Boolean.new(false), LessThan.new(Number.new(5), Number.new(1)).reduce
  end

end


_q = "«»"
