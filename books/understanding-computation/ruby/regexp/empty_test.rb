require './empty'

require './test_setup'

class EmptyTest < Test::Unit::TestCase
  
  def test_to_s
    assert_equal '', Empty.new.to_s
  end

  def test_inspect
    assert_equal '//', Empty.new.inspect
  end

end
