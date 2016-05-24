package main

import(
    "os"
    "errors"
    "testing"
    )
    
func TestWrfile(t *testing.T) {
    testStr := "Test String"
    testLoc := "testfile.txt"
    // store the current WriteFile and reinstate it at the end of test
    oldWriteFile := writeFile
    
    defer func() {
        writeFile = oldWriteFile
    }()
    
    // create new function that tests the passed paramaerters
    writeFile = func(loc string, data []byte, perm os.FileMode) error {
        if loc != testLoc {
            t.Error("Expected loc to be", testLoc, " but got" , loc)
        }
        if string(data) != testStr {
            t.Error("Expected data to be",  testStr, " but got" , string(data))
        }
        return nil
    }
    
    err := Wrfile(testStr, testLoc)
    if err != nil {
        t.Error("Expected no error, but got", err)
    }

}

func TestWrfileError(t *testing.T) {
    testStr := "Test String"
    testLoc := "testfile.txt"
    testErr := errors.New("Ah! An Error.")
    
    oldWriteFile := writeFile
    defer func() {
        writeFile = oldWriteFile
    }()
    
    writeFile = func(loc string, data []byte, perm os.FileMode) error {
        return testErr
    }
    
    err := Wrfile(testStr, testLoc)
    if err != testErr {
        t.Error("Expected an error, but got ", err )
    }
}