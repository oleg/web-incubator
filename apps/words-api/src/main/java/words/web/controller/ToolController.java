package words.web.controller;

import lombok.extern.slf4j.Slf4j;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;
import words.app.model.User;
import words.app.repository.UserRepository;
import words.app.service.WordsService;
import words.web.json.WordCount;
import words.web.json.WordsRequest;

import javax.inject.Inject;
import java.util.List;

import static org.springframework.http.MediaType.APPLICATION_JSON_VALUE;

@Slf4j
@RestController
@SessionAttributes("username")
@RequestMapping(value = "api/tool/to-words", produces = APPLICATION_JSON_VALUE)
public class ToolController {

    private final WordsService wordsService;
    private final UserRepository userRepository;

    @Inject
    public ToolController(WordsService wordsService, UserRepository userRepository) {
        this.wordsService = wordsService;
        this.userRepository = userRepository;
    }

    @RequestMapping(method = RequestMethod.POST, consumes = MediaType.APPLICATION_JSON_VALUE)
    public ResponseEntity<List<WordCount>> textToWords(@RequestBody WordsRequest wordsRequest,
                                                       @ModelAttribute("username") String username) {

        log.info("text to words for user {} ignoring {}", username, wordsRequest.getIgnoreWordLists());
        User myUser = userRepository.findOne(username);

        return ResponseEntity.ok(wordsService.toWords(wordsRequest, myUser));
    }
}
