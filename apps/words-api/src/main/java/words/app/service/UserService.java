package words.app.service;

import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;
import words.app.model.User;
import words.app.model.Word;
import words.app.model.WordsList;
import words.app.repository.UserRepository;
import words.app.repository.WordRepository;
import words.app.repository.WordsListRepository;
import words.web.json.Words;

import javax.inject.Inject;

@Slf4j
@Service
public class UserService {

    private final UserRepository userRepository;
    private final WordRepository wordRepository;
    private final WordsListRepository wordsListRepository;

    @Inject
    public UserService(UserRepository userRepository, WordRepository wordRepository, WordsListRepository wordsListRepository) {
        this.userRepository = userRepository;
        this.wordRepository = wordRepository;
        this.wordsListRepository = wordsListRepository;
    }

    public void addWord(String userId, String list, Word word) {
        User user = userRepository.findOne(userId);
        Word foundWord = findOrCreateWord(word);
        user.addWord(list, foundWord);
        userRepository.save(user);

        log.info("Added word '{}' to list '{}' of user '{}'", foundWord.getWord(), list, user.getName());
    }

    private Word findOrCreateWord(Word word) {
        Word found = wordRepository.findByWord(word.getWord());
        if (found == null) {
            found = wordRepository.save(word);
        }
        return found;
    }

    public void addList(User user, WordsList wordsList) {
        if (user.getListsNames().contains(wordsList.getName())) {
            throw new AlreadyExistException();
        }

        WordsList save = wordsListRepository.save(wordsList);
        user.addWordsList(save);
        userRepository.save(user);
    }

    @Transactional
    public void removeList(String userName, String list) {
        User user = findUser(userName);
        user.removeList(list);
    }

    public Words readWordsFromList(User user, String list) {
        WordsList wordsList = user.getList(list);
        Words words = new Words(wordsList.getWordsAsSortedString());
        log.info("Found words '{}' for list '{}' and user '{}'", words.getWords(), list, user.getName());
        return words;
    }


    private User findUser(String userName) {
        User user = userRepository.findOne(userName);
        if (user == null) {
            throw new NotFoundException();
        }
        return user;
    }

}
