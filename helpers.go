package rfip

import "strconv"

func PanicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func calculateArgLength(
	contentLength int,
	funcNameLength int,
) (length int, err Error) {
	length = contentLength - funcNameLength
	if length < 0 {
		err = ErrInvalidContentLength
	}
	return
}

func intWithSpaceAsBytes(i int) []byte {
	return []byte(strconv.Itoa(i) + " ")
}
