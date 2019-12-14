require './number'
require './sequence'
require './assign'
require './add'
require './variable'

require './test_setup'

class SequenceTest < Test::Unit::TestCase
  
  def test_simple
    st = Sequence.new(Assign.new(:x, Number.new(2)),
                      Assign.new(:y, Number.new(3)))

    assert_equal Hash[x: Number.new(2),
                      y: Number.new(3)], st.evaluate({})
  end

  def test_complex
    st = Sequence.new(Assign.new(:x, Number.new(2)),
                      Assign.new(:x, Add.new(Variable.new(:x), Number.new(3))))

    assert_equal Hash[x: Number.new(5)], st.evaluate({})
  end

end
