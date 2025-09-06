package processor

import (
	"regexp"
	"strings"
)

// SanitizeText removes escape sequences and formatting artifacts from a string.
func SanitizeText(text string) string {
	// Replace common escape sequences AND remove asterisks
	replacer := strings.NewReplacer(
		"\\n", "\n",
		"\\t", "\t",
		"\\\"", "\"",
		"\\*", "", // This was changed to remove single asterisks
		"*", "",   // This removes any remaining asterisks
	)
	cleanedText := replacer.Replace(text)

	// This regex for multiple asterisks is now redundant but harmless.
	re := regexp.MustCompile(`\*{2,}`)
	cleanedText = re.ReplaceAllString(cleanedText, "")

	return cleanedText
}