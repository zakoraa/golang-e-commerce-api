package logger

import "log"

func LogSuccess(action string, data ...any) {
	log.Printf("✅ SUCCESS: %s", action)
	if len(data) > 0 {
		log.Printf("→ Data: %+v", data[0])
	}
}

func LogError(action string, err error, meta ...map[string]any) {
	log.Printf("❌ ERROR: %s", action)
	log.Printf("→ Message: %s", err.Error())

	if len(meta) > 0 {
		for k, v := range meta[0] {
			log.Printf("→ %s: %+v", k, v)
		}
	}
}
