require './fa_rule'
require 'test/unit'

class FARuleTest < Test::Unit::TestCase

    def test_create
        r = FARule.new(1, 'a', 2)
        assert_equal 1, r.state
        assert_equal 'a', r.character
        assert_equal 2, r.next_state
    end

    def test_applies_to
        r = FARule.new(1, 'a', 2)

        assert_true r.applies_to?(1, 'a')

        assert_false r.applies_to?(2, 'a')
        assert_false r.applies_to?(1, 'b')
        assert_false r.applies_to?(3, 'c')
    end

    def test_follow
        r = FARule.new(1, 'a', 2)
        assert_equal 2, r.follow
    end
end
