package main

import (
    "io"
    "os"
    "strings"
)

type rot13Reader struct {
    r io.Reader
}

// Implement io.Reader interface
func (rot13r *rot13Reader) Read(bytes []byte) (int, error) {
    n, err := rot13r.r.Read(bytes)
    for i := 0; i < n; i++ {
        if bytes[i] == ' ' {
            continue
        } else if bytes[i] < 'N' {
            bytes[i] = bytes[i] + 13
        } else {
            bytes[i] = bytes[i] - 13
        }
    }
    return n, err
}

func main() {
    s := strings.NewReader("LBH PENPXRQ GUR PBQR")
    r := rot13Reader{s}

    // Note *r is of type io.Reader since it implements io.Reader interface
    var r1 io.Reader
    r1 = &r
    io.Copy(os.Stdout, r1)
}
