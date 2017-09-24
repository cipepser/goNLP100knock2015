package q71

func IsStopWords(s string) bool {
	// sw := []string{"a", "did", "in", "only", "then",
	// 	"whereall", "do", "into", "onto", "there",
	// 	"whether", "almost", "does", "is", "therefore",
	// 	"which", "also", "either", "it", "our", "these", "while",
	// 	"although", "for", "its", "ours", "they", "whoan",
	// 	"from", "just", "s", "this", "whose", "had",
	// 	"ll", "shall", "those", "why", "any", "has", "me",
	// 	"she", "though", "will", "are", "have", "might",
	// 	"should", "through", "with", "as", "having",
	// 	"Mr", "since", "thus", "would", "at",
	// 	"he", "Mrs", "so", "to", "yet", "be", "her",
	// 	"Ms", "some", "too", "you", "because", "here", "my",
	// 	"still", "until", "your", "been", "hers", "no", "such",
	// 	"ve", "yours", "both", "him", "non", "t", "very",
	// 	"but", "his", "nor", "than", "was", "by",
	// 	"how", "not", "that", "we", "can",
	// 	"however", "of", "the", "were", "could", "i", "on", "their",
	// 	"what", "d", "if", "one", "them", "when", "or"}

	// https://pypi.python.org/pypi/stop-words
	sw := []string{"a", "about", "above", "after", "again", "against", "all",
		"am", "an", "and", "any", "are", "aren't", "as", "at",
		"be", "because", "been", "before", "being", "below",
		"between", "both", "but", "by", "can't", "cannot",
		"could", "couldn't", "did", "didn't", "do", "does", "doesn't",
		"doing", "don't", "down", "during", "each", "few", "for", "from",
		"further", "had", "hadn't", "has", "hasn't", "have", "haven't", "having",
		"he", "he'd", "he'll", "he's", "her", "here", "here's", "hers", "herself",
		"him", "himself", "his", "how", "how's", "i", "i'd", "i'll", "i'm",
		"i've", "if", "in", "into", "is", "isn't", "it", "it's", "its",
		"itself", "let's", "me", "more", "most", "mustn't", "my", "myself", "no",
		"nor", "not", "of", "off", "on", "once", "only", "or", "other", "ought",
		"our", "ours", "ourselves", "out", "over", "own", "same", "shan't", "she",
		"she'd", "she'll", "she's", "should", "shouldn't", "so", "some", "such",
		"than", "that", "that's", "the", "their", "theirs", "them", "themselves",
		"then", "there", "there's", "these", "they", "they'd", "they'll",
		"they're", "they've", "this", "those", "through", "to", "too",
		"under", "until", "up", "very", "was", "wasn't", "we", "we'd",
		"we'll", "we're", "we've", "were", "weren't", "what", "what's",
		"when", "when's", "where", "where's", "which", "while", "who", "who's",
		"whom", "why", "why's", "with", "won't", "would", "wouldn't", "you", "you'd",
		"you'll", "you're", "you've", "your", "yours", "yourself", "yourselves",
		"will"}

	for _, v := range sw {
		if v == s {
			return true
		}
	}
	return false
}
