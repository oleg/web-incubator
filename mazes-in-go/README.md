# Maze generator
Run `go run cli/cli.go -width 10 -height 10`

to generate random maze like the one below

```

+---+---+---+---+---+---+---+---+---+---+
|                                       |
+   +---+---+   +   +---+   +---+   +   +
|   |           |   |       |       |   |
+   +---+   +---+   +---+   +   +---+   +
|   |       |       |       |   |       |
+---+---+   +---+---+---+---+   +---+   +
|           |                   |       |
+---+---+---+---+---+   +---+---+   +   +
|                       |           |   |
+---+---+   +---+   +   +   +---+   +   +
|           |       |   |   |       |   |
+---+   +   +   +---+   +---+   +   +   +
|       |   |   |       |       |   |   |
+   +---+   +   +---+   +---+   +   +   +
|   |       |   |       |       |   |   |
+---+   +   +   +---+   +---+---+---+   +
|       |   |   |       |               |
+   +---+---+   +---+---+   +---+---+   +
|   |           |           |           |
+---+---+---+---+---+---+---+---+---+---+
```

Run `go run cli/cli.go --algorithm sidewinder --render png --width 100 --height 100 > maze.png`

to render as png

![Mze](maze.png)