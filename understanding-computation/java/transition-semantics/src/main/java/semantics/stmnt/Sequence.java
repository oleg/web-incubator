package semantics.stmnt;

import lombok.EqualsAndHashCode;
import lombok.ToString;
import semantics.env.Environment;

@EqualsAndHashCode
@ToString
public class Sequence implements Statement {
    private final Statement first;
    private final Statement second;

    public Sequence(Statement first, Statement second) {
        this.first = first;
        this.second = second;
    }

    @Override
    public boolean isReducible() {
        return true;
    }

    @Override
    public StatementResult reduce(Environment environment) {
        if (first.isReducible()) {
            StatementResult reduced = first.reduce(environment);
            return new StatementResult(
                    new Sequence(reduced.getStatement(), second),
                    reduced.getEnvironment());
        }
        return new StatementResult(second, environment);
    }
}
