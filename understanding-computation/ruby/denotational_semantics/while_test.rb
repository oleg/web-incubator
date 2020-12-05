require './number'
require './while'
require './assign'
require './variable'
require './add'
require './less_than'

require './test_setup'

class WhileTest < Test::Unit::TestCase

  def setup
    @w = While.new(LessThan.new(Variable.new(:x), Number.new(10)),
                   Assign.new(:x, Add.new(Variable.new(:x), Number.new(1))))
  end

  def test_to_ruby
    assert_equal "-> e { while (-> e { (-> e { e[:x] })[e] < (-> e { 10 })[e] })[e]; e = (-> e { e.merge({ :x => (-> e { (-> e { e[:x] })[e] + (-> e { 1 })[e] })[e] }) })[e]; end; e }",
      @w.to_ruby
  end
  
  def test_eval
    env = {x: 9}
    env_1 = eval(@w.to_ruby)[env]
    assert_equal Hash[x: 10], env_1
  end

end
