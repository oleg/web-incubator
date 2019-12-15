require './evaluate'

require "./number"
require "./add"
require "./less_than"
require "./do_nothing"

require './test_setup'


class EvaluateTest < Test::Unit::TestCase
  
  def test_new
    evaluate = Evaluate.new(Add.new(Number.new(1), Number.new(100)))
    assert_equal Add.new(Number.new(1), Number.new(100)), evaluate.expression
  end

  def test_simple_to_s
    assert_equal "eval(200 + 22)", Evaluate.new(Add.new(Number.new(200), Number.new(22))).to_s
  end

  def test_inspect
    assert_equal "«eval(10 < 100)»", Evaluate.new(LessThan.new(Number.new(10), Number.new(100))).inspect
  end
  
  def test_reducible
    assert_true Evaluate.new(Number.new(1)).reducible?
  end

  def test_reduce_no_change_env
    statement = Evaluate.new(Add.new(Number.new(200), Number.new(22)))
    env = {}
    
    statement_1, env_1 = statement.reduce(env)
    assert_equal Hash[], env_1
    assert_equal Evaluate.new(Number.new(222)), statement_1
        
    statement_2, env_2 = statement_1.reduce(env)
    assert_equal Hash[], env_2
    assert_equal DoNothing.new, statement_2
  end

end  

