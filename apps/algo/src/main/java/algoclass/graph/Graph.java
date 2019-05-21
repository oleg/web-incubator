package algoclass.graph;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.Iterator;
import java.util.List;
import java.util.Map;
import java.util.Random;

//TODO make immutable?
public class Graph {
  private final Map<Integer, List<Integer>> vertices;

  private Graph(Map<Integer, List<Integer>> input) {
    ensureCorrectGraph(input);
    this.vertices = input;
  }

  public static Graph copyOf(Map<Integer, List<Integer>> input) {
    return new Graph(copyMapAndLists(input));
  }

  private static Map<Integer, List<Integer>> copyMapAndLists(Map<Integer, List<Integer>> input) {
    final Map<Integer, List<Integer>> copy = new HashMap<>();
    for (Map.Entry<Integer, List<Integer>> e : input.entrySet()) {
      copy.put(e.getKey(), new ArrayList<>(e.getValue()));
    }
    return copy;
  }

  private static void ensureCorrectGraph(Map<Integer, List<Integer>> input) {
    for (Map.Entry<Integer, List<Integer>> vertex : input.entrySet()) {
      final Integer thisId = vertex.getKey();
      final List<Integer> thisEnds = vertex.getValue();

      for (Integer endVertexId : thisEnds) {
        List<Integer> endVertexEnds = input.get(endVertexId);
        if (endVertexEnds == null || !endVertexEnds.contains(thisId)) {
          throw new IllegalStateException(endVertexId + " has no link to " + thisId);
        }
      }
    }
  }

  public int verticesCount() {
    return vertices.size();
  }

  public List<Integer> get(int id) {
    return vertices.get(id);
  }

  public Graph merge(int o, int n) {
    Map<Integer, List<Integer>> copy = copyMapAndLists(vertices);
    collapseVertices(o, n, copy);
    return new Graph(copy);
  }

  //TODO rewrite to make it possible to work with unmodifiable (or immutable) map
  private void collapseVertices(int o, int n, Map<Integer, List<Integer>> map) {
    if (o == n) {
      return;
    }

    List<Integer> removed = map.remove(o);
    map.get(n).addAll(removed);

    for (List<Integer> links : map.values()) {
      while (links.remove((Object) o)) {
        links.add(n);
      }
    }
    while (map.get(n).remove((Object) n)) ;
  }

  public int calculateMinCut() {
    int i = oneMinCutRandomlyCalculation();
    for (int j = 0; j < vertices.size() * 2; j++) {   //TODO
      i = Math.min(i, oneMinCutRandomlyCalculation());
    }
    return i;
  }

  private int oneMinCutRandomlyCalculation() {
    final Random r = new Random();

    Map<Integer, List<Integer>> copy = copyMapAndLists(vertices);
    while (copy.size() > 2) {
      final int randomVertex = randomVertex(copy, r);
      final List<Integer> links = copy.get(randomVertex);
      final int vertexFromLink = links.get(r.nextInt(links.size()));
      collapseVertices(randomVertex, vertexFromLink, copy);
    }
    return copy.values().iterator().next().size();
  }

  private static Integer randomVertex(Map<Integer, List<Integer>> vs, Random r) {
    final int i = r.nextInt(vs.size());
    final Iterator<Integer> iterator = vs.keySet().iterator();
    for (int j = 0; j < i; j++) {
      iterator.next();
    }
    return iterator.next();
  }

  @Override
  public String toString() {
    final StringBuilder res = new StringBuilder();
    for (Map.Entry<Integer, List<Integer>> ile : vertices.entrySet()) {
      res.append(ile).append("\n");
    }
    return res.toString();
  }

}
