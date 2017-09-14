package q71

func IsStopWords(s string) bool {
	sw := []string{"a", "did", "in", "only", "then",
		"whereall", "do", "into", "onto", "there",
		"whether", "almost", "does", "is", "therefore",
		"which", "also", "either", "it", "our", "these", "while",
		"although", "for", "its", "ours", "they", "whoan",
		"from", "just", "s", "this", "whose", "had",
		"ll", "shall", "those", "why", "any", "has", "me",
		"she", "though", "will", "are", "have", "might",
		"should", "through", "with", "as", "having",
		"Mr", "since", "thus", "would", "at",
		"he", "Mrs", "so", "to", "yet", "be", "her",
		"Ms", "some", "too", "you", "because", "here", "my",
		"still", "until", "your", "been", "hers", "no", "such",
		"ve", "yours", "both", "him", "non", "t", "very",
		"but", "his", "nor", "than", "was", "by",
		"how", "not", "that", "we", "can",
		"however", "of", "the", "were", "could", "i", "on", "their",
		"what", "d", "if", "one", "them", "when"}

	for _, v := range sw {
		if v == s {
			return true
		}
	}
	return false
}
