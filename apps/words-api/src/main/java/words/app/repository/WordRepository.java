package words.app.repository;

import org.springframework.data.repository.PagingAndSortingRepository;
import org.springframework.stereotype.Repository;
import words.app.model.Word;

/**
 * Created by oleg.prozorov on 16.07.2015.
 */
@Repository
public interface WordRepository extends PagingAndSortingRepository<Word, String> {

    Word findByWord(String word);

}

