require './nfa'
require './fa_rule'
require './nfa_rulebook'

require 'set'

require 'test/unit'

class NFATest < Test::Unit::TestCase

  def setup
    @rulebook = NFARulebook.new([
      FARule.new(1, 'a', 1), FARule.new(1, 'b', 1), FARule.new(1, 'b', 2),
      FARule.new(2, 'a', 3), FARule.new(2, 'b', 3),
      FARule.new(3, 'a', 4), FARule.new(3, 'b', 4)
    ])
  end
  
  def test_accepting
    assert_false NFA.new(Set[1], [4], @rulebook).accepting?
    assert_true NFA.new(Set[1, 2, 4], [4], @rulebook).accepting?
  end

  def test_read_character
    nfa = NFA.new(Set[1], [4], @rulebook)
    assert_false nfa.accepting?

    nfa.read_character('b')
    assert_false nfa.accepting?

    nfa.read_character('a')
    assert_false nfa.accepting?

    nfa.read_character('b')
    assert_true nfa.accepting?
  end

  def test_read_string
    nfa = NFA.new(Set[1], [4], @rulebook)
    assert_false nfa.accepting?

    nfa.read_string('bbbbb')
    assert_true nfa.accepting?
  end
  
end
