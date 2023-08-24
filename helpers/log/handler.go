package log

import (
	"context"
	"fmt"
	"grpc-layout/helpers/utils"
	"os"
	"runtime"

	"github.com/gookit/color"
	"golang.org/x/exp/slog"
)

const (
	LevelTrace = slog.Level(-8)
	LevelDebug = slog.Level(-4)
	LevelInfo  = slog.Level(0)
	LevelWarn  = slog.Level(4)
	LevelError = slog.Level(8)
	LevelPanic = slog.Level(12)
	LevelFatal = slog.Level(16)
)

type PrettyHandler struct {
	slog.Handler
	Color bool
	Attrs []slog.Attr
	Group slog.Attr
}

func Trace(msg string, args ...slog.Attr) {
	slog.LogAttrs(context.TODO(), LevelTrace, msg, args...)
}

func Debug(msg string, args ...any) {
	slog.Debug(msg, args...)
}

func Info(msg string, args ...any) {
	slog.Info(msg, args...)
}

func Warn(msg string, args ...any) {
	slog.Warn(msg, args...)
}

func Error(msg string, args ...any) {
	slog.Error(msg, args...)
}
func Panic(msg string, args ...slog.Attr) {
	slog.LogAttrs(context.TODO(), LevelPanic, msg, args...)
	panic(msg)
}

func Fatal(msg string, args ...slog.Attr) {
	slog.LogAttrs(context.TODO(), LevelFatal, msg, args...)
	os.Exit(1)
}

func (h *PrettyHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.Handler.Enabled(ctx, level)
}

func (h *PrettyHandler) Handle(ctx context.Context, r slog.Record) (err error) {
	pc, file, line, _ := runtime.Caller(4)
	r.PC = pc + 1

	if err = h.Handler.Handle(ctx, r); err != nil {
		return err
	}

	level := r.Level.String()
	message := r.Message
	prefix := r.Time.Format("2006-01-02 15:05:05.000")
	prefix = color.HEX("#A9B7C6").Sprint(prefix)
	switch r.Level {
	case LevelTrace:
		level = color.Hex("#7970A9").Sprint("TRACE")
		message = color.Hex("#7970A9").Sprint(message)
	case slog.LevelDebug:
		level = color.Hex("#808080").Sprint(level)
		message = color.Hex("#808080").Sprint(message)
	case slog.LevelInfo:
		level = color.Green.Sprint(level + " ")
		message = color.Green.Sprint(message)
	case slog.LevelWarn:
		level = color.Yellow.Sprint(level + " ")
		message = color.Yellow.Sprint(message)
	case slog.LevelError:
		level = color.Hex("#ff3800").Sprint(level)
		message = color.Hex("#ff3800").Sprint(message)
	case LevelPanic:
		level = color.Hex("#F998CC").Sprint("PANIC")
		message = color.Hex("#F998CC").Sprint(message)
	case LevelFatal:
		level = color.Hex("#FE4EDA").Sprint("FATAL")
		message = color.Hex("#FE4EDA").Sprint(message)
	}

	var attrs []slog.Attr
	r.Attrs(func(a slog.Attr) bool {
		attrs = append(attrs, a)
		return true
	})

	if h.Group.Key != "" {
		h.AddGroupAttr(attrs...)
		attrs = []slog.Attr{h.Group}
	}

	s := AttrString(append(h.Attrs, attrs...)...)
	s = color.Cyan.Sprint(s)

	source := utils.BaseN(fmt.Sprintf("%s:%d", file, line), 3)
	if h.Color {
		fmt.Printf("%s | %s | %s > %s %s\n", prefix, level, source, message, s)
	}
	return
}

func (h *PrettyHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	newHandler := h.Handler.WithAttrs(attrs)
	handler := &PrettyHandler{Handler: newHandler, Attrs: h.Attrs, Group: h.Group}

	if handler.Group.Key != "" {
		handler.AddGroupAttr(attrs...)
	} else {
		handler.Attrs = append(h.Attrs, attrs...)
	}

	return handler
}

func (h *PrettyHandler) WithGroup(name string) slog.Handler {
	newHandler := h.Handler.WithGroup(name)
	handler := &PrettyHandler{Handler: newHandler, Attrs: h.Attrs, Group: h.Group}

	if group := slog.Group(name); h.Group.Key == "" {
		handler.Group = group
	} else {
		handler.AddGroupAttr(group)
	}
	return handler
}

func (h *PrettyHandler) AddGroupAttr(attrs ...slog.Attr) {
	if v := &LastGroup(&h.Group).Value; v.Kind() == slog.KindGroup {
		*v = slog.GroupValue(append(v.Group(), attrs...)...)
	}
}
