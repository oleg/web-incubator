require './test_setup'

require './nfa_simulation'

require './nfa_rulebook'
require './nfa_design'
require './fa_rule'


class NFASimulationTest < Test::Unit::TestCase

  def setup
    rulebook = NFARulebook.new([
      FARule.new(1, 'a', 1), FARule.new(1, 'a', 2), FARule.new(1, nil, 2),
      FARule.new(2, 'b', 3),
      FARule.new(3, 'b', 1), FARule.new(3, nil, 2)
    ])
    @nfa_design = NFADesign.new(1, [3], rulebook)
    @simulation = NFASimulation.new(@nfa_design)
  end
  
  def test_create
    assert_equal Set[1, 2], @simulation.next_state(Set[1, 2], 'a')
    assert_equal Set[2, 3], @simulation.next_state(Set[1, 2], 'b')
    assert_equal Set[1, 2, 3], @simulation.next_state(Set[3, 2], 'b')
    assert_equal Set[1, 2, 3], @simulation.next_state(Set[1, 2, 3], 'b')
    assert_equal Set[1, 2], @simulation.next_state(Set[1, 2, 3], 'a')
  end

  def test_rules_for
    assert_equal [
      FARule.new(Set[1, 2], 'a', Set[1, 2]),
      FARule.new(Set[1, 2], 'b', Set[2, 3])
    ],
      @simulation.rules_for(Set[1, 2])
  end

  def test_rules_for_v2
    assert_equal [
      FARule.new(Set[2, 3], 'a', Set[]),
      FARule.new(Set[2, 3], 'b', Set[1, 2, 3])
    ],
      @simulation.rules_for(Set[2, 3])
  end

  def test_discover_states_and_rules
    start_state = @nfa_design.to_nfa.current_states
    states, rules = @simulation.discover_states_and_rules(Set[start_state])
    assert_equal Set[
      Set[1,2],
      Set[2,3],
      Set[],
      Set[1,2,3]
    ], states
    
    assert_equal [
      FARule.new(Set[1,2], 'a', Set[1,2]),
      FARule.new(Set[1,2], 'b', Set[2,3]),
      FARule.new(Set[2,3], 'a', Set[]),
      FARule.new(Set[2,3], 'b', Set[1,2,3]),
      FARule.new(Set[], 'a', Set[]),
      FARule.new(Set[], 'b', Set[]),
      FARule.new(Set[1,2,3], 'a', Set[1,2]),
      FARule.new(Set[1,2,3], 'b', Set[1,2,3])
    ], rules
  end

  def test_accepting
    assert_true @nfa_design.to_nfa(Set[2,3]).accepting?
    assert_false @nfa_design.to_nfa(Set[1,2]).accepting?
  end

  def test_to_dfa_design
    dfa_design = @simulation.to_dfa_design

    assert_false dfa_design.accepts?('aaa')
    assert_true dfa_design.accepts?('aab')
    assert_true dfa_design.accepts?('bbbabb')
  end
  
end
