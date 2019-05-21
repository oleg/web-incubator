package words.app.service;

import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Service;
import words.app.model.User;
import words.web.json.WordCount;
import words.web.json.WordsRequest;

import java.util.Collections;
import java.util.Comparator;
import java.util.List;
import java.util.Map;

import static java.util.function.Function.identity;
import static java.util.stream.Collectors.*;

@Slf4j
@Service
public class WordsService {

    private final CleanupStrategy cleanupStrategy = new CleanupStrategy();

    public <T> List<WordCount> toWords(WordsRequest wordsRequest, User myUser) {
        if (wordsRequest == null) {
            return Collections.emptyList();
        }

        Map<String, Long> countWords = countWords(wordsRequest);

        List<String> ignoreWords = getIgnoreWords(wordsRequest, myUser);

        countWords.keySet().removeAll(ignoreWords);

        return toWordCountList(countWords);
    }

    private List<String> getIgnoreWords(WordsRequest wordsRequest, User myUser) {
        return myUser.getLists(wordsRequest.getIgnoreWordLists())
                .stream()
                .flatMap(wl -> wl.getWordsAsSortedString().stream())
                .collect(toList());
    }

    private List<WordCount> toWordCountList(Map<String, Long> collect) {
        return collect.entrySet().stream()
                .sorted(
                        Comparator.comparing(Map.Entry<String, Long>::getValue).reversed()
                                .thenComparing(Map.Entry::getKey))
                .map(e -> new WordCount(e.getKey(), e.getValue()))
                .collect(toList());
    }

    private Map<String, Long> countWords(WordsRequest wordsRequest) {
        return cleanupStrategy.toWords(wordsRequest.getText())
                .collect(groupingBy(identity(), counting()));
    }

}
