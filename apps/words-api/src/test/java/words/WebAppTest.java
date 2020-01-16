package words;

import com.fasterxml.jackson.databind.ObjectMapper;
import org.junit.Before;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.boot.test.web.client.TestRestTemplate;
import org.springframework.http.*;
import org.springframework.test.annotation.DirtiesContext;
import org.springframework.test.annotation.Rollback;
import org.springframework.test.context.ActiveProfiles;
import org.springframework.test.context.junit4.SpringJUnit4ClassRunner;
import words.app.model.User;
import words.app.model.Word;
import words.app.model.WordsList;
import words.app.repository.UserRepository;
import words.app.repository.WordRepository;
import words.app.repository.WordsListRepository;
import words.web.json.Words;

import javax.inject.Inject;
import java.util.Base64;

import static java.util.Arrays.asList;
import static org.hamcrest.CoreMatchers.*;
import static org.hamcrest.beans.HasPropertyWithValue.hasProperty;
import static org.hamcrest.core.IsNot.not;
import static org.junit.Assert.assertNotNull;
import static org.junit.Assert.assertThat;

@RunWith(SpringJUnit4ClassRunner.class)
@SpringBootTest(classes = WebApp.class, webEnvironment = SpringBootTest.WebEnvironment.DEFINED_PORT)
@Rollback
@ActiveProfiles("test")
@DirtiesContext(classMode = DirtiesContext.ClassMode.AFTER_EACH_TEST_METHOD)
public class WebAppTest {

    private String root = "http://localhost:9099/api/";
    private TestRestTemplate restTemplate = new TestRestTemplate();
    private ObjectMapper mapper = new ObjectMapper();
    @Inject
    private UserRepository userRepository;
    @Inject
    private WordsListRepository wordsListRepository;
    @Inject
    private WordRepository wordRepository;

    @Before
    public void setUp() throws Exception {
        WordsList wordsList1 = new WordsList("study");
        wordsListRepository.save(wordsList1);

        User user1 = new User("oleg", "123");
        user1.addWordsList(wordsList1);
        userRepository.save(user1);

        WordsList wordsList2 = new WordsList("study");
        wordsListRepository.save(wordsList2);

        User user2 = new User("zheka", "456");
        user2.addWordsList(wordsList2);
        userRepository.save(user2);
    }

    @Test
    public void whenAddListWithSameName_throw400() throws Exception {
        WordsList entity = new WordsList("study");

        ResponseEntity<String> addWord = post(url("user/oleg/list"), entity, "oleg", "123");

        assertThat(addWord.getStatusCode().value(), is(400));
    }

    @Test
    public void add_new_word_one_user_one_list_as_json() throws Exception {
        Word entity = new Word();
        entity.setWord("me");

        ResponseEntity<String> addWord = post(url("user/oleg/list/study"), entity, "oleg", "123");
        assertThat(addWord.getStatusCode().value(), is(200));
        assertThat(addWord.getBody(), is("{}"));

        ResponseEntity<Words> studyList = get(url("user/oleg/list/study"), Words.class, "oleg", "123");
        assertThat(studyList.getStatusCode().value(), is(200));

        Words actual = studyList.getBody();
        assertThat(actual.getWords(), is(asList("me")));
    }

    @Test
    public void word_added_by_one_user_should_not_be_visible_to_other() throws Exception {
        Word entity = new Word();
        entity.setWord("me");

        ResponseEntity<String> addWord = post(url("user/oleg/list/study"), entity, "oleg", "123");
        assertThat(addWord.getStatusCode().value(), is(200));
        assertThat(addWord.getBody(), is("{}"));

        ResponseEntity<String> studyList = get(url("user/zheka/list/study"), String.class, "zheka", "456");
        assertThat(studyList.getStatusCode().value(), is(200));
        assertThat(studyList.getBody(), is("{\"words\":[]}"));
    }

    @Test
    public void two_users_could_add_same_word() throws Exception {
        Word entity = new Word();
        entity.setWord("me");

        ResponseEntity<String> addWord = post(url("user/oleg/list/study"), entity, "oleg", "123");
        assertThat(addWord.getStatusCode().value(), is(200));
        assertThat(addWord.getBody(), is("{}"));

        ResponseEntity<String> addWord2 = post(url("user/zheka/list/study"), entity, "zheka", "456");
        assertThat(addWord2.getStatusCode().value(), is(200));
        assertThat(addWord2.getBody(), is("{}"));

        assertNotNull(wordRepository.findByWord("me"));
    }

    @Test
    public void removeList() throws Exception {
        ResponseEntity<String> entity = delete(url("user/zheka/list/study"), String.class, "zheka", "456");

        assertThat(entity.getStatusCode().value(), is(HttpStatus.NO_CONTENT.value()));
        User user = userRepository.findById("zheka").orElse(null);
        assertThat(user.getListsNames().size(), is(0));
    }

    @Test
    public void removeList_noListForUser() throws Exception {
        ResponseEntity<String> entity = delete(url("user/zheka/list/blablalist"), String.class, "zheka", "456");

        assertThat(entity.getStatusCode().value(), is(HttpStatus.NO_CONTENT.value()));
        User user = userRepository.findById("zheka").orElse(null);
        assertThat(user.getListsNames().size(), is(1));
        assertThat(user.getListsNames(), not(hasItem(hasProperty("name", equalTo("blablalist")))));
    }

    //utils
    private <T> ResponseEntity<T> get(String url, Class<T> clazz, String login, String password) throws Exception {
        HttpHeaders headers = headers(login, password);
        HttpEntity<String> requestEntity = new HttpEntity<>(headers);

        return restTemplate.exchange(url, HttpMethod.GET, requestEntity, clazz);
    }

    private ResponseEntity<String> post(String url, Object entity, String login, String password) throws Exception {
        HttpHeaders headers = headers(login, password);
        HttpEntity<String> requestEntity = new HttpEntity<>(mapper.writeValueAsString(entity), headers);

        return restTemplate.exchange(url, HttpMethod.POST, requestEntity, String.class);
    }

    private <T> ResponseEntity<T> delete(String url, Class<T> clazz, String login, String password) {
        HttpHeaders headers = headers(login, password);
        HttpEntity<String> requestEntity = new HttpEntity<>(headers);

        return restTemplate.exchange(url, HttpMethod.DELETE, requestEntity, clazz);
    }

    private HttpHeaders headers(String login, String password) {
        HttpHeaders headers = new HttpHeaders();
        headers.setContentType(MediaType.APPLICATION_JSON);
        headers.set("Authorization", "Basic " + Base64.getEncoder().encodeToString((login + ":" + password).getBytes()));
        return headers;
    }

    private String url(String part) {
        return root + part;
    }

}
