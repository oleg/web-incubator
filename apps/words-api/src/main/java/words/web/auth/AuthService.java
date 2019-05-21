package words.web.auth;

import org.springframework.stereotype.Service;
import words.app.model.User;
import words.app.repository.UserRepository;

import javax.inject.Inject;

@Service
public class AuthService {

    private final UserRepository userRepository;

    @Inject
    public AuthService(UserRepository userRepository) {
        this.userRepository = userRepository;
    }

    public boolean isUserExist(String login, String password) {
        User user = userRepository.findOne(login);
        return user != null && user.getPassword().equals(password);
    }

}
