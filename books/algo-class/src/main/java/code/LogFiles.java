package code;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.Comparator;
import java.util.List;

public class LogFiles {
    public String[] reorderLogFiles_v1(String[] logs) {
        List<String> letterLogs = new ArrayList<>(logs.length);
        List<String> digitLogs = new ArrayList<>(logs.length);

        for (String log : logs) {
            if (isLetterLog(log)) {
                letterLogs.add(log);
            } else {
                digitLogs.add(log);
            }
        }

        letterLogs.sort(Comparator.comparing(log -> {
            int i = log.indexOf(" ");
            return log.substring(i) + log.substring(0, i);
        }));

        letterLogs.addAll(digitLogs);
        return letterLogs.toArray(new String[logs.length]);
    }

    private boolean isLetterLog(String log) {
        int i = log.indexOf(" ");
        return !Character.isDigit(log.charAt(i + 1));
    }


    public String[] reorderLogFiles(String[] logs) {
        Arrays.sort(logs, (l1, l2) -> {
            String[] ls1 = l1.split(" ", 2);
            String l1Key = ls1[0];
            String l1Value = ls1[1];
            boolean l1Digit = Character.isDigit(l1Value.charAt(0));

            String[] ls2 = l2.split(" ", 2);
            String l2Key = ls2[0];
            String l2Value = ls2[1];
            boolean l2Digit = Character.isDigit(l2Value.charAt(0));

            if (l1Digit || l2Digit) {
                if (!l2Digit) {
                    return 1;
                }
                if (!l1Digit) {
                    return -1;
                }
                return 0;
            }

            int compare = l1Value.compareTo(l2Value);
            if (compare != 0) {
                return compare;
            }
            return l1Key.compareTo(l2Key);
        });
        return logs;
    }

}
