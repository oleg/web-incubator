require './boolean'
require './test_setup'

class NumberTest < Test::Unit::TestCase

  def test_number
    assert_equal "-> e { true }", Boolean.new(true).to_ruby
  end
  
end
