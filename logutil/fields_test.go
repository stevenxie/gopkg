package logutil_test

import (
	"os"

	"github.com/sirupsen/logrus"
	"go.stevenxie.me/gopkg/logutil"
)

type SomeStruct struct {
	log *logrus.Entry
}

func SomeFunc() {}

func ExampleWithType() {
	l := logrus.New()
	l.SetOutput(os.Stdout)
	l.SetFormatter(_logrusFormatter)
	log := logrus.NewEntry(l)

	log1 := logutil.WithType(log, (*SomeStruct)(nil))
	log1.Info("Added 'SomeStruct' component.")

	log2 := logutil.WithType(log1, SomeFunc)
	log2.Info("Added 'SomeFunc' component.")

	log1.Info("First logger component unchanged.")

	// Output:
	// level=info msg="Added 'SomeStruct' component." type=logutil_test.SomeStruct
	// level=info msg="Added 'SomeFunc' component." type=logutil_test.SomeFunc
	// level=info msg="First logger component unchanged." type=logutil_test.SomeStruct
}

func (ss SomeStruct) LogWithMethod() {
	logutil.
		WithMethod(ss.log, SomeStruct.LogWithMethod).
		Info("Hello from method!")
}

func ExampleWithMethod() {
	l := logrus.New()
	l.SetOutput(os.Stdout)
	l.SetFormatter(_logrusFormatter)

	ss := SomeStruct{log: logrus.NewEntry(l)}
	ss.LogWithMethod()

	// Output:
	// level=info msg="Hello from method!" method=LogWithMethod
}
