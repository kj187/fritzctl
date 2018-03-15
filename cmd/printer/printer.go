package printer

import (
	"encoding/json"
	"io"

	"github.com/bpicode/fritzctl/console"
)

type Printer interface {
	Print(data interface{}, writer io.Writer) error
}

func Print(data interface{}, writer io.Writer) {
	if table, ok := data.(*console.Table); ok {
		table.Print(writer)
		return
	}
	encoder := json.NewEncoder(writer)
	encoder.SetIndent("", "  ")
	encoder.Encode(data)
}
