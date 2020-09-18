package compiler

import "go.aporeto.io/elemental"

type CompilerConfig struct {
	TranslateKeysFromSpec bool
	AttrSpecs             map[string]elemental.AttributeSpecification
}

type CompilerOption func(*CompilerConfig) *CompilerConfig

func CompilerOptionTranslateKeysFromSpec(attrSpecs map[string]elemental.AttributeSpecification) CompilerOption {
	return func(config *CompilerConfig) *CompilerConfig {
		config.AttrSpecs = attrSpecs
		config.TranslateKeysFromSpec = true
		return config
	}
}
