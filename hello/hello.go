package hello

const (
	Spanish = "Spanish"
	French  = "French"

	prefixEnglish = "Hello, "
	prefixSpanish = "Hola, "
	prefixFrench  = "Bonjour, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "world"
	}

	prefix := greetingPrefix(language)
	return prefix + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case Spanish:
		prefix = prefixSpanish
	case French:
		prefix = prefixFrench
	default:
		prefix = prefixEnglish
	}
	return
}
