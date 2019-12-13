require './number'

require "test/unit"
require "test/unit/assertions"
Test::Unit::Assertions.use_pp = false


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
