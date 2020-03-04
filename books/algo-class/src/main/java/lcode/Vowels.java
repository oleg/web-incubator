package lcode;

public class Vowels {
    public String removeVowels_v1(String s) {
        return s.replaceAll("[aeiou]*", "");
    }

    public String removeVowels(String s) {
        StringBuilder builder = new StringBuilder();
        for (char c : s.toCharArray()) {
            if (c != 'a' && c != 'e' && c != 'i' && c != 'o' && c != 'u') {
                builder.append(c);
            }
        }
        return builder.toString();
    }
}
