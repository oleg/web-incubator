package code;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;

class VowelsTest {
    @Test
    void test1() {
        assertEquals("ltcdscmmntyfrcdrs", new Vowels().removeVowels("leetcodeisacommunityforcoders"));
    }

    @Test
    void test2() {
        assertEquals("", new Vowels().removeVowels("aeiou"));
    }
}