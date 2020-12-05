package words.app.model;

import org.springframework.util.CollectionUtils;

import javax.persistence.CascadeType;
import javax.persistence.Entity;
import javax.persistence.Id;
import javax.persistence.OneToMany;
import java.util.*;
import java.util.stream.Collectors;

import static java.util.Arrays.asList;
import static java.util.stream.Collectors.toList;

@Entity(name = "wuser")
public class User {

    @Id
    private String name;
    private String password;
    @OneToMany(cascade = CascadeType.ALL)
    private Set<WordsList> lists = new HashSet<>();

    public User() {
        //for hibernate
    }

    public User(String name, String password) {
        this.name = name;
        this.password = password;
    }

    public WordsList getList(String list) {
        return lists.stream()
                .filter(l -> l.getName().equals(list))
                .findFirst()
                .orElseThrow(() -> new IllegalArgumentException("Not found list: " + list));
    }

    public Collection<WordsList> getLists(List<String> names) {
        if (CollectionUtils.isEmpty(names)) {
            return Collections.emptyList();
        }
        return lists.stream()
                .filter(wordsList -> names.contains(wordsList.getName()))
                .collect(toList());
    }

    public void removeList(String name) {
        lists.removeIf(wordsList -> wordsList.getName().equals(name));
    }

    public String getName() {
        return name;
    }

    public String getPassword() {
        return password;
    }

    public void setName(String name) {
        this.name = name;
    }

    public void addWordsList(WordsList wordsList) {
        lists.add(wordsList);
    }

    public void addWordsLists(WordsList... wordsList) {
        lists.addAll(asList(wordsList));
    }

    public void addWord(String list, Word word) {
        getList(list).addWord(word);
    }

    public Set<String> getListsNames() {
        return lists.stream().map(WordsList::getName).collect(Collectors.toSet());
    }

}
