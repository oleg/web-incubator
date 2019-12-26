require './dfa'
require './dfa_rulebook'
require './fa_rule'


require "test/unit"

class DFATest < Test::Unit::TestCase
  
    def setup
        @rulebook = DFARulebook.new([
            FARule.new(1, 'a', 2), FARule.new(1, 'b', 1),
            FARule.new(2, 'a', 2), FARule.new(2, 'b', 3),
            FARule.new(3, 'a', 3), FARule.new(3, 'b', 3)
        ])
    end

    def test_create
        dfa = DFA.new(1, [1, 3], @rulebook)
        assert_equal 1, dfa.current_state
        assert_equal [1, 3], dfa.accept_states
        assert_equal @rulebook, dfa.rulebook
    end

    def test_accepting
        assert_true DFA.new(1, [1, 3], @rulebook).accepting?
        assert_false DFA.new(1, [3], @rulebook).accepting?
    end

    def test_read_character
        dfa = DFA.new(1, [3], @rulebook)
        assert_false dfa.accepting?

        dfa.read_character('b')
        assert_false dfa.accepting?

        3.times { dfa.read_character('a') }
        assert_false dfa.accepting?

        dfa.read_character('b')
        assert_true dfa.accepting?
    end

    def test_read_string
      dfa = DFA.new(1, [3], @rulebook)
      assert_false dfa.accepting?

      dfa.read_string('baaab')
      assert_true dfa.accepting?
    end
    

end
