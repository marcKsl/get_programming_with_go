package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type safeWriter struct {
	// We can pass any type to w that implements the io.Writer interface
	// For example os.File does implement the io.Writer interface, because
	// it contains a method Write(b []byte) n int, err error, which implicitly
	// implements the io.Writer interface
	w   io.Writer
	err error
}

// We use a pointertype here because we want to actually modify
// the value of the given safeWriter.err value
func (sw *safeWriter) writeln(s string) {
	if sw.err != nil {
		return
	}
	_, sw.err = fmt.Fprintln(sw.w, s)
}

func proverbs(name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()

	sw := safeWriter{w: f}
	sw.writeln("Errors are values.")
	sw.writeln("Don't just check errors, handle them gracefully.")
	sw.writeln("Don't panic.")
	sw.writeln("Make the zero value useful.")
	sw.writeln("The bigger the interface, the weaker the abstraction.")
	sw.writeln("interface{} says nothing.")
	sw.writeln("Gofmt's style is no one's favorite, yet gofmt is everyone'sfavorite.")
	sw.writeln("Documentation is for users.")
	sw.writeln("A little copying is better than a little dependency.")
	sw.writeln("Clear is better than clever.")
	sw.writeln("Concurrency is not parallelism.")
	sw.writeln("Don't communicate by sharing memory, share memory bycommunicating.")
	sw.writeln("Channels orchestrate; mutexes serialize.")

	return sw.err
}

func main() {

	files, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}

	err = proverbs("proverbs.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	files, err = ioutil.ReadDir(".")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}
