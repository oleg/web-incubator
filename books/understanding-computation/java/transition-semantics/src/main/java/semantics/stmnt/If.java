package semantics.stmnt;

import lombok.EqualsAndHashCode;
import lombok.ToString;
import semantics.env.Environment;
import semantics.exp.Expression;
import semantics.value.BoolVal;

@ToString
@EqualsAndHashCode
public class If implements Statement {
    private final Expression<BoolVal> condition;
    private final Statement consequence;
    private final Statement alternative;

    public If(Expression<BoolVal> condition, Statement consequence, Statement alternative) {
        this.condition = condition;
        this.consequence = consequence;
        this.alternative = alternative;
    }


    @Override
    public boolean isReducible() {
        return true;
    }

    @Override
    public StatementResult reduce(Environment environment) {
        if (condition.isReducible()) {
            return new StatementResult(new If(condition.reduce(environment), consequence, alternative), environment);
        }
        if (condition.getReduced().get()) {
            return new StatementResult(consequence, environment);
        } else {
            return new StatementResult(alternative, environment);
        }
    }
}
