package algoclass.graph;

import org.junit.Test;
import util.PathUtil;

import java.nio.file.Paths;

import static algoclass.graph.GraphBuilder.graph;
import static org.hamcrest.CoreMatchers.is;
import static org.junit.Assert.assertThat;

public class MinCutTest {

  @Test
  public void test_1() throws Exception {
    Graph graph = graph()
        .vertex(1, 2, 3)
        .vertex(2, 1, 3, 4)
        .vertex(3, 1, 2, 4)
        .vertex(4, 2, 3).build();

    assertThat(graph.calculateMinCut(), is(2));
  }

  @Test
  public void test_2() throws Exception {
    Graph graph = graph()
        .vertex(1, 2)
        .vertex(2, 1, 3, 4)
        .vertex(3, 2, 4)
        .vertex(4, 2, 3).build();

    assertThat(graph.calculateMinCut(), is(1));
  }

  @Test
  public void test_big() throws Exception {
    Graph graph = graph()
        .vertex(1, 2, 3, 7)
        .vertex(2, 1, 3, 4, 7)
        .vertex(3, 1, 2, 6, 7)
        .vertex(4, 2, 5, 6, 8)
        .vertex(5, 4, 6, 8)
        .vertex(6, 3, 4, 5, 8)
        .vertex(7, 2, 1, 3)
        .vertex(8, 4, 5, 6).build();

    assertThat(graph.calculateMinCut(), is(2));
  }

  @Test
  public void test_3() throws Exception {
    Graph graph = PathUtil.readGraph(Paths.get("test/kargerMinCut.txt"));
    assertThat(graph.calculateMinCut(), is(17));
  }
}
