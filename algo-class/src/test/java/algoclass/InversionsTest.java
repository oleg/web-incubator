package algoclass;

import org.junit.jupiter.api.Test;
import util.PathUtil;

import java.io.*;
import java.net.URI;
import java.net.URL;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;

import static algoclass.Inversions.count;
import static org.hamcrest.CoreMatchers.is;
import static org.hamcrest.CoreMatchers.is;
import static org.hamcrest.MatcherAssert.assertThat;


public class InversionsTest {

  @Test
  public void simple() throws Exception {
    assertThat(count(3, 2), is(1L));
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
    int[] ints = PathUtil.readAllInts("IntegerArray.txt");
    assertThat(count(ints), is(2407905288L));
  }

}