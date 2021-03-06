package hello

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greeting(language) + name
}

func greeting(language string) (prefix string) {
	switch language {
	case "frances":
		prefix = "Bonjour "
	case "portugues":
		prefix = "Ola "
	default:
		prefix = "Hello "
	}
	return
}