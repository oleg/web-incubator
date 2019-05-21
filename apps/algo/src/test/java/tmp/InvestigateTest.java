package tmp;

import org.junit.Test;

import java.util.List;

import static java.util.Arrays.asList;

public class InvestigateTest {

    public interface Visitable {

        void accept(Visitor visitor);

    }

    public static class Book implements Visitable {
        String name;

        public Book(String name) {
            this.name = name;
        }

        @Override
        public void accept(Visitor visitor) {
            visitor.visitBook(this);
        }
    }

    public static class Cd implements Visitable {
        String title;

        public Cd(String title) {
            this.title = title;
        }

        @Override
        public void accept(Visitor visitor) {
            visitor.visitCd(this);
        }
    }


    public interface Visitor {

        void visitBook(Book book);

        void visitCd(Cd cd);

    }

    public static class NameCollectorVisitor implements Visitor {
        String description = "Shopping card: \n";

        @Override
        public void visitBook(Book book) {
            description += "=> Book " + book.name + "\n";
        }

        @Override
        public void visitCd(Cd cd) {
            description += "=> Cd " + cd.title + "\n";
        }

        public String getDescription() {
            return description;
        }
    }

    @Test
    public void test_visitor_pattern() throws Exception {
        NameCollectorVisitor visitor = new NameCollectorVisitor();

        List<Visitable> shoppingCard = asList(
            new Book("Peace and war"),
            new Cd("How much is the fish"),
            new Book("Adventure time"),
            new Cd("Lap story"),
            new Cd("Boom!"));

        shoppingCard.forEach(e -> e.accept(visitor));

        System.out.println(visitor.getDescription());
    }
}