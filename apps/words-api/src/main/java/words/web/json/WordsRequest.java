package words.web.json;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.util.List;

/**
 * Created by evgeniia.prozorova
 * on 4/25/16.
 */
@Data
@AllArgsConstructor
@NoArgsConstructor
public class WordsRequest {

    private String text;
    private List<String> ignoreWordLists;

}
