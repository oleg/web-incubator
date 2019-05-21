package words.app.repository;

import org.springframework.data.repository.PagingAndSortingRepository;
import org.springframework.stereotype.Repository;
import words.app.model.WordsList;

import java.util.Collection;

/**
 * Created by oleg.prozorov on 16.07.2015.
 */
@Repository
public interface WordsListRepository extends PagingAndSortingRepository<WordsList, String> {

    String findOneByName(String name);

    Collection<WordsList> findByNameIn(Collection<String> names);

}
