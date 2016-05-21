// Author: Alexandre Wilhelm
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package manipcassandra

import "bytes"

// Parameter is a parameter struct which can be used with Cassandra
type Parameter struct {
	IfNotExists bool
	IfExists    bool
	UsingTTL    bool
	OrderByDesc string
	OrderByAsc  string
}

// Compile returns the string of the current parameter
func (p *Parameter) Compile() string {

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
