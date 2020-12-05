require './number'
require './multiply'
require './test_setup'

class MultiplyTest < Test::Unit::TestCase

  def test_to_ruby
    m = Multiply.new(Number.new(1), Number.new(2)).to_ruby
    assert_equal "-> e { (-> e { 1 })[e] * (-> e { 2 })[e] }", m
  end
  
  def test_eval
    m = Multiply.new(Number.new(7), Number.new(8))
    assert_equal 56, eval(m.to_ruby)[nil]
  end

end
