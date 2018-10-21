package error

import "log"

const (
	WARNING = "[WARNING]:"
	ERROR   = "[ERROR]:"
	MESSAGE = "[MESSAGE]:"
)

func Cw(err error, msg string) {
	if err != nil {
		log.Println(MESSAGE + msg)
		log.Println(WARNING+"%v", err)
	}
}

func Ce(err error, msg string) {
	if err != nil {
		log.Println(MESSAGE + msg)
		log.Fatalf(ERROR+"%v", err)
	}
}
