require './number'
require './do_nothing'

require './test_setup'

class DoNothingTest < Test::Unit::TestCase
  
  def test_empty
    assert_equal Hash[], DoNothing.new.evaluate({})
  end

  def test_non_empty
    assert_equal Hash[x: Number.new(7)], DoNothing.new.evaluate({x: Number.new(7)})
  end
end
