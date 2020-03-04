package lcode;

public class ReverseInt {

    public int reverse_v1(int x) {
        StringBuilder str = new StringBuilder(Integer.toString(x));
        str.reverse();
        if (x < 0) {
            str.deleteCharAt(str.length() - 1);
            str.insert(0, "-");
        }
        try {
            return Integer.parseInt(str.toString());
        } catch (NumberFormatException e) {
            return 0;
        }
    }

    public int reverse_v2(int x) {
        char[] val = Integer.toString(x).toCharArray();
        int extr = x < 0 ? 1 : 0;
        int len = val.length - extr;
        int n = len - 1;
        for (int j = (n - 1) >> 1; j >= 0; j--) {
            int mj = j + extr;
            int k = n - j + extr;

            char cj = val[mj];
            val[mj] = val[k];
            val[k] = cj;
        }
        try {
            return Integer.parseInt(String.valueOf(val));
        } catch (NumberFormatException e) {
            return 0;
        }
    }

    public int reverse(int x) {
        int xm;
        short sign;
        if (x < 0) {
            sign = -1;
            xm = -x;
        } else {
            sign = 1;
            xm = x;
        }
        char[] val = Integer.toString(xm).toCharArray();
        int n = val.length - 1;
        for (int j = (n - 1) >> 1; j >= 0; j--) {
            int k = n - j;
            char cj = val[j];
            val[j] = val[k];
            val[k] = cj;
        }
        try {
            return sign * Integer.parseInt(String.valueOf(val));
        } catch (NumberFormatException e) {
            return 0;
        }
    }
}
