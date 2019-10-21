package jat;

import org.junit.jupiter.api.Test;

import java.util.Collections;
import java.util.List;

import static org.mockito.Mockito.mock;
import static org.mockito.Mockito.verify;

class CollectionsTest {

    @Test
    void sort_delegates_to_list_sort() {
        List<Integer> list = mock(List.class);
        Collections.sort(list);

        verify(list).sort(null);
    }
}
