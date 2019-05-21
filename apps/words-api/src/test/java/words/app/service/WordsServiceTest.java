package words.app.service;

import org.junit.Test;
import words.app.model.User;
import words.app.model.Word;
import words.app.model.WordsList;
import words.web.json.WordCount;
import words.web.json.WordsRequest;

import java.util.Collections;
import java.util.List;

import static org.hamcrest.CoreMatchers.hasItems;
import static org.hamcrest.CoreMatchers.is;
import static org.junit.Assert.assertThat;

public class WordsServiceTest {

    private static final User user = new User("oleg", "123");

    private WordsService wordsService = new WordsService();

    @Test
    public void test_empty() throws Exception {

        assertThat(wordsService.toWords(new WordsRequest("", Collections.emptyList()), user).size(), is(0));
        assertThat(wordsService.toWords(null, user).size(), is(0));
    }

    @Test
    public void should_be_split_by_whitespaces() throws Exception {
        List<WordCount> wordCounts = wordsService.toWords(new WordsRequest("The simple sentence.", Collections.emptyList()), user);
        assertThat(wordCounts.size(), is(3));
    }

    @Test
    public void should_transform_to_lower_case() throws Exception {
        List<WordCount> wordCounts = wordsService.toWords(new WordsRequest("The simple sentence.", Collections.emptyList()), user);
        assertThat(wordCounts.get(2).getWord(), is("the"));
    }

    @Test
    public void should_sort_alphabetically_inside_same_usage() throws Exception {
        List<WordCount> wordCounts = wordsService.toWords(new WordsRequest("c b a", Collections.emptyList()), user);
        assertThat(wordCounts.get(0).getWord(), is("a"));
        assertThat(wordCounts.get(1).getWord(), is("b"));
        assertThat(wordCounts.get(2).getWord(), is("c"));
    }

    @Test
    public void should_report_usage_number() throws Exception {
        List<WordCount> wordCounts = wordsService.toWords(new WordsRequest("c b a", Collections.emptyList()), user);
        assertThat(wordCounts.get(0).getCount(), is(1L));
        assertThat(wordCounts.get(1).getCount(), is(1L));
        assertThat(wordCounts.get(2).getCount(), is(1L));
    }

    @Test
    public void should_sort_by_usage_and_report_usage_number_for_repeated_words() throws Exception {
        List<WordCount> wordCounts = wordsService.toWords(new WordsRequest("c c c b b a", Collections.emptyList()), user);
        assertThat(wordCounts.get(0), is(new WordCount("c", 3L)));
        assertThat(wordCounts.get(1), is(new WordCount("b", 2L)));
        assertThat(wordCounts.get(2), is(new WordCount("a", 1L)));
    }

    @Test
    public void test_withIgnoreList() throws Exception {
        WordsList wordsList = new WordsList();
        wordsList.setName("study");
        wordsList.addWord(new Word("a"));
        wordsList.addWord(new Word("e"));
        user.addWordsList(wordsList);

        WordsRequest wordsRequest = new WordsRequest("a b c d e", Collections.singletonList("study"));
        List<WordCount> wordCounts = wordsService.toWords(wordsRequest, user);

        assertThat(wordCounts.size(), is(3));
        assertThat(wordCounts, hasItems(
                new WordCount("b", 1),
                new WordCount("c", 1),
                new WordCount("d", 1)
        ));
    }
}
