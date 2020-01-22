package jat;

import org.junit.jupiter.api.Test;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
import java.util.Currency;
import java.util.List;
import java.util.function.IntFunction;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.mockito.Mockito.mock;
import static org.mockito.Mockito.verify;

class CollectionsTest {

    @Test
    void sort_delegates_to_list_sort() {
        @SuppressWarnings("unchecked")
        List<Integer> list = mock(List.class);
        Collections.sort(list);

        verify(list).sort(null);
    }

    @Test
    void replaceAll_replaces_all_objects() {
        List<String> strings = new ArrayList<>(List.of("a", "b", "c"));

        strings.replaceAll(s -> s + "!");

        assertEquals(List.of("a!", "b!", "c!"), strings);
    }

    @Test
    void name() {
//        Set<String> objects = Collections.newSetFromMap(new ConcurrentHashMap<>());
//        objects.add("a");
//
//        ArrayList<String> objects1 = new ArrayList<>(List.of("a", "b", "c", "d", "e", "f"));
//        Collections.rotate(objects1, objects1.size()-1);
//        System.out.println(objects1);
        String[] arr = {"a", "b", "c"};
        Arrays.parallelPrefix(arr, (a, b) -> a + b);
        System.out.println(Arrays.toString(arr));

        String[] arr1 = {"a", "b", "c"};
        Arrays.setAll(arr1, i -> i + "!");
        System.out.println(Arrays.toString(arr1));
    }

    @Test
    void nam1e() {
        Currency eur = Currency.getInstance("EUR");
        System.out.println(eur);
    }
}
