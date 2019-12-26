require './literal'
require './test_setup'

class LiteralTest < Test::Unit::TestCase
  
  def test_to_s
    assert_equal 'a', Literal.new('a').to_s
    assert_equal 'b', Literal.new('b').to_s    
  end

  def test_inspect
    assert_equal '/c/', Literal.new('c').inspect
    assert_equal '/d/', Literal.new('d').inspect
  end

  def test_to_nfa_design
    d = Literal.new('a').to_nfa_design
    
    assert_false d.accepts?('')
    assert_true d.accepts?('a')
    assert_false d.accepts?('b')
  end

  def test_matches
    assert_false Literal.new('a').matches?('')
    assert_true Literal.new('b').matches?('b')
    assert_false Literal.new('b').matches?('a')    
  end
  
end
