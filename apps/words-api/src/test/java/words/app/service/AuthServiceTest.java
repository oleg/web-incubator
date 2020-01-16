package words.app.service;

import org.junit.Before;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.runners.MockitoJUnitRunner;
import words.app.model.User;
import words.app.repository.UserRepository;
import words.web.auth.AuthService;

import java.util.Optional;

import static org.hamcrest.CoreMatchers.is;
import static org.junit.Assert.assertThat;
import static org.mockito.Mockito.when;

@RunWith(MockitoJUnitRunner.class)
public class AuthServiceTest {

    @Mock UserRepository userRepository;
    @InjectMocks AuthService authService;

    @Before
    public void setUp() throws Exception {
        when(userRepository.findById("admin")).thenReturn(Optional.of(new User("admin", "admin")));
    }

    @Test
    public void ok() throws Exception {
        assertThat(authService.isUserExist("admin", "admin"), is(true));
    }

    @Test
    public void fail() throws Exception {
        assertThat(authService.isUserExist("admin", "user"), is(false));
    }

    @Test
    public void npe() throws Exception {
        assertThat(authService.isUserExist("user", "user"), is(false));
    }
}
