package algoclass;

import org.junit.Test;
import util.PathUtil;

import java.nio.file.Paths;

import static algoclass.MergeInversion.count;
import static org.hamcrest.CoreMatchers.is;
import static org.junit.Assert.assertThat;

public class MergeInversionsTest {

  @Test
  public void empty() throws Exception {
    assertThat(count(), is(0L));
  }

  @Test
  public void one() throws Exception {
    assertThat(count(1), is(0L));
    assertThat(count(33), is(0L));
  }

  @Test
  public void simple() throws Exception {
    assertThat(count(3, 2), is(1L));
    assertThat(count(5, 1), is(1L));
  }

  @Test
  public void example_from_book() throws Exception {
    assertThat(count(2, 3, 8, 6, 1), is(5L));
  }

  @Test
  public void my_sample() throws Exception {
    assertThat(count(5, 4, 3, 2, 1), is(10L));
  }

  @Test
  public void no_inversions() throws Exception {
    assertThat(count(1, 2, 3, 4, 5), is(0L));
  }

  @Test
  public void from_file() throws Exception {
    int[] ints = PathUtil.readAllInts(Paths.get("test/IntegerArray.txt"));
    assertThat(count(ints), is(2407905288L));
  }

}