require './variable'
require './number'
require './test_setup'

class VariableTest < Test::Unit::TestCase

  def test_variable
    assert_equal "-> e { e[:x] }", Variable.new(:x).to_ruby
  end

  def test_eval_nil
    assert_nil eval(Variable.new(:x).to_ruby)[{}]
  end
  
  def test_eval_val
    assert_equal Number.new(7), eval(Variable.new(:y).to_ruby)[{y: Number.new(7)}]
  end
  
end
