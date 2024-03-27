/**
 * Author: Andrei Mikhailov
 * File: go-lib.go
 */

package lib

import (
	"bufio"
	"io"
	"log"
	"os"
)

// like unwrap_or_else in Rust
func UnwrapResultOrElse[T interface{}](result T, problem error) func(func(error) T) T {
	if problem == nil {
		return func(f func(error) T) T { return result }
	} else {
		return func(f func(error) T) T { return f(problem) }
	}
}

// like unwrap in Rust
func UnwrapResult[T interface{}](result T, problem error) T {
	if problem != nil {
		log.Panic(problem)
	}
	return result
}

type OldLine = string
type NewLine = string

// replace in file
// `backupName` is required
func ReplaceInFile[State interface{ skip() bool }](file string, transformer func(OldLine, State) (NewLine, State), initialState State, backupFile string) {
	UnwrapResult(CopyFile(file, backupFile))
	readFile := UnwrapResult(os.Open(backupFile))
	defer readFile.Close()
	newFile := UnwrapResult(os.Create(file))
	w := bufio.NewWriter(newFile)
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var state = initialState
	for fileScanner.Scan() {
		line := fileScanner.Text()
		newLine, newState := transformer(line, state)
		state = newState
		if state.skip() {
		} else {
			w.WriteString(newLine + "\n")
		}
	}
	w.Flush()
}

// copy file
func CopyFile(src, dst string) (int64, error) {
	source := UnwrapResult(os.Open(src))
	defer source.Close()
	destination := UnwrapResult(os.Create(dst))
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}
