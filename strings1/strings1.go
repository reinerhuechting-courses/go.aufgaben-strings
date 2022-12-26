package strings1

import "strings"

// Erwartet einen string s und zählt, wie oft der Buchstabe 'A' in s vorkommt.
func CountA(s string) int {
	// ANMERKUNG: Diese Funktion ist ein Beispiel, hier ist (noch) nichts zu tun.
	result := 0
	for _, char := range s {
		if char == 'A' {
			result++
		}
	}
	return result
}

// Erwartet einen string s und einen Buchstaben c.
// Zählt, wie oft c in s vorkommt.
func CountChar(s string, c rune) int {
	result := 0
	for _, char := range s {
		if char == c {
			result++
		}
	}
	return result
	// ZUSATZAUFGABE: Ändern Sie anschließend CountA(),
	// so dass insgesamt möglichst wenig doppelter Code existiert.
}

// Erwartet einen String s und liefert einen neuen String,
// in dem jeder Buchstabe aus s zweimal hintereinander vorkommt.
func DuplicateChars(s string) string {
	// ANMERKUNG: Diese Funktion ist ein Beispiel, hier ist nichts zu tun.
	result := ""
	for _, char := range s {
		result += string(char)
		result += string(char)
	}
	return result
}

// Erwartet einen String s und liefert s rückwärts zurück.
func Reverse(s string) string {
	result := ""
	for i := len(s) - 1; i >= 0; i-- {
		result += string(s[i])
	}
	return result
}

// Erwartet zwei Strings s1 und s2 und prüft, ob der eine der andere umgedreht ist.
func IsReverse(s1, s2 string) bool {
	return s1 == Reverse(s2)
}

// Erwartet einen String s und prüft, ob s ein Palindrom ist.
// Ein Palindrom ist ein String, der vorwärts und rückwärts gelesen gleich ist.
func IsPalindrome(s string) bool {
	return s == Reverse(s)
}

// Erwartet zwei Strings s1 und s2 und prüft, ob die beiden Anagramme voneinander sind.
// Ein Anagramm von s1 ist ein String, der exakt die gleichen Buchstaben wie s1
// enthält, aber nicht unbedingt in der gleichen Reihenfolge.
func IsAnagram(s1, s2 string) bool {
	for _, c := range s1 {
		if CountChar(s1, c) != CountChar(s2, c) {
			return false
		}
	}
	return true
}

// Erwartet zwei Strings s1 und s2 und prüft, ob die beiden Anagramme voneinander sind.
// Ein Anagramm von s1 ist ein String, der exakt die gleichen Buchstaben wie s1
// enthält, aber nicht unbedingt in der gleichen Reihenfolge.
// Diese Funktion soll keinen Unterschied zwischen Groß- und Kleinschreibung machen.
func IsAnagramIgnoreCase(s1, s2 string) bool {
	// HINWEIS: Verwenden Sie die Funktion strings.ToLower() oder strings.ToUpper()
	return IsAnagram(strings.ToLower(s1), strings.ToLower(s2))
}
