package main

import (
	"library"
	"log/slog"
	"os"
	"time"
)

func main() {
	library.HelloWithoutLogger()
	time.Sleep(100 * time.Millisecond)

	h := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource:   false,
		Level:       slog.LevelInfo,
		ReplaceAttr: nil,
	})

	l := slog.New(h)
	slog.SetDefault(l)

	library.HelloWithoutLogger()
	library.HelloWithLogger(slog.NewLogLogger(h, slog.LevelWarn))
}

/*

https://dusted.codes/creating-a-pretty-console-logger-using-gos-slog-package



var appEnv = os.Getenv("APP_ENV")

func main() {

	var appEnv = os.Getenv("APP_ENV")

    opts := &slog.HandlerOptions{
        Level: slog.LevelDebug,
    }

    var handler slog.Handler = slog.NewTextHandler(os.Stdout, opts)
    if appEnv == "production" {
        handler = slog.NewJSONHandler(os.Stdout, opts)
    }

    logger := slog.New(handler)

    logger.Info("Info message")
}

*/
//l.LogAttrs(context.Background(), slog.LevelInfo, "hello world",
//	slog.Int("user_id", 123),
//	slog.String("user_name", "John Doe"),
//)

//l.With(slog.Int("user_id", 123), slog.String("user_name", "John Doe")).Info("hello world")

//}
