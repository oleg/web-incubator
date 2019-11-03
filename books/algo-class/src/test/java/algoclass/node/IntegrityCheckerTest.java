package algoclass.node;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertThrows;

class IntegrityCheckerTest {

    IntegrityChecker checker = new IntegrityChecker();

    @Test
    void first_must_not_have_prev() {
        IllegalStateException exception = assertThrows(IllegalStateException.class,
                () -> {
                    ListNode prev = new ListNode(900, null, null);
                    ListNode first = new ListNode(100, prev, null);
                    checker.onChange(first, null);
                });

        assertEquals("first prev is not null", exception.getMessage());
    }

}