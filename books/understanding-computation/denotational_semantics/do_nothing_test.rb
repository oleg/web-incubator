require './number'
require './do_nothing'

require './test_setup'

class DoNothingTest < Test::Unit::TestCase
  
  def test_to_ruby
    assert_equal "-> e { e }", DoNothing.new.to_ruby
  end

  def test_eval
    env = {x: 1, y: 2, z: 3}
    env_1 = eval(DoNothing.new.to_ruby)[env]
    assert_equal Hash[x: 1, y: 2, z: 3], env_1
  end

end
