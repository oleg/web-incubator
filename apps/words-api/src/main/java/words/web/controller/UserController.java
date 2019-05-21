package words.web.controller;

import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;
import words.app.model.User;
import words.app.repository.UserRepository;

import javax.inject.Inject;

@RestController
@RequestMapping(value = "api/user")
public class UserController {

    private final UserRepository userRepository;

    @Inject
    public UserController(UserRepository userRepository) {
        this.userRepository = userRepository;
    }

    @RequestMapping(method = RequestMethod.GET, value = "{user}")
    public ResponseEntity<?> readUser(@PathVariable String user) {
        User u = userRepository.findOne(user);
        if (u != null) {
            return ResponseEntity.ok(u);
        } else {
            return ResponseEntity.notFound().build();
        }
    }

}
