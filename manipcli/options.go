package manipcli

// generateCommandConfig hods the configuration to generate commands.
type cmdConfig struct {
	argumentsPrefix string
}

// Option represents an option can for the generate command.
type cmdOption func(*cmdConfig)

// optionArgumentsPrefix sets the argument prefixes.
func optionArgumentsPrefix(prefix string) cmdOption {
	return func(g *cmdConfig) {

		g.argumentsPrefix = prefix
	}
}
