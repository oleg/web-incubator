package util;

import algoclass.graph.Graph;
import algoclass.graph.GraphBuilder;

import java.io.IOException;
import java.nio.charset.StandardCharsets;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.Arrays;
import java.util.List;

import static java.nio.charset.StandardCharsets.UTF_8;

public class PathUtil {

  public static int[] readAllInts(Path path) throws IOException {
    final List<String> lines = Files.readAllLines(path, UTF_8);
    int[] result = new int[lines.size()];
    for (int i = 0; i < result.length; i++) {
      result[i] = Integer.parseInt(lines.get(i));
    }
    return result;
  }

  public static Graph readGraph(Path path) throws IOException {
    GraphBuilder graph = GraphBuilder.graph();
    final List<String> lines = Files.readAllLines(path, UTF_8);

    for (String line : lines) {
      String[] split = line.split("\\s");
      graph.vertex(Integer.parseInt(split[0].trim()), Arrays.copyOfRange(toInt(split), 1, split.length));
    }
    return graph.build();
  }

  public static int[] toInt(String[] strings) {
    int[] result = new int[strings.length];
    for (int i = 0; i < strings.length; i++) {
      result[i] = Integer.parseInt(strings[i].trim());
    }
    return result;
  }
}
