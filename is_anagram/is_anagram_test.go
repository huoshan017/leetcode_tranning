package is_anagram

import "testing"

func TestIsAnagram(t *testing.T) {
	s1 := "anagram"
	s2 := "nagaram"
	result := isAnagram(s1, s2)
	t.Logf("result: %v", result)

	s1 = "rat"
	s2 = "car"
	result = isAnagram(s1, s2)
	t.Logf("result: %v", result)
}
