package shad.gui;

import shad.NetPath;
import shad.Permutations;

import javax.swing.*;
import java.awt.*;
import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;
import java.awt.event.MouseAdapter;
import java.awt.event.MouseEvent;
import java.util.ArrayList;
import java.util.Collections;
import java.util.Comparator;
import java.util.HashSet;
import java.util.List;
import java.util.Set;

import static shad.NetPath.Direction.DOWN;
import static shad.NetPath.Direction.LEFT;

public class DrawNetPaths extends JFrame {
  private final List<List<NetPath.Direction>> directions;
  private int current;

  public DrawNetPaths(Set<List<NetPath.Direction>> directions) {
    super("Draw Line");
    this.directions = new ArrayList<>(directions);

    Collections.sort(this.directions, new Comparator<List<NetPath.Direction>>() {
      @Override
      public int compare(List<NetPath.Direction> o1, List<NetPath.Direction> o2) {
        for (int i = 0; i < o1.size(); i++) {
          final NetPath.Direction a = o1.get(i);
          final NetPath.Direction b = o2.get(i);
          if (a != b) {
            return a == DOWN ? 1 : -1;
          }
        }
        return 0;
      }
    });

    final MouseAdapter l = new MyMouseAdapter();
    final DrawPane contentPane = new DrawPane();
    contentPane.addMouseListener(l);
    setContentPane(contentPane);
    setDefaultCloseOperation(WindowConstants.EXIT_ON_CLOSE);
    setSize(400, 400);
    setVisible(true);

    final MenuBar mb = new MenuBar();
    final Menu m = new Menu("Paths");
    for (int i = 0; i < this.directions.size(); i++) {
      final MenuItem item = new MenuItem((i + 1) + this.directions.get(i).toString());
      final int finalI = i;
      item.addActionListener(new ActionListener() {
        @Override
        public void actionPerformed(ActionEvent e) {
          current = finalI;
          repaint();
        }
      });
      m.add(item);
    }
    mb.add(m);
    setMenuBar(mb);
  }

  class DrawPane extends JPanel {
    public void paintComponent(Graphics g) {
      int xf = 20;
      int yf = 20;
      int xt = 20;
      int yt = 20;
      for (NetPath.Direction direction : directions.get(current)) {
        if (direction == LEFT) {
          xt += 100;
        } else {
          yt += 100;
        }
        g.drawLine(xf, yf, xt, yt);
        xf = xt;
        yf = yt;
      }
      g.drawString((current + 1) + "/" + directions.size() + ": " + directions.get(current).toString(), 10, 10);
    }
  }

  private class MyMouseAdapter extends MouseAdapter {
    @Override
    public void mouseClicked(MouseEvent e) {
      if (current == directions.size() - 1) {
        current = 0;
      } else {
        current++;
      }
      repaint();
    }
  }

  public static void main(String args[]) {
    final List<List<NetPath.Direction>> lists = Permutations.get(DOWN, DOWN, DOWN, LEFT, LEFT, LEFT);
    new DrawNetPaths(new HashSet<>(lists));
  }
}
