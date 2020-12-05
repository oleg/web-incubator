require './number'
require './boolean'
require './variable'
require './test_setup'

class VariableTest < Test::Unit::TestCase
  
  def test_found_number
    assert_equal Number.new(9), Variable.new(:x).evaluate({x: Number.new(9)})
  end
  
  def test_found_boolean
    assert_equal Boolean.new(false), Variable.new(:x).evaluate({x: Boolean.new(false)})
  end
    
  def test_not_found
    assert_nil Variable.new(:z).evaluate({y: Number.new(10)})
  end
  
end
