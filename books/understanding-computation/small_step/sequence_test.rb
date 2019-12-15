require './sequence'
require './number'
require './assign'
require './add'
require './do_nothing'
require './variable'

require './test_setup'


class SequenceTest < Test::Unit::TestCase
  
  def test_new
    seq = Sequence.new(Assign.new(:x, Number.new(10)),
                       Assign.new(:y, Number.new(11)))
    assert_equal Assign.new(:x, Number.new(10)), seq.first
    assert_equal Assign.new(:y, Number.new(11)), seq.second
  end

  def test_to_s
    assert_equal "b = 1; a = 2", Sequence.new(Assign.new(:b, Number.new(1)),
                                              Assign.new(:a, Number.new(2))).to_s
  end

  def test_inspect
    assert_equal "«a = 10; b = 20»", Sequence.new(Assign.new(:a, Number.new(10)),
                                                  Assign.new(:b, Number.new(20))).inspect
  end
  
  def test_reducible
    assert_true Sequence.new(Assign.new(:a, Number.new(10)),
                             Assign.new(:b, Number.new(20))).reducible?
  end

  def test_always_reducible
    assert_true Sequence.new(DoNothing.new, DoNothing.new).reducible?
  end

  def test_reduce
    seq = Sequence.new(Assign.new(:x, Add.new(Number.new(1), Number.new(1))),
                       Assign.new(:y, Add.new(Variable.new(:x), Number.new(3))))
    env = {}
    seq_1, env_1 = seq.reduce(env)
    assert_equal Sequence.new(Assign.new(:x, Number.new(2)),
                              Assign.new(:y, Add.new(Variable.new(:x), Number.new(3)))), seq_1
    
    seq_2, env_2 = seq_1.reduce(env_1)
    assert_equal Sequence.new(DoNothing.new,
                              Assign.new(:y, Add.new(Variable.new(:x), Number.new(3)))), seq_2

    seq_3, env_3 = seq_2.reduce(env_2)
    assert_equal Assign.new(:y, Add.new(Variable.new(:x), Number.new(3))), seq_3
  end
  

end
