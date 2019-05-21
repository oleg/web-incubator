package words.app.service;

import org.junit.Before;
import org.junit.Rule;
import org.junit.Test;
import org.junit.rules.ExpectedException;
import org.junit.runner.RunWith;
import org.mockito.Mock;
import org.mockito.runners.MockitoJUnitRunner;
import words.app.model.User;
import words.app.model.WordsList;
import words.app.repository.UserRepository;
import words.app.repository.WordRepository;
import words.app.repository.WordsListRepository;

import static org.hamcrest.Matchers.is;
import static org.junit.Assert.assertThat;
import static org.mockito.Mockito.when;

@RunWith(MockitoJUnitRunner.class)
public class UserServiceTest {

    private static final String USER_NAME = "oleg";

    @Mock
    private UserRepository userRepository;
    @Mock
    private WordRepository wordRepository;
    @Mock
    private WordsListRepository wordsListRepository;

    private UserService userService;

    @Rule
    public ExpectedException exception = ExpectedException.none();

    @Before
    public void setUp() throws Exception {
        userService = new UserService(userRepository, wordRepository, wordsListRepository);
    }

    @Test
    public void removeList_noUserFound() throws Exception {
        exception.expect(NotFoundException.class);

        when(userRepository.findOne(USER_NAME)).thenReturn(null);

        userService.removeList(USER_NAME, "deutsch");
    }

    @Test
    public void removeList_noListFound() throws Exception {
        User user = new User();
        when(userRepository.findOne(USER_NAME))
                .thenReturn(user);

        userService.removeList(USER_NAME, "deutsch");

        assertThat(user.getListsNames().size(), is(0));
    }

    @Test
    public void removeList() throws Exception {
        User user = new User();
        user.addWordsList(buildWordsList("deutsch"));
        when(userRepository.findOne(USER_NAME))
                .thenReturn(user);

        userService.removeList(USER_NAME, "deutsch");

        assertThat(user.getListsNames().size(), is(0));
    }

    private WordsList buildWordsList(String english) {
        WordsList wordsList = new WordsList();
        wordsList.setName(english);
        return wordsList;
    }

}
