package code;

import java.util.ArrayDeque;

public class ValidParentheses {

    public boolean isValid(String str) {
        var deq = new ArrayDeque<Character>();
        for (char curr : str.toCharArray()) {
            if (curr == '(' || curr == '[' || curr == '{') {
                deq.push(curr);
            } else if (curr == ')' || curr == ']' || curr == '}') {
                Character prev = deq.poll();
                if (prev == null) {
                    return false;
                }
                if ((curr == ')' && prev != '(') ||
                    (curr == ']' && prev != '[') ||
                    (curr == '}' && prev != '{')) {
                    return false;
                }
            } else {
                throw new IllegalArgumentException("unexpected char " + curr);
            }
        }
        return deq.isEmpty();
    }

}
