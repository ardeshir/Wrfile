package main

import (
    "io/ioutil"
    "os"
    )
    
func main() {
    str := "How to test write to file"
    loc := "test.txt"
    err :=  Wrfile(str, loc)
    
    if err != nil {
        panic(err)
    }
}

func Wrfile(str, loc string) error {
    err := writeFile(loc,[]byte(str), os.ModeAppend)
    if err != nil {
        return err
    }
    return nil
}
var writeFile = ioutil.WriteFile 