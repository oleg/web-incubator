require './variable'
require './number'
require './boolean'

require './test_setup'

class VariableTest < Test::Unit::TestCase
  def test_new
    v = Variable.new :x
    assert_equal :x, v.name
  end

  def test_to_s
    assert_equal "yy", Variable.new(:yy).to_s
  end

  def test_inspect
    assert_equal "«zzz»", Variable.new(:zzz).inspect
  end
  
  def test_reducible
    assert_true Variable.new(:xyz).reducible?
  end
  
  def test_reduce
    assert_equal Number.new(100), Variable.new(:y).reduce({y: Number.new(100)})
    assert_equal Number.new(200), Variable.new(:y).reduce({y: Number.new(200)})
    assert_equal Boolean.new(true), Variable.new(:z).reduce({z: Boolean.new(true)})
  end

end
