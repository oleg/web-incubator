require './number'
require './test_setup'

class NumberTest < Test::Unit::TestCase

  def test_number
    assert_equal "-> e { 10 }", Number.new(10).to_ruby
  end
  
end
