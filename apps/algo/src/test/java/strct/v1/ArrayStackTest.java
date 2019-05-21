package strct.v1;

import org.junit.Test;

import static org.hamcrest.CoreMatchers.is;
import static org.junit.Assert.assertThat;

public class ArrayStackTest {

    @Test
    public void get_at_should_return_element_at_position() throws Exception {
        ArrayStack<String> stack = new ArrayStack<>(String.class, 2);
        stack.push("world");
        stack.push("hello");

        assertThat(stack.pop(), is("hello"));
        assertThat(stack.pop(), is("world"));
    }

    @Test(expected = IllegalStateException.class)
    public void pop_should_throw_error_if_stack_is_empty() throws Exception {
        new ArrayStack<>(String.class, 10).pop();
    }

    @Test(expected = IllegalStateException.class)
    public void pop_should_throw_error_if_stack_is_overflow() throws Exception {
        ArrayStack<String> stack = new ArrayStack<>(String.class, 2);
        stack.push("aaa");
        stack.push("bbb");
        stack.push("ccc");
    }

}