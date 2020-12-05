require './dfa_design'
require './dfa'
require './fa_rule'
require './dfa_rulebook'

require 'test/unit'


class DFADesignTest < Test::Unit::TestCase
  
  def setup
    @rulebook = DFARulebook.new([
      FARule.new(1, 'a', 2), FARule.new(1, 'b', 1),
      FARule.new(2, 'a', 2), FARule.new(2, 'b', 3),
      FARule.new(3, 'a', 3), FARule.new(3, 'b', 3)
    ])
  end

  def test_to_dfa
    assert_equal DFA.new(1, [3], @rulebook),
      DFADesign.new(1, [3], @rulebook).to_dfa
    
    assert_equal DFA.new(3, [4, 5, 6], @rulebook),
      DFADesign.new(3, [4, 5, 6], @rulebook).to_dfa
  end

  def test_accepts
    dfa_design = DFADesign.new(1, [3], @rulebook)
    
    assert_false dfa_design.accepts?('a')
    assert_false dfa_design.accepts?('baa')
    assert_true dfa_design.accepts?('baba')
  end
  
end
