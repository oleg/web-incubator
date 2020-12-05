require './number'
require './less_than'
require './add'
require './variable'
require './test_setup'

class LessThanTest < Test::Unit::TestCase

  def test_to_ruby
    lt = LessThan.new(Number.new(1), Number.new(2)).to_ruby
    assert_equal "-> e { (-> e { 1 })[e] < (-> e { 2 })[e] }", lt
  end

  def test_to_ruby_complex
    lt = LessThan.new(Add.new(Variable.new(:x), Number.new(1)), Number.new(3)).to_ruby
    assert_equal "-> e { (-> e { (-> e { e[:x] })[e] + (-> e { 1 })[e] })[e] < (-> e { 3 })[e] }", lt
  end
  
  def test_complex_eval
    lt = LessThan.new(Add.new(Variable.new(:x), Number.new(1)), Number.new(3)).to_ruby
    assert_false eval(lt)[{x: 777}]
  end

  def test_eval_false
    lt = LessThan.new(Number.new(7), Number.new(8))
    assert_true eval(lt.to_ruby)[nil]
  end

  def test_eval_true
    lt = LessThan.new(Number.new(8), Number.new(4))
    assert_false eval(lt.to_ruby)[nil]
  end


end
