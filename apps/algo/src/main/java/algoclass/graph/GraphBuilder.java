package algoclass.graph;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class GraphBuilder {
  private final Map<Integer, List<Integer>> map = new HashMap<>();

  public static GraphBuilder graph() {
    return new GraphBuilder();
  }

  public GraphBuilder vertex(int id, int... links) {
    map.put(id, toList(links));
    return this;
  }

  public Graph build() {
    return Graph.copyOf(map);
  }


  //TODO how to write it short
  private List<Integer> toList(int[] links) {
    //List<Integer> ints = Arrays.asList(links); -- wtf?
    final List<Integer> result = new ArrayList<>();
    for (int link : links) {
      result.add(link);
    }
    return result;
  }

}
