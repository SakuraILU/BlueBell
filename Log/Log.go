package log

import (
	"io"
	"log"
	"os"
	"sync"
)

var (
	infolog *log.Logger
	warnlog *log.Logger
	errlog  *log.Logger
	logs    []*log.Logger
)

var (
	Info   = infolog.Println
	Infof  = infolog.Printf
	Warn   = warnlog.Println
	Warnf  = warnlog.Printf
	Error  = errlog.Println
	Errorf = errlog.Printf
	Fatal  = errlog.Fatal
	Fatalf = errlog.Fatalf
	llk    = sync.Mutex{}
)

var (
	blue   string = "\033[34m"
	orange string = "\033[33m"
	red    string = "\033[31m"
	none   string = "\033[0m"
)

func init() {
	fd := os.Stdout
	flag := log.Ldate | log.Lshortfile
	infolog = log.New(fd, setColor("[INFO ] ", blue), flag)
	warnlog = log.New(fd, setColor("[WARN ] ", orange), flag)
	errlog = log.New(fd, setColor("[ERROR] ", red), flag)
	logs = []*log.Logger{infolog, warnlog, errlog}

	Info = infolog.Println
	Infof = infolog.Printf
	Warn = warnlog.Println
	Warnf = warnlog.Printf
	Error = errlog.Println
	Errorf = errlog.Printf
}

func setColor(str string, color string) string {
	return color + str + none
}

const (
	INFO int = iota
	WARN
	ERROR
	DISABLE
)

func SetLogLevel(level int) {
	llk.Lock()
	defer llk.Unlock()

	for _, log := range logs {
		log.SetOutput(os.Stdout)
	}

	if level > INFO {
		infolog.SetOutput(io.Discard)
	}

	if level > WARN {
		warnlog.SetOutput(io.Discard)
	}

	if level > ERROR {
		errlog.SetOutput(io.Discard)
	}
}

func Panic(str string, v ...any) {
	Errorf(str, v)
	panic("")
}
