package command

import (
	"strings"
	"Minireddis/storage"
)

func HandleCommand(line string, store *storage.Store) string {
	// 🔥 Clean input (important for Windows terminal)
	line = strings.TrimSpace(line)
	if line == "" {
		return ""
	}

	// Split by spaces (not limited)
	parts := strings.Fields(line)
	if len(parts) == 0 {
		return ""
	}

	cmd := strings.ToUpper(parts[0])

	switch cmd {

	case "SET":
		if len(parts) < 3 {
			return "ERR SET usage: SET key value"
		}
		key := parts[1]

		// 🔥 JOIN EVERYTHING AFTER KEY AS VALUE
		value := strings.Join(parts[2:], " ")

		store.Set(key, value)
		return "OK"

	case "GET":
		if len(parts) != 2 {
			return "ERR GET usage: GET key"
		}
		val, ok := store.Get(parts[1])
		if !ok {
			return "nil"
		}
		return val

	default:
		return "ERR unknown command"
	}
}
