package words.app.service;

import org.junit.Test;

import static org.hamcrest.CoreMatchers.is;
import static org.junit.Assert.assertThat;

public class CleanupStrategyTest {

    private CleanupStrategy cleanupStrategy = new CleanupStrategy();

    @Test
    public void should_remove_punctuation() throws Exception {
        assertThat(cleanupStrategy.toWordsList("hello.").get(0), is("hello"));
        assertThat(cleanupStrategy.toWordsList(".world").get(0), is("world"));

        assertThat(cleanupStrategy.toWordsList(",other,").get(0), is("other"));

        assertThat(cleanupStrategy.toWordsList("'quotes\"").get(0), is("quotes"));
    }

    @Test
    public void should_convert_to_lowercase() throws Exception {
        assertThat(cleanupStrategy.toWordsList("What").get(0), is("what"));
    }

    @Test
    public void should_treat_dots_as_separators() throws Exception {
        assertThat(cleanupStrategy.toWordsList("do.what").get(0), is("do"));
        assertThat(cleanupStrategy.toWordsList("do.what").get(1), is("what"));
    }

    @Test
    public void should_new_lines_as_separators() throws Exception {
        assertThat(cleanupStrategy.toWordsList("line\nbreak").get(0), is("line"));
        assertThat(cleanupStrategy.toWordsList("line\nbreak").get(1), is("break"));
    }

}
