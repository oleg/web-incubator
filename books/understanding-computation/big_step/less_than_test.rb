require './number'
require './multiply'
require './add'
require './boolean'
require './variable'
require './less_than'
require './test_setup'

class LessThanTest < Test::Unit::TestCase
  
  def test_true
    lt = LessThan.new(Number.new(1), Number.new(2))
    assert_equal Boolean.new(true), lt.evaluate({})
  end

  def test_false
    l = Multiply.new(Variable.new(:x), Number.new(100))
    r = Add.new(Variable.new(:y), Number.new(100))
    lt = LessThan.new(l, r)
    assert_equal Boolean.new(false), lt.evaluate({x: Number.new(7), y: Number.new(3)})
  end

end
