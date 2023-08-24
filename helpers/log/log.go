package log

import (
	"errors"
	"grpc-layout/helpers/utils"
	"io"
	"os"
	"strings"
	"sync"
	"time"

	"golang.org/x/exp/slog"
)

type LogDir struct {
	Dir    string
	Format string
	lock   bool
	open   string
	file   *os.File
	mu     sync.Mutex
}

func NewLogDir(dir string, lock bool, perm ...os.FileMode) *LogDir {
	_perm := append(perm, 0)[0]
	if err := os.Mkdir(dir, _perm); err != nil && !errors.Is(err, os.ErrExist) {
		panic(err)
	}

	return &LogDir{Dir: dir, Format: "20060102", lock: lock}
}

func (d *LogDir) Write(b []byte) (n int, err error) {
	format := time.Now().Format(d.Format)
	fileName := strings.TrimRight(d.Dir, "/") + "/day_" + format + ".log"
	if d.open != fileName {
		if d.file != nil {
			_ = d.file.Close()
		}
		d.open = fileName
		d.file, err = os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	}

	if err == nil {
		if d.lock {
			d.mu.Lock()
			defer d.mu.Unlock()
		}
		n, err = d.file.Write(b)
	}

	return
}

var LevelNames = map[slog.Leveler]string{
	LevelTrace: "TRACE",
	LevelPanic: "PANIC",
	LevelFatal: "FATAL",
}

func SetOutput(w io.Writer, level slog.Leveler, color bool) {
	opts := slog.HandlerOptions{
		AddSource: true,
		Level:     level,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.SourceKey {
				a.Value = slog.StringValue(utils.BaseN(a.Value.String()[:len(a.Value.String())-1], 3))
			} else if a.Key == slog.TimeKey {
				value := a.Value.Time().Format("2006-01-02 15:04:05.000")
				a.Value = slog.StringValue(value)
			} else if a.Key == slog.LevelKey {
				level := a.Value.Any().(slog.Level)
				levelLabel, exists := LevelNames[level]
				if !exists {
					levelLabel = level.String()
				}
				a.Value = slog.StringValue(levelLabel)
			}
			return a
		},
	}
	h := &PrettyHandler{Handler: slog.NewJSONHandler(w, &opts), Color: color}
	slog.SetDefault(slog.New(h))
}
