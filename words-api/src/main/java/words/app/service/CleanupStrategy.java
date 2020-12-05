package words.app.service;

import java.util.Arrays;
import java.util.List;
import java.util.stream.Collectors;
import java.util.stream.Stream;

public class CleanupStrategy {

    public Stream<String> toWords(String text) {
        return Arrays.stream(text.split("[ .,\n]"))
            .map(t -> t.replaceAll("\\P{L}", ""))
            .map(String::toLowerCase)
            .filter(w -> !w.isEmpty());
    }

    public List<String> toWordsList(String text) {
        return toWords(text).collect(Collectors.toList());
    }

}
