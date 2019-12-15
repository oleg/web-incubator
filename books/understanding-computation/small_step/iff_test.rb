require './iff'
require './number'
require './less_than'
require './assign'
require './boolean'

require './test_setup'


class IfTest < Test::Unit::TestCase

  def test_to_s
    iff = If.new(LessThan.new(Number.new(10), Number.new(11)),
                 Assign.new(:x, Number.new(100)),
                 Assign.new(:y, Number.new(200)))
    
    assert_equal "if (10 < 11) { x = 100 } else { y = 200 }", iff.to_s
  end

  def test_inspect
    iff = If.new(Boolean.new(true),
                 Assign.new(:x, Number.new(1)),
                 Assign.new(:x, Number.new(2)))
    
    assert_equal "«if (true) { x = 1 } else { x = 2 }»", iff.inspect
  end

  def test_reducible
    iff = If.new(Boolean.new(true),
                 Assign.new(:x, Number.new(1)),
                 Assign.new(:x, Number.new(2)))

    assert_true iff.reducible?
  end
  
  def test_reduce_true
    st = If.new(LessThan.new(Number.new(10), Number.new(11)),
                Assign.new(:x, Number.new(100)),
                Assign.new(:y, Number.new(200)))
    env = {}

    st_1, env_1 = st.reduce(env)
    assert_equal Hash[], env_1
    assert_equal If.new(Boolean.new(true),
                        Assign.new(:x, Number.new(100)),
                        Assign.new(:y, Number.new(200))), st_1

    st_2, env_2 = st_1.reduce(env_1)
    assert_equal Hash[], env_2
    assert_equal Assign.new(:x, Number.new(100)), st_2
  end

  def test_reduce_false
    st = If.new(LessThan.new(Number.new(11), Number.new(10)),
                Assign.new(:x, Number.new(100)),
                Assign.new(:y, Number.new(200)))
    env = {}

    st_1, env_1 = st.reduce(env)
    assert_equal Hash[], env_1
    assert_equal If.new(Boolean.new(false),
                        Assign.new(:x, Number.new(100)),
                        Assign.new(:y, Number.new(200))), st_1

    st_2, env_2 = st_1.reduce(env_1)
    assert_equal Hash[], env_2
    assert_equal Assign.new(:y, Number.new(200)), st_2
  end

end
