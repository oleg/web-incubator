package lcode;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertFalse;
import static org.junit.jupiter.api.Assertions.assertTrue;

class ValidParenthesesTest {

    @Test
    void empty_string_is_valid() {
        assertTrue(new ValidParentheses().isValid(""));
    }

    @Test
    void one_paren_pair_is_valid() {
        assertTrue(new ValidParentheses().isValid("()"));
    }

    @Test
    void two_paren_pair_is_valid() {
        assertTrue(new ValidParentheses().isValid("()()"));
    }

    @Test
    void three_paren_pair_is_valid() {
        assertTrue(new ValidParentheses().isValid("()()()"));
    }

    @Test
    void nested_paren_pair_is_valid() {
        assertTrue(new ValidParentheses().isValid("(())"));
    }

    @Test
    void complex_valid() {
        assertTrue(new ValidParentheses().isValid("(())()()()((())(()()))"));
    }

    //false
    @Test
    void one_left_is_invalid() {
        assertFalse(new ValidParentheses().isValid("("));
    }

    @Test
    void two_left_is_invalid() {
        assertFalse(new ValidParentheses().isValid("(("));
    }

    @Test
    void one_pair_and_one_left_is_invalid() {
        assertFalse(new ValidParentheses().isValid("()("));
    }

    @Test
    void nested_in_pair_and_one_left_is_invalid() {
        assertFalse(new ValidParentheses().isValid("(()"));
    }

    @Test
    void one_right_is_invalid() {
        assertFalse(new ValidParentheses().isValid(")"));
    }

    //square
    @Test
    void one_square_paren_pair_is_valid() {
        assertTrue(new ValidParentheses().isValid("[]"));
    }

    @Test
    void mixed_paren_pair_is_invalid() {
        assertFalse(new ValidParentheses().isValid("(]"));
    }

    @Test
    void mixed_paren_pair_is_invalid_2() {
        assertFalse(new ValidParentheses().isValid("[)"));
    }

    @Test
    void complex_mixed_pairs_is_valid() {
        assertTrue(new ValidParentheses().isValid("[()()[]]([]()[])[]()[]"));
    }

    @Test
    void complex_extra_mixed_pairs_is_valid() {
        assertTrue(new ValidParentheses().isValid("[({}){}()[{{}}]]([]()[])[]()[][{{}{}}()]"));
    }

}