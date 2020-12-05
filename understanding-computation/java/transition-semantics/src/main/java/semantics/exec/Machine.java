package semantics.exec;

import semantics.env.Environment;
import semantics.stmnt.Statement;
import semantics.stmnt.StatementResult;

public class Machine {

    public Environment run(Statement statement,
                           Environment environment) {
        StatementResult result;
        while (statement.isReducible()) {
            result = statement.reduce(environment);
            statement = result.getStatement();
            environment = result.getEnvironment();
        }
        return environment;
    }

}
