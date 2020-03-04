package lcode;

public class Jewels {
    public int numJewelsInStones(String J, String S) {
        char[] ss = S.toCharArray();
        char[] js = J.toCharArray();

        int count = 0;
        for (char s : ss)
            if (contains(s, js))
                count++;

        return count;
    }

    private boolean contains(char s, char[] js) {
        for (char j : js)
            if (j == s)
                return true;
        return false;
    }
}
