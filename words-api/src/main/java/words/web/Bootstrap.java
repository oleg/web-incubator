package words.web;

import lombok.extern.slf4j.Slf4j;
import org.springframework.context.annotation.Profile;
import org.springframework.stereotype.Component;
import words.app.model.User;
import words.app.model.WordsList;
import words.app.repository.UserRepository;

import javax.annotation.PostConstruct;
import javax.inject.Inject;
import javax.transaction.Transactional;

@Slf4j
@Component
@Profile("!test")
public class Bootstrap {

    @Inject BootstrapJob bootstrapJob;

    @PostConstruct
    public void init() {
        log.info("invoking bootstrap");
        bootstrapJob.bootstrap();
    }

    @Slf4j
    @Component
    @Profile("!test")
    public static class BootstrapJob {

        private final UserRepository userRepository;

        @Inject
        public BootstrapJob(UserRepository userRepository) {
            this.userRepository = userRepository;
        }

        @Transactional
        public void bootstrap() {
            log.info("Bootstrap started");

            addNewUser("oleg", "123");
            addNewUser("zheka", "456");

            log.info("Bootstrap ended");
        }

        private void addNewUser(String name, String password) {
            User user = getOrCreate(name, password);

            if (user.getListsNames() == null || user.getListsNames().isEmpty()) {
                user.addWordsLists(
                        new WordsList("study"),
                        new WordsList("known"),
                        new WordsList("ignore"));
            }

            userRepository.save(user);
        }

        private User getOrCreate(String name, String password) {
            User user = userRepository.findByName(name);
            if (user == null) {
                user = new User(name, password);
            }
            return user;
        }
    }
}
