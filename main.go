package main

import (
    "crypto/sha256"
    "fmt"
)

func hashwork(message []byte, rounds int) []byte {
    var output []byte
    for i:=0;i<rounds;i++{
        o := sha256.Sum256(message)
        //fmt.Printf("%d %x\n", i, o[:])
        message = o[:]
        output = o[:]
    }
    return output
}

func main() {
    d := []byte("Hello World")
    //h := sha256.Sum256(d)
    //fmt.Printf("%x\n",h[:])
    //s := fmt.Sprintf("%x", h)
    //o := sha256.Sum256([]byte(s))
    //fmt.Printf("%x\n", o)
    //fmt.Printf("\n")
    //m := sha256.Sum256(d)
    //fmt.Printf("%x\n",m[:])
    //n := sha256.Sum256(m[:])
    //fmt.Printf("%x\n",n[:])
    g := hashwork(d, 2)
    fmt.Printf("%x\n", g)
    g = hashwork(d, 3)
    fmt.Printf("%x\n", g)
    g = hashwork(d, 1000)
    fmt.Printf("%x\n", g)
    g = hashwork(d, 1000000)
    fmt.Printf("%x\n", g)
    g = hashwork(d, 1000000000)
    fmt.Printf("%x\n", g)
}
