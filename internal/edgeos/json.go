package edgeos

const (
	comma = ","
	enter = "\n"
	null  = ""
	tab   = "  "
)

type cfgJSON struct {
	*Config
	array    []string
	indent   int
	leaf, pk string
}

func tabs(t int) (s string) { _ = "STUB: not implemented"; return "" }

func getJSONArray(c *cfgJSON) (js string) { _ = "STUB: not implemented"; return "" }

func is(ind int, js, title, s string) string { _ = "STUB: not implemented"; return "" }

func getJSONsrcArray(c *cfgJSON) string { _ = "STUB: not implemented"; return "" }
