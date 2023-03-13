package todo_list

import (
	"github.com/faceair/jio"
)

var RequestBodyUpdateSchema = jio.Object().Optional().Keys(jio.K{
	"title":       jio.String(),
	"description": jio.String(),
	"status":      jio.String(),
})

var RequestBodyCreateSchema = jio.Object().Keys(jio.K{
	"title":       jio.String().Required(),
	"description": jio.String().Default(""),
	"status":      jio.String().Default(""),
})
