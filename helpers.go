package rfip

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func readContentLength(bufr *bufio.Reader) (contentLength int, err error) {
	bodySizeString, err := bufr.ReadString(' ')
	if err != nil {
		return 0, ErrBadContentLengthRead
	}
	contentLength, err = strconv.Atoi(strings.TrimSpace(bodySizeString))
	if err != nil {
		return 0, ErrBadContentLengthRead
	}
	return
}

func readURL(bufr *bufio.Reader) (length int, url []string, err error) {
	withSpace, err := bufr.ReadString(' ')
	if err != nil {
		err = ErrBadURLRead
		return
	}
	length = len(withSpace)
	url = strings.Split(strings.TrimSpace(withSpace), "/")
	return
}

func calculateArgLength(contentLength int, urlLength int) (n int, err error) {
	n = contentLength - urlLength
	if n < 0 {
		err = ErrInvalidContentLength
	}
	return
}

func readBody(bufr *bufio.Reader, contentLength int) (data []byte, err error) {
	data = make([]byte, contentLength)
	if _, err = io.ReadFull(bufr, data); err != nil {
		err = ErrBadBodyRead
	}
	return
}
