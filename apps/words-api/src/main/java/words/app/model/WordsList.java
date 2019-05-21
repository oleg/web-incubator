package words.app.model;

import com.google.common.base.MoreObjects;

import javax.persistence.*;
import javax.validation.constraints.NotNull;
import java.util.HashSet;
import java.util.List;
import java.util.Set;

import static java.util.stream.Collectors.toList;

@Entity
public class WordsList {

    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    private Long id;

    @NotNull
    private String name;

    public WordsList(){}

    public WordsList(String name) {
        this.name = name;
    }

    @NotNull
    @ManyToMany
    private Set<Word> words = new HashSet<>();

    public List<String> getWordsAsSortedString() {
        return words.stream()
                .map(Word::getWord)
                .sorted()
                .collect(toList());
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public void addWord(Word word) {
        words.add(word);
    }

    @Override
    public String toString() {
        return MoreObjects.toStringHelper(this)
            .add("words", words)
            .toString();
    }

}
