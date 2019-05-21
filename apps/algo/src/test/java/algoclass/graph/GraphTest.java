package algoclass.graph;

import org.junit.Test;

import static algoclass.graph.GraphBuilder.graph;
import static org.hamcrest.CoreMatchers.is;
import static org.junit.Assert.assertThat;
import static org.junit.matchers.JUnitMatchers.hasItems;

public class GraphTest {

  @Test
  public void vertices_count() throws Exception {

    final Graph graph = graph()
        .vertex(1, 2, 3)
        .vertex(2, 3, 1)
        .vertex(3, 2, 1).build();

    assertThat(graph.verticesCount(), is(3));
  }

  @Test
  public void vertices_count_2() throws Exception {
    final Graph graph = graph()
        .vertex(1, 2)
        .vertex(2, 1).build();

    assertThat(graph.verticesCount(), is(2));
  }

  @Test
  public void vertices_collapse() throws Exception {
    final Graph graph = graph()
        .vertex(1, 2, 3)
        .vertex(2, 1)
        .vertex(3, 1)
        .build();

    final Graph merged = graph.merge(1, 2);

    assertThat(merged.verticesCount(), is(2));

    assertThat(merged.get(2), is(hasItems(3)));
    assertThat(merged.get(3), is(hasItems(2)));
  }

  @Test
  public void vertices_collapse_2() throws Exception {
    final Graph graph = graph()
        .vertex(1, 2, 3, 4)
        .vertex(2, 1, 4)
        .vertex(3, 1, 4)
        .vertex(4, 1, 2, 3)
        .build();

    final Graph merged = graph.merge(4, 1);

    assertThat(merged.verticesCount(), is(3));

    assertThat(merged.get(1), is(hasItems(2, 3)));
    assertThat(merged.get(2), is(hasItems(1)));
    assertThat(merged.get(3), is(hasItems(1)));
  }

  @Test(expected = IllegalStateException.class)
  public void incorrect_1() throws Exception {
    graph()
        .vertex(1, 3)
        .vertex(2, 1).build();
  }
}

