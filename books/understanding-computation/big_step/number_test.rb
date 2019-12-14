require './number'
require './test_setup'

class NumberTest < Test::Unit::TestCase

  def test_evaluate
    assert_equal Number.new(7), Number.new(7).evaluate({})
    assert_equal Number.new(100), Number.new(100).evaluate({})
  end
  
end
