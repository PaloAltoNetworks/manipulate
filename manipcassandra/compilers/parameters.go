package compilers

import (
	"bytes"

	"github.com/aporeto-inc/manipulate"
)

// CompileParameters returns the string of the current parameter
func CompileParameters(p *manipulate.Parameters) string {

	var buffer bytes.Buffer

	if p.OrderByDesc != "" {
		buffer.WriteString(`ORDER BY `)
		buffer.WriteString(p.OrderByDesc)
		buffer.WriteString(` DESC `)
	} else if p.OrderByAsc != "" {
		buffer.WriteString(`ORDER BY `)
		buffer.WriteString(p.OrderByAsc)
		buffer.WriteString(` ASC `)
	}

	if p.UsingTTL {
		buffer.WriteString(`USING TTL `)
	}

	if p.IfNotExists {
		buffer.WriteString(`IF NOT EXISTS `)
	} else if p.IfExists {
		buffer.WriteString(`IF EXISTS `)
	}

	return buffer.String()
}
