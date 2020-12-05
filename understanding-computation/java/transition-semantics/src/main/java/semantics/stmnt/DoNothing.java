package semantics.stmnt;

import lombok.EqualsAndHashCode;
import lombok.ToString;
import semantics.env.Environment;

@EqualsAndHashCode
@ToString
public class DoNothing implements Statement {

    public static final DoNothing INSTANCE = new DoNothing();

    private DoNothing() {
    }

    @Override
    public boolean isReducible() {
        return false;
    }

    @Override
    public StatementResult reduce(Environment environment) {
        throw new UnsupportedOperationException();
    }
}
