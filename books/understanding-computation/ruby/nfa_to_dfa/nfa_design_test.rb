require './nfa_design'
require './nfa_rulebook'
require './fa_rule'

require 'test/unit'

class NFADesignTest < Test::Unit::TestCase

  def setup
    @rulebook = NFARulebook.new([
      FARule.new(1, 'a', 1), FARule.new(1, 'b', 1), FARule.new(1, 'b', 2),
      FARule.new(2, 'a', 3), FARule.new(2, 'b', 3),
      FARule.new(3, 'a', 4), FARule.new(3, 'b', 4)
    ])
  end
  
  def test_accepts
    nfa_design = NFADesign.new(1, [4], @rulebook)
    
    assert_true nfa_design.accepts?('bab')
    assert_true nfa_design.accepts?('bbbbb')
    assert_false nfa_design.accepts?('bbabb')
  end
  
end

class NFADesignFreeMovesTest < Test::Unit::TestCase
  
  def setup
    @rulebook = NFARulebook.new([
     FARule.new(1, nil, 2), FARule.new(1, nil, 4),
     FARule.new(2, 'a', 3),
     FARule.new(3, 'a', 2),
     FARule.new(4, 'a', 5),
     FARule.new(5, 'a', 6),
     FARule.new(6, 'a', 4)
   ])
  end

  def test_free_moves
    nfa_design = NFADesign.new(1, [2, 4], @rulebook)
    
    assert_true nfa_design.accepts?('aa')
    assert_true nfa_design.accepts?('aaa')
    assert_false nfa_design.accepts?('aaaaa')
    assert_true nfa_design.accepts?('aaaaaa')
  end
  
end

