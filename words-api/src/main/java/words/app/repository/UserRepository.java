package words.app.repository;

import org.springframework.data.repository.PagingAndSortingRepository;
import org.springframework.stereotype.Repository;
import words.app.model.User;

import javax.transaction.Transactional;

@Repository
public interface UserRepository extends PagingAndSortingRepository<User, String> {

    @Transactional(value = Transactional.TxType.SUPPORTS)
    User findByName(String name);

}
