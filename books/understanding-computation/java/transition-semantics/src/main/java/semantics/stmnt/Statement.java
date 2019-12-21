package semantics.stmnt;

import semantics.env.Environment;

public interface Statement {

    boolean isReducible();

    StatementResult reduce(Environment environment);
}
