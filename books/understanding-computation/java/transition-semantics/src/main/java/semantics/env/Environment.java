package semantics.env;

import lombok.EqualsAndHashCode;
import lombok.ToString;
import semantics.value.Val;

import java.util.Collections;
import java.util.HashMap;
import java.util.Map;

@EqualsAndHashCode
@ToString
public class Environment {

    private final Map<String, Val<?>> state;

    public Environment() {
        this(Collections.emptyMap());
    }

    private Environment(Map<String, Val<?>> state) {
        this.state = Collections.unmodifiableMap(state);
    }

    public Environment merge(String name, Val<?> vl) {
        Map<String, Val<?>> newState = new HashMap<>(state);
        newState.put(name, vl);
        return new Environment(newState);
    }

    @SuppressWarnings("unchecked")
    public <T> Val<T> get(String name) {
        return (Val<T>) state.get(name);
    }
}
