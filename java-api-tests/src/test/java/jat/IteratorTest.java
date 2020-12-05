package jat;

import org.junit.jupiter.api.DisplayNameGeneration;
import org.junit.jupiter.api.DisplayNameGenerator;
import org.junit.jupiter.api.Test;

import java.util.Iterator;
import java.util.NoSuchElementException;

import static org.junit.jupiter.api.Assertions.*;

@DisplayNameGeneration(DisplayNameGenerator.ReplaceUnderscores.class)
class IteratorTest {

    @Test
    void can_be_used_in_loop() {
        Iterator<String> iterator = new StringNumberIterator(3);

        StringBuilder result = new StringBuilder();

        while (iterator.hasNext())
            result.append(iterator.next());

        assertEquals("321", result.toString());
    }

    @Test
    void has_next_return_true_if_there_is_next_element() {
        Iterator<String> iterator = new StringNumberIterator(3);

        assertTrue(iterator.hasNext());

        iterator.next();
        assertTrue(iterator.hasNext());

        iterator.next();
        assertTrue(iterator.hasNext());

        iterator.next();
        assertFalse(iterator.hasNext());
    }

    @Test
    void has_next_return_same_value_until_next_is_invoked() {
        Iterator<String> iterator = new StringNumberIterator(3);

        assertTrue(iterator.hasNext());
        assertTrue(iterator.hasNext());
        assertTrue(iterator.hasNext());
        assertTrue(iterator.hasNext());

        iterator.next();
        iterator.next();
        iterator.next();

        assertFalse(iterator.hasNext());
        assertFalse(iterator.hasNext());
        assertFalse(iterator.hasNext());
        assertFalse(iterator.hasNext());
    }

    @Test
    void next_return_value_if_there_is_next_element() {
        Iterator<String> iterator = new StringNumberIterator(3);

        assertEquals("3", iterator.next());
        assertEquals("2", iterator.next());
        assertEquals("1", iterator.next());
    }

    @Test
    void next_throws_exception_if_there_are_no_more_elements() {
        Iterator<String> iterator = new StringNumberIterator(3);

        iterator.next();
        iterator.next();
        iterator.next();
        assertThrows(NoSuchElementException.class, iterator::next);
        assertThrows(NoSuchElementException.class, iterator::next);
    }

    @Test
    void forEachRemaining_has_default_implementation() {
        Iterator<String> iterator = new StringNumberIterator(3);

        StringBuilder result = new StringBuilder();
        iterator.forEachRemaining(result::append);

        assertEquals("321", result.toString());
    }

    @Test
    void forEachRemaining_iterates_over_remaining_elements() {
        Iterator<String> iterator = new StringNumberIterator(3);

        iterator.next();

        StringBuilder result = new StringBuilder();
        iterator.forEachRemaining(result::append);

        assertEquals("21", result.toString());
    }

    @Test
    void forEachRemaining_does_nothing_if_iterator_is_empty() {
        Iterator<String> iterator = new StringNumberIterator(3);

        iterator.next();
        iterator.next();
        iterator.next();

        StringBuilder result = new StringBuilder();
        iterator.forEachRemaining(result::append);

        assertEquals("", result.toString());
    }

    @Test
    void remove_throws_exception_by_default() {
        Iterator<String> iterator = new StringNumberIterator(3);

        iterator.next();
        iterator.next();
        iterator.next();

        assertThrows(UnsupportedOperationException.class, iterator::remove);
    }

    static class StringNumberIterator implements Iterator<String> {
        private int i;

        StringNumberIterator(int i) {
            this.i = i;
        }

        @Override
        public boolean hasNext() {
            return i > 0;
        }

        @Override
        public String next() {
            if (hasNext()) {
                return String.valueOf(i--);
            }
            throw new NoSuchElementException();
        }
    }
}
