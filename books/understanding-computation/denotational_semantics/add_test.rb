require './number'
require './add'
require './test_setup'

class AddTest < Test::Unit::TestCase

  def test_add
    add = Add.new(Number.new(1), Number.new(2)).to_ruby
    assert_equal "-> e { (-> e { 1 })[e] + (-> e { 2 })[e] }", add
  end
  
  def test_eval_new
    add = Add.new(Number.new(7), Number.new(8))
    assert_equal 15, eval(add.to_ruby)[nil]
  end

end
