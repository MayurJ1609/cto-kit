package logging

import "runtime"

func getFileAndLine() (file string, line int, ok bool) {
	_, file, line, ok = runtime.Caller(3)
	return file, line, ok
}

func getMethod() string {
	fPcs := make([]uintptr, 1)

	n := runtime.Callers(3, fPcs)
	if n == 0 {
		return ""
	}

	methodObj := runtime.FuncForPC(fPcs[0])
	if methodObj == nil {
		return "n/a"
	}

	return methodObj.Name()
}
