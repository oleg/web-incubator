require './boolean'

require './test_setup'


class BooleanTest < Test::Unit::TestCase
  
  def test_new_true
    assert_equal true, Boolean.new(true).value
  end
  
  def test_new_false
    assert_equal false, Boolean.new(false).value
  end

  def test_to_s
    assert_equal "true", Boolean.new(true).to_s
    assert_equal "false", Boolean.new(false).to_s
  end

  def test_inspect
    assert_equal "«true»", Boolean.new(true).inspect
    assert_equal "«false»", Boolean.new(false).inspect    
  end
  
  def test_number
    assert_false Boolean.new(true).reducible?
    assert_false Boolean.new(false).reducible?
  end
  
end
