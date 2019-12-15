require './number'
require './add'
require './variable'
require './test_setup'

class AddTest < Test::Unit::TestCase
  
  def test_add_numbers
    add = Add.new(Number.new(10), Number.new(5))
    assert_equal Number.new(15), add.evaluate({})
  end

  def test_variables
    add = Add.new(Variable.new(:x), Variable.new(:y))
    assert_equal Number.new(10), add.evaluate({x: Number.new(7), y: Number.new(3)})
  end

end
