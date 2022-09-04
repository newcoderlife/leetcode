package logger_rate_limiter

type Logger struct {
	buffer map[string]int
}

func Constructor() Logger {
	return Logger{buffer: make(map[string]int)}
}

func (this *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	if ts, ok := this.buffer[message]; ok {
		if timestamp-ts <= 10 {
			return false
		}
	}

	this.buffer[message] = timestamp
	return true
}
