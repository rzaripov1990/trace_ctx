package trace_ctx

import "strings"

var (
	rp = strings.NewReplacer("-", "")
)

func replace(s string) string {
	return rp.Replace(s)
}
