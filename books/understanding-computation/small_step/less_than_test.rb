require './less_than'

require "./number"
require "./boolean"
require "./add"

require "test/unit"
require "test/unit/assertions"
Test::Unit::Assertions.use_pp = false


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

    lt_1 = lt.reduce({})
    assert_equal LessThan.new(Number.new(2),
                              Add.new(Number.new(2), Number.new(2))), lt_1

    lt_2 = lt_1.reduce({})
    assert_equal LessThan.new(Number.new(2), Number.new(4)), lt_2

    lt_3 = lt_2.reduce({})
    assert_equal Boolean.new(true), lt_3
  end

  def test_reduce_false
    assert_equal Boolean.new(false), LessThan.new(Number.new(5), Number.new(1)).reduce({})
  end

end
