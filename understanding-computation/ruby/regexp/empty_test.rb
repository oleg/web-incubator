require './empty'

require './test_setup'

class EmptyTest < Test::Unit::TestCase
  
  def test_to_s
    assert_equal '', Empty.new.to_s
  end

  def test_inspect
    assert_equal '//', Empty.new.inspect
  end

  def test_nfa_design
    d = Empty.new.to_nfa_design
    assert_true d.accepts?('')
    assert_false d.accepts?('a')
  end

  def test_matches
    assert_true Empty.new.matches?('')
    assert_false Empty.new.matches?('a')
  end

end
