/*
   This library was taken from https://github.com/ianberdin/l/blob/master/README.md when the repo was first opened
   it suits my purpose, but I will most likely change it in the future.
*/

package logger

// Subscribe on every log
func (t *Logger) Subscribe(handler func(text string, lvl Level)) {
	t.handlers = append(t.handlers, handler)
}

func (t *Logger) executeHandlers(text string, lvl Level) {
	for _, handler := range t.handlers {
		handler(text, lvl)
	}
}
