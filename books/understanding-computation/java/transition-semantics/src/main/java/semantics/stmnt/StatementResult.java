package semantics.stmnt;

import lombok.Value;
import semantics.env.Environment;

@Value
public class StatementResult {
    private final Statement statement;
    private final Environment environment;
}
