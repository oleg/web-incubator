package words.web.json;

import java.util.ArrayList;
import java.util.List;

public class Words {

    private final List<String> words;

    public Words() {
        words = new ArrayList<>();
    }

    public Words(List<String> words) {this.words = words;}

    public List<String> getWords() {
        return words;
    }
}
