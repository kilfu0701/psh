package langs

var X = map[string]interface{}{
	"go": {
		KEYWORD: "break default func interface select case map struct chan else goto package switch " +
			"const fallthrough if range type continue for import return var go defer " +
			"bool byte complex64 complex128 float32 float64 int8 int16 int32 int64 string uint8 " +
			"uint16 uint32 uint64 int uint uintptr rune",
		LITERAL: "true false iota nil",
		BUILTIN: "append cap close complex copy imag len make new panic print println real recover delete",
	},
}


func Comment(begin string, end string) map[string]interface{} {
	mode :=  map[string]interface{}{
		"className": "comment",
		"begin": begin,
		"end": end,
		"contains": []interface{}{},
	}

	mode["contains"] = append(mode["contains"].([]string), map[string]interface{}{
		"className": "doctag",
		"begin": "(?:TODO|FIXME|NOTE|BUG|XXX):",
		"relevance": 0,
	})

	return mode
}
