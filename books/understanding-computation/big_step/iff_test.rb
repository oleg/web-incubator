require './number'
require './boolean'
require './assign'
require './iff'

require './test_setup'

class IfTest < Test::Unit::TestCase
  
  def test_true
    st = If.new(Boolean.new(true),
                Assign.new(:x, Number.new(2)),
                Assign.new(:y, Number.new(3)))

    assert_equal Hash[x: Number.new(2)], st.evaluate({})
  end

  def test_false
    st = If.new(Boolean.new(false),
                Assign.new(:x, Number.new(2)),
                Assign.new(:y, Number.new(3)))
    
    assert_equal Hash[y: Number.new(3)], st.evaluate({})
  end

end
