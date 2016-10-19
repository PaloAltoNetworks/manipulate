package compiler

import (
	"bytes"

	"github.com/aporeto-inc/manipulate"
)

// CompileParameters returns the string of the current parameter
func CompileParameters(p *manipulate.Parameters) string {

	var buffer bytes.Buffer

	if p.OrderByDesc != "" {
		manipulate.WriteString(&buffer, `ORDER BY `)
		manipulate.WriteString(&buffer, p.OrderByDesc)
		manipulate.WriteString(&buffer, ` DESC `)
	} else if p.OrderByAsc != "" {
		manipulate.WriteString(&buffer, `ORDER BY `)
		manipulate.WriteString(&buffer, p.OrderByAsc)
		manipulate.WriteString(&buffer, ` ASC `)
	}

	if p.UsingTTL {
		manipulate.WriteString(&buffer, `USING TTL `)
	}

	if p.IfNotExists {
		manipulate.WriteString(&buffer, `IF NOT EXISTS `)
	} else if p.IfExists {
		manipulate.WriteString(&buffer, `IF EXISTS `)
	}

	return buffer.String()
}
