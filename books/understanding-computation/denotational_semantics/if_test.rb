require './number'
require './if'
require './assign'
require './less_than'

require './test_setup'

class IfTest < Test::Unit::TestCase

  def setup
    @iff = If.new(LessThan.new(Number.new(1), Number.new(2)),
                  Assign.new(:x, Number.new(10)),
                  Assign.new(:y, Number.new(20)))
  end

  def test_to_ruby
    assert_equal "-> e { 
if (-> e { (-> e { 1 })[e] < (-> e { 2 })[e] })[e] 
then (-> e { e.merge({ :x => (-> e { 10 })[e] }) })[e] 
else (-> e { e.merge({ :y => (-> e { 20 })[e] }) })[e] 
end 
}", @iff.to_ruby
  end
  
  def test_eval
    env_1 = eval(@iff.to_ruby)[{}]
    assert_equal Hash[x: 10], env_1
  end

end
