require './number'
require './assign'

require './test_setup'

class AssignTest < Test::Unit::TestCase
  
  def test_to_ruby
    assign = Assign.new(:x, Number.new(100))
    
    assert_equal "-> e { e.merge({ :x => (-> e { 100 })[e] }) }", assign.to_ruby
  end

  def test_eval
    assign = Assign.new(:x, Number.new(100))
    result = eval(assign.to_ruby)[{}]
    assert_equal Hash[x: 100], result
  end

end
