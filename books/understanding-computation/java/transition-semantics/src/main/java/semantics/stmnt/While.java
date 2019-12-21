package semantics.stmnt;

import lombok.EqualsAndHashCode;
import lombok.ToString;
import semantics.env.Environment;
import semantics.exp.Expression;
import semantics.value.BoolVal;

@EqualsAndHashCode
@ToString
public class While implements Statement {
    private final Expression<BoolVal> condition;
    private final Statement body;

    public While(Expression<BoolVal> condition, Statement body) {
        this.condition = condition;
        this.body = body;
    }

    @Override
    public boolean isReducible() {
        return true;
    }

    @Override
    public StatementResult reduce(Environment environment) {
        return new StatementResult(
                new If(condition, new Sequence(body, this), DoNothing.INSTANCE),
                environment);
    }
}
