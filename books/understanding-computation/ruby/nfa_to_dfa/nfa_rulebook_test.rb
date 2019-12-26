require './nfa_rulebook'
require './fa_rule'

require 'set'

require 'test/unit'


class NFARulebookTest < Test::Unit::TestCase

  def setup
    @rulebook = NFARulebook.new([
      FARule.new(1, 'a', 1), FARule.new(1, 'b', 1), FARule.new(1, 'b', 2),
      FARule.new(2, 'a', 3), FARule.new(2, 'b', 3),
      FARule.new(3, 'a', 4), FARule.new(3, 'b', 4)
    ])
  end

  def test_next_states
    assert_equal Set[1, 2], @rulebook.next_states(Set[1], 'b')
    assert_equal Set[1, 3], @rulebook.next_states(Set[1, 2], 'a')
    assert_equal Set[1, 2, 4], @rulebook.next_states(Set[1, 3], 'b')
  end
  
end

class NFARulebookFreeMovesTest < Test::Unit::TestCase
  
  def setup
    @rulebook = NFARulebook.new([
      FARule.new(1, nil, 2), FARule.new(1, nil, 4),
      FARule.new(2, 'a', 3),
      FARule.new(3, 'a', 2)
    ])
  end
  
  def test_next_states_nil
    assert_equal Set[2, 4], @rulebook.next_states(Set[1], nil)
  end

  def test_follow_free_moves
    assert_equal Set[1, 2, 4], @rulebook.follow_free_moves(Set[1])
  end
end

class NFARulebookOperatorTest < Test::Unit::TestCase
  
  def test_plus
    a = NFARulebook.new([1, 2, 3])
    b = NFARulebook.new([4, 5, 6])
    assert_equal [1, 2, 3, 4, 5, 6], (a + b).rules
  end
  
  def test_plus_position
    a = NFARulebook.new([1, 2, 3])
    b = NFARulebook.new([4, 5, 6])
    assert_equal [4, 5, 6, 1, 2, 3], (b + a).rules
  end

  def test_create
    a = NFARulebook.new([1, 2, 3])
    b = NFARulebook.new([4, 5, 6])
    assert_not_equal a, (a + b)
    assert_not_equal b, (a + b)
  end

end
