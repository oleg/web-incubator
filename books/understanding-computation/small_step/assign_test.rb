require './assign'
require './add'
require './do_nothing'
require './number'
require './variable'
require './less_than'

require './test_setup'


class AssignTest < Test::Unit::TestCase
  
  def test_new
    assign = Assign.new(:x, Number.new(100))
    assert_equal :x, assign.name
    assert_equal Number.new(100), assign.expression
  end

  def test_simple_to_s
    assert_equal "y = 200", Assign.new(:y, Number.new(200)).to_s
  end
  
  def test_Add_to_s
    assert_equal "e = 9 + 3", Assign.new(:e, Add.new(Number.new(9), Number.new(3))).to_s
  end

  def test_inspect
    assert_equal "«f = 1 + 1»", Assign.new(:f, Add.new(Number.new(1), Number.new(1))).inspect
  end
  
  def test_reducible
    assert_true Assign.new(:a, Number.new(1)).reducible?
  end

  def test_reducible_complex
    assert_true Assign.new(:b, LessThan.new(Number.new(2), Number.new(1))).reducible?
  end

  def test_reduce_simple
    statement, env = Assign.new(:b, Number.new(123)).reduce({})
    assert_equal DoNothing.new, statement
    assert_equal Hash[b: Number.new(123)], env 
  end

  def test_reduce_complex
    assign = Assign.new(:c, Add.new(Number.new(123), Number.new(321)))
    env = {}
    
    assign_1, env_1 = assign.reduce(env)
    assert_equal Assign.new(:c, Number.new(444)), assign_1
    assert_equal({}, env_1)

    assign_2, env_2 = assign_1.reduce(env_1)
    assert_equal DoNothing.new, assign_2
    assert_equal Hash[c: Number.new(444)], env_2
  end
  
  def test_assign_returns_updated_environment
    statement = Assign.new(:x, Add.new(Variable.new(:x), Number.new(1)))
    environment = {x: Number.new(2)}
    
    assert_true statement.reducible?
    
    statement, environment = statement.reduce(environment)
    assert_equal Assign.new(:x, Add.new(Number.new(2), Number.new(1))), statement
    assert_equal({ x: Number.new(2) }, environment)
    
    statement, environment = statement.reduce(environment)
    assert_equal Assign.new(:x, Number.new(3)), statement
    assert_equal({ x: Number.new(2) }, environment)
    
    statement, environment = statement.reduce(environment)
    assert_equal DoNothing.new, statement
    assert_equal({ x: Number.new(3) }, environment)
  end

end
