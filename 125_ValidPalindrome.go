package main

import (
    "fmt"
    "unicode"
)

func isPalindrome(s string) bool {
    runes := []rune(s)
    l, r := 0, len(runes)-1

    for l < r {
        for l < r && !isAlnum(runes[l]) {
            l++
        }
        for l < r && !isAlnum(runes[r]) {
            r--
        }
        if toLower(runes[l]) != toLower(runes[r]) {
            return false
        }
        l++
        r--
    }
    return true
}

func isAlnum(r rune) bool {
    return unicode.IsLetter(r) || unicode.IsDigit(r)
}

func toLower(r rune) rune {
    if 'A' <= r && r <= 'Z' {
        return r + ('a' - 'A')
    }
    return r
}

func main() {
    fmt.Println(isPalindrome("A man, a plan, a canal: Panama")) // true
    fmt.Println(isPalindrome("race a car"))                     // false
}