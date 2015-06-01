package main

import (
    "fmt";
    "io"
    "xml"
    "strings"
)

// https://gist.github.com/chrisfarms/1377218
const XML = `<?xml version="1.0" encoding="UTF-8" ?>
<langs>
<en>English</en>
</langs>`

func main() {
    r := strings.NewReader(XML)
    m := xmlToMap(r)
    fmt.Println(m)
}

func xmlToMap(r io.Reader) map[string]string {
    // result
    m := make(map[string]string)
    // the current value stack
    values := make([]string,0)
    // parser
    p := xml.NewParser(r)
    for token, err := p.Token(); err == nil; token, err = p.Token() {
        switch t := token.(type) {
        case xml.CharData:
            // push
            values = append(values, string([]byte(t)))
        case xml.EndElement:            
            if t.Name.Local == "langs" {
                continue
            }
            m[t.Name.Local] = values[len(values)-1]
            // pop
            values = values[:len(values)]
        }
    }
    // done
    return m
}