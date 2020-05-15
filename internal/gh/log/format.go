package log

import (
	"bytes"
	"github.com/sirupsen/logrus"
)

type Formatter struct {}

func (f *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	out := bytes.Buffer{}

}

