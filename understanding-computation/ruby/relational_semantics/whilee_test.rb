require './number'
require './sequence'
require './assign'
require './add'
require './whilee'
require './variable'
require './less_than'
require './boolean'

require './test_setup'

class WhileTest < Test::Unit::TestCase
  
  def test_simple
    st = While.new(LessThan.new(Variable.new(:x), Number.new(5)),
                   Sequence.new(
                     Assign.new(:x, Add.new(Variable.new(:x), Number.new(1))),
                     Assign.new(:y, Add.new(Variable.new(:y), Number.new(10)))))

    res = st.evaluate({x: Number.new(1),
                       y: Number.new(7)})
    
    assert_equal Hash[x: Number.new(5),
                      y: Number.new(47)], res
  end

end
