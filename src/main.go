package main

import (
	"fmt"
	"log/slog"
	"os"
	"pedeveaux/go-dev/greetings"
)

var LogLevel = new(slog.LevelVar)

func main() {
	// setting the log level
	LogLevel.Set(slog.LevelDebug - 1)

	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: LogLevel,
	})
	slog.SetDefault(slog.New(handler))

	// logger := slog.Default()
	message, err := greetings.Hello("")
	if err != nil {
		slog.Error(err.Error())
	}
	fmt.Println(message)
}
