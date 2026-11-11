package helpers

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func InputInt(reader *bufio.Reader, out io.Writer, text string) (int, error) {
	fmt.Fprint(out, text)

	s, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(strings.TrimSpace(s))
}
