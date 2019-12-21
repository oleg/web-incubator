package semantics.stmnt;

import lombok.EqualsAndHashCode;
import lombok.ToString;
import semantics.env.Environment;
import semantics.exp.Expression;
import semantics.value.Val;

@ToString
@EqualsAndHashCode
public class Assign<T extends Val<?>> implements Statement {
    private final String name;
    private final Expression<T> expression;

    public Assign(String name, Expression<T> expression) {
        this.name = name;
        this.expression = expression;
    }

    @Override
    public boolean isReducible() {
        return true;
    }

    @Override
    public StatementResult reduce(Environment environment) {
        if (expression.isReducible())
            return new StatementResult(new Assign<>(name, expression.reduce(environment)), environment);

        T val = expression.getReduced();
        return new StatementResult(DoNothing.INSTANCE, environment.merge(name, val));
    }
}
