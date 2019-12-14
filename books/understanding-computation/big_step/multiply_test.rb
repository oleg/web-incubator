require './number'
require './multiply'
require './add'
require './variable'
require './test_setup'

class MultiplyTest < Test::Unit::TestCase
  
  def test_multiply_numbers
    m = Multiply.new(Number.new(10), Number.new(5))
    assert_equal Number.new(50), m.evaluate({})
  end

  def test_variables
    m = Multiply.new(Variable.new(:x), Variable.new(:y))
    assert_equal Number.new(21), m.evaluate({x: Number.new(7), y: Number.new(3)})
  end

  def test_multiply_and_add_numbers
    a5 = Add.new(Number.new(3), Number.new(2))
    a6 = Add.new(Number.new(2), Number.new(4))
    m = Multiply.new(a5, a6)
    assert_equal Number.new(30), m.evaluate({})
  end

end
