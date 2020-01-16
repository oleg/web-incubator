package words.web.controller;

import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;
import words.app.model.User;
import words.app.model.Word;
import words.app.model.WordsList;
import words.app.repository.UserRepository;
import words.app.service.UserService;
import words.web.json.Words;

import javax.inject.Inject;
import java.util.Optional;
import java.util.Set;

import static org.springframework.http.MediaType.APPLICATION_JSON_VALUE;

@RestController
@RequestMapping(value = "api/user/{user}/list", produces = APPLICATION_JSON_VALUE, consumes = APPLICATION_JSON_VALUE)
public class WordsListController {

    private final UserRepository userRepository;
    private final UserService userService;

    @Inject
    public WordsListController(UserRepository userRepository, UserService userService) {
        this.userRepository = userRepository;
        this.userService = userService;
    }

    @RequestMapping(method = RequestMethod.GET)
    public ResponseEntity<Set<String>> readLists(@PathVariable String user) {
        User byId = userRepository.findById(user).orElseThrow();
        Set<String> listsNames = byId.getListsNames();
        return ResponseEntity.ok(listsNames);
    }

    @RequestMapping(method = RequestMethod.POST)
    public ResponseEntity<String> addList(@PathVariable String user,
                                          @RequestBody WordsList wordsList) {
        User u = userRepository.findById(user).orElse(null);
        userService.addList(u, wordsList);
        return ResponseEntity.ok("{}");
    }

    @RequestMapping(method = RequestMethod.DELETE, value = "{list}")
    public ResponseEntity<Void> removeList(@PathVariable String user,
                                           @PathVariable String list) {

        userService.removeList(user, list);
        return ResponseEntity.noContent().build();
    }

    @RequestMapping(method = RequestMethod.GET, value = "{list}", produces = APPLICATION_JSON_VALUE)
    public ResponseEntity<Words> readOneWordsList(@PathVariable String user, @PathVariable String list) {
        User u = userRepository.findById(user).orElse(null);
        Words body = userService.readWordsFromList(u, list);
        return ResponseEntity.ok(body);
    }

    @RequestMapping(method = RequestMethod.POST, value = "{list}")
    public ResponseEntity<String> addWord(@PathVariable String user, @PathVariable String list, @RequestBody Word word) {
        userService.addWord(user, list, word);
        return ResponseEntity.ok("{}");
    }
}
