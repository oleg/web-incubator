require './number'
require './sequence'
require './assign'
require './less_than'

require './test_setup'

class SequenceTest < Test::Unit::TestCase

  def setup
    @seq = Sequence.new(Assign.new(:x, Number.new(10)),
                        Assign.new(:y, Number.new(20)))
  end

  def test_to_ruby
    assert_equal "-> e { (-> e { e.merge({ :y => (-> e { 20 })[e] }) })[(-> e { e.merge({ :x => (-> e { 10 })[e] }) })[e]] }", @seq.to_ruby
  end
  
  def test_eval
    env_1 = eval(@seq.to_ruby)[{}]
    assert_equal Hash[x: 10, y: 20], env_1
  end

end
