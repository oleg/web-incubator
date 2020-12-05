package algoclass.node;

import org.junit.jupiter.api.Nested;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.Arguments;
import org.junit.jupiter.params.provider.MethodSource;

import java.util.stream.Stream;

import static algoclass.node.ListNodeAssertions.assertList;
import static algoclass.node.ListNodeAssertions.makeNodes;
import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertThrows;
import static org.junit.jupiter.params.provider.Arguments.arguments;

class ListNodeMergerTest {
    ListNodeMerger merger = new ListNodeMerger();

    @Nested
    class Args {
        @Test
        void first_argument_must_not_be_null() {
            IllegalArgumentException exception = assertThrows(IllegalArgumentException.class,
                    () -> {
                        ListNode sortedA = null;
                        ListNode sortedB = new ListNode(0, null, null);
                        merger.merge(sortedA, sortedB);
                    });
            assertEquals("first argument must not be null", exception.getMessage());
        }

        @Test
        void second_argument_must_not_be_null() {
            IllegalArgumentException exception = assertThrows(IllegalArgumentException.class,
                    () -> {
                        ListNode sortedA = new ListNode(0, null, null);
                        ListNode sortedB = null;
                        merger.merge(sortedA, sortedB);
                    });
            assertEquals("second argument must not be null", exception.getMessage());
        }

    }

    @MethodSource("data_2")
    @ParameterizedTest
    void merge_2(ListNode a, ListNode b) {
        ListNode merged = merger.merge(a, b);
        assertList(merged, 1, 2);
    }

    static Stream<Arguments> data_2() {
        return Stream.of(
                arguments(makeNodes(1), makeNodes(2)),
                arguments(makeNodes(2), makeNodes(1)));
    }

    @MethodSource("data_3")
    @ParameterizedTest
    void merge_3(ListNode a, ListNode b) {
        ListNode merged = merger.merge(a, b);

        assertList(merged, 1, 2, 3);
    }

    static Stream<Arguments> data_3() {
        return Stream.of(
                arguments(makeNodes(1, 2), makeNodes(3)),
                arguments(makeNodes(1, 3), makeNodes(2)),
                arguments(makeNodes(2, 3), makeNodes(1)),
                arguments(makeNodes(3), makeNodes(1, 2)),
                arguments(makeNodes(2), makeNodes(1, 3)),
                arguments(makeNodes(1), makeNodes(2, 3)));
    }

    @MethodSource("data_4")
    @ParameterizedTest
    void merge_4(ListNode a, ListNode b) {
        ListNode merged = merger.merge(a, b);

        assertList(merged, 1, 2, 3, 4);
    }

    static Stream<Arguments> data_4() {
        return Stream.of(
                arguments(makeNodes(1, 2, 3), makeNodes(4)),
                arguments(makeNodes(1, 2, 4), makeNodes(3)),
                arguments(makeNodes(1, 3, 4), makeNodes(2)),
                arguments(makeNodes(2, 3, 4), makeNodes(1)),
                arguments(makeNodes(4), makeNodes(1, 2, 3)),
                arguments(makeNodes(3), makeNodes(1, 2, 4)),
                arguments(makeNodes(2), makeNodes(1, 3, 4)),
                arguments(makeNodes(1), makeNodes(2, 3, 4)));
    }


}