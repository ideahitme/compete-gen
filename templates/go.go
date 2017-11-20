package templates

const GO = `package main

import (
    "bufio"
    "os"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Buffer(make([]byte, 2048*2048), 2048*2048)
    scanner.Split(bufio.ScanWords)
}
`
