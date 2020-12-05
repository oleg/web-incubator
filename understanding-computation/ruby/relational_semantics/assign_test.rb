require './number'
require './assign'
require './variable'
require './add'
require './multiply'
require './less_than'
require './boolean'

require './test_setup'

class AssignTest < Test::Unit::TestCase
  
  def test_simple
    assert_equal Hash[x: Number.new(100)], Assign.new(:x, Number.new(100)).evaluate({})
  end

  def test_complex
    a = Assign.new(:y, LessThan.new(Add.new(Number.new(3), Variable.new(:x)),
                                    Multiply.new(Number.new(2), Number.new(3))))
    
    assert_equal Hash[x: Number.new(1), y: Boolean.new(true)], a.evaluate({x: Number.new(1)})
  end

end
