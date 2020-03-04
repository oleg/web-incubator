package lcode;

import java.util.Arrays;

public class TwoSumLessThanK {
    public int twoSumLessThanK(int[] data, int k) {
        Arrays.sort(data);
        int start = 0;
        int end = data.length - 1;
        int max = -1;
        while (start < end) {
            int sum = data[start] + data[end];
            if (sum < k) {
                if (sum > max) {
                    max = sum;
                }
                start++;
            } else {
                end--;
            }
        }
        return max;
    }

    public int twoSumLessThanK_v1(int[] data, int k) {
        int maxSum = -1;
        int len = data.length;
        for (int i = 0; i < len; i++) {
            for (int j = i; j < len; j++) {
                if (data[i] < k && data[j] < k && i != j) {
                    int sum = data[i] + data[j];
                    if (sum < k && sum > maxSum) {
                        maxSum = sum;
                    }
                }
            }
        }
        return maxSum;
    }
}
