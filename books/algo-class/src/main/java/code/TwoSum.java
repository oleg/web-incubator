package code;

import java.util.HashMap;
import java.util.Map;

public class TwoSum {
    public int[] twoSum_v1(int[] nums, int target) {
        int length = nums.length;
        for (int i = 0; i < length; i++) {
            if (nums[i] <= target) {
                for (int j = i + 1; j < length; j++) {
                    if (nums[j] <= target) {
                        if (nums[i] + nums[j] == target) {
                            return new int[]{i, j};
                        }
                    }
                }
            }
        }
        return null;
    }

    public int[] twoSum(int[] nums, int target) {
        Map<Integer, Integer> valToIndex = new HashMap<>();
        for (int aIndex = 0; aIndex < nums.length; aIndex++) {
            int a = nums[aIndex];
            int b = target - a;
            Integer bIndex = valToIndex.get(b);
            if (bIndex != null) {
                return new int[]{aIndex, bIndex};
            }
            valToIndex.put(a, aIndex);
        }
        return null;
    }

}

