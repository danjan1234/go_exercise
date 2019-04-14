package main

import (
    "golang.org/x/tour/wc"
    "strings"
)

func WordCount(s string) map[string]int {
    dic := map[string]int{}
    for _, word := range strings.Fields(s) {
        freq, ok := dic[word]
        if ok {
            dic[word] = freq + 1
        } else {
            dic[word] = 1
        }
    }
    return dic
}

func main() {
    wc.Test(WordCount)
}
