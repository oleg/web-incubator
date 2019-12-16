require './test_setup'
require './whilee'

require './less_than'
require './assign'
require './variable'
require './multiply'
require './add'
require './number'
require './sequence'
require './iff'
require './do_nothing'
require './boolean'

class WhileTest < Test::Unit::TestCase

  #while (x < 5) { x = x * 3 }
  def test_to_s
    w = While.new(LessThan.new(Variable.new(:x), Number.new(5)),
                  Assign.new(:x, Multiply.new(Variable.new(:x), Number.new(3))))

    assert_equal "while (x < 5) { x = x * 3 }", w.to_s
  end

  def test_inspect
    w = While.new(LessThan.new(Variable.new(:x), Number.new(10)),
                  Assign.new(:x, Add.new(Variable.new(:x), Number.new(2))))
    
    assert_equal "«while (x < 10) { x = x + 2 }»", w.inspect    
  end

  def test_reducible
    w = While.new(LessThan.new(Variable.new(:x), Number.new(10)),
                  Assign.new(:x, Add.new(Variable.new(:x), Number.new(4))))
    
    assert_true w.reducible?
  end

  def test_reduce
    c = LessThan.new(Variable.new(:x), Number.new(8))
    b = Assign.new(:x, Add.new(Variable.new(:x), Number.new(2)))
    w = While.new(c, b)
    s = Sequence.new(b, w)
    env = {x: Number.new(7)}

    w_1, env_1 = w.reduce(env)
    assert_equal Hash[x: Number.new(7)], env_1
    assert_equal If.new(c, s, DoNothing.new), w_1

    w_2, env_2 = w_1.reduce(env_1)
    assert_equal Hash[x: Number.new(7)], env_2
    assert_equal If.new(LessThan.new(Number.new(7), Number.new(8)), s, DoNothing.new), w_2

    w_3, env_3 = w_2.reduce(env_2)
    assert_equal Hash[x: Number.new(7)], env_3
    assert_equal If.new(Boolean.new(true), s, DoNothing.new), w_3
    
    w_4, env_4 = w_3.reduce(env_3)
    assert_equal Hash[x: Number.new(7)], env_4
    assert_equal s, w_4

    w_5, env_5 = w_4.reduce(env_4)
    assert_equal Hash[x: Number.new(7)], env_5
    assert_equal Sequence.new(Assign.new(:x, Add.new(Number.new(7), Number.new(2))), w), w_5

    w_6, env_6 = w_5.reduce(env_5)
    assert_equal Hash[x: Number.new(7)], env_6
    assert_equal Sequence.new(Assign.new(:x, Number.new(9)), w), w_6

    w_7, env_7 = w_6.reduce(env_6)
    assert_equal Hash[x: Number.new(9)], env_7
    assert_equal Sequence.new(DoNothing.new, w), w_7

    w_8, env_8 = w_7.reduce(env_7)
    assert_equal Hash[x: Number.new(9)], env_8
    assert_equal w, w_8

    w_9, env_9 = w_8.reduce(env_8)
    assert_equal Hash[x: Number.new(9)], env_9
    assert_equal If.new(LessThan.new(Variable.new(:x), Number.new(8)), s, DoNothing.new), w_9

    w_10, env_10 = w_9.reduce(env_9)
    assert_equal Hash[x: Number.new(9)], env_10
    assert_equal If.new(LessThan.new(Number.new(9), Number.new(8)), s, DoNothing.new), w_10

    w_11, env_11 = w_10.reduce(env_10)
    assert_equal Hash[x: Number.new(9)], env_11
    assert_equal If.new(Boolean.new(false), s, DoNothing.new), w_11

    w_12, env_12 = w_11.reduce(env_11)
    assert_equal Hash[x: Number.new(9)], env_12
    assert_equal DoNothing.new, w_12
  end

end
