package todo_item

import (
	"github.com/faceair/jio"
)

var RequestBodyUpdateSchema = jio.Object().Optional().Keys(jio.K{
	"list_id":     jio.String(),
	"title":       jio.String(),
	"description": jio.String(),
	"status":      jio.String(),
})

var RequestBodyCreateSchema = jio.Object().Keys(jio.K{
	"list_id":     jio.String().Required(),
	"title":       jio.String().Required(),
	"description": jio.String().Default(""),
	"status":      jio.String().Default(""),
})
