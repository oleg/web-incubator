package words.web.controller;

import org.junit.Test;
import org.junit.runner.RunWith;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.runners.MockitoJUnitRunner;
import words.app.model.User;
import words.app.repository.UserRepository;
import words.app.service.WordsService;
import words.web.json.WordsRequest;

import java.util.Collections;

import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.when;

@RunWith(MockitoJUnitRunner.class)
public class ToolControllerTest {

    @Mock WordsService wordsService;
    @Mock UserRepository userRepository;

    @InjectMocks ToolController toolController;

    private User user = new User("oleg", "123");

    @Test
    public void should_exclude_known_words() throws Exception {
        when(userRepository.findOne("oleg")).thenReturn(user);

        toolController.textToWords(new WordsRequest("hello hello a world", Collections.emptyList()), "oleg");

        verify(wordsService).toWords(new WordsRequest("hello hello a world", Collections.emptyList()), user);
    }

}
