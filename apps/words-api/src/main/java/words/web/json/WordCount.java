package words.web.json;

import java.util.Objects;

public class WordCount {

    private String word;
    private long count;

    public WordCount(String word, long count) {
        this.word = word;
        this.count = count;
    }

    public String getWord() {
        return word;
    }

    public long getCount() {
        return count;
    }

    @Override
    public String toString() {
        return "{" + word + "(" + count + ")}";
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) {
            return true;
        }
        if (o == null || getClass() != o.getClass()) {
            return false;
        }
        WordCount wordCount = (WordCount) o;
        return count == wordCount.count &&
            Objects.equals(word, wordCount.word);
    }

    @Override
    public int hashCode() {
        return Objects.hash(word, count);
    }
}
