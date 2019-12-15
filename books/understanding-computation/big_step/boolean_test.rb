require './boolean'
require './test_setup'

class BooleanTest < Test::Unit::TestCase

  def test_evaluate_true
    assert_equal Boolean.new(true), Boolean.new(true).evaluate({})
  end
  
  def test_evaluate_false
    assert_equal Boolean.new(false), Boolean.new(false).evaluate({})
  end
  
end
