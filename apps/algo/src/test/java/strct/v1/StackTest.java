package strct.v1;

import org.junit.Test;

import static org.hamcrest.CoreMatchers.is;
import static org.junit.Assert.assertThat;

public class StackTest {

    @Test
    public void get_at_should_return_element_at_position() throws Exception {
        Stack<String> stack = new Stack<>();
        stack.push("world");
        stack.push("hello");

        assertThat(stack.pop(), is("hello"));
        assertThat(stack.pop(), is("world"));
    }

    @Test(expected = IllegalStateException.class)
    public void pop_should_throw_error_if_stack_is_empty() throws Exception {
        new Stack<>().pop();
    }
}