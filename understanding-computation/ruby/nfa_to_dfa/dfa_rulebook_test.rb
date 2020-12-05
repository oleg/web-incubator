require './fa_rule'
require './dfa_rulebook'

require 'test/unit'

class DFARulebookTest < Test::Unit::TestCase

    def setup
        @rulebook = DFARulebook.new([
            FARule.new(1, 'a', 2), FARule.new(1, 'b', 1),
            FARule.new(2, 'a', 2), FARule.new(2, 'b', 3),
            FARule.new(3, 'a', 3), FARule.new(3, 'b', 3)
        ])

    end

    def test_next_state
        assert_equal 2, @rulebook.next_state(1, 'a')
        assert_equal 1, @rulebook.next_state(1, 'b')
        assert_equal 3, @rulebook.next_state(2, 'b')
    end
end
