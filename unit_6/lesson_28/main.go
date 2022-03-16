package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/url"
	"os"
	"strings"
)

func proverbs(name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = fmt.Fprintln(f, "Errors are values.")
	if err != nil {
		// f.Close() // Unnecessary due to `defer` statement
		return err
	}

	_, err = fmt.Fprintln(f, "Don't just check errors, handle them gracefully")
	// f.Close() // Unnecessary due to `defer` statement

	sw := safeWriter{w: f}
	sw.writeln("Don't panic.")
	sw.writeln("Make the zero value useful.")
	sw.writeln("The bigger the interface, the weaker the abstraction")
	sw.writeln("interface{} says nothing.")
	sw.writeln("gofmt's stype is no one's favourite, yet gofmt is everyone's favourite")
	sw.writeln("Documentation is for users")
	sw.writeln("A little copying is better than a little dependency")
	sw.writeln("Clear is better than clever")
	sw.writeln("Concurrency is not parallelism")
	sw.writeln("Don't communicate by sharing memory, share memory by communicating")
	sw.writeln("Channels orchestrate, mutexes serialize")
	return sw.err
}

type safeWriter struct {
	w   io.Writer
	err error
}

func (sw *safeWriter) writeln(s string) {
	if sw.err != nil {
		return // Don't write if the last time we tried something it err'd
	}

	_, sw.err = fmt.Fprintln(sw.w, s) //Write a line and store the error in the struct
}

const rows, columns = 9, 9

type Grid [rows][columns]int8

func (g *Grid) Set(row, column int, digit int8) error {
	var errs SudokuError
	if !inBounds(row, column) {
		errs = append(errs, ErrBounds)
	}
	if !validDigit(digit) {
		errs = append(errs, ErrDigit)
	}
	if len(errs) > 0 {
		return errs
	}
	g[row][column] = digit
	return nil

}

func inBounds(row, column int) bool {
	if row < 0 || row >= rows {
		return false
	}
	if column < 0 || column >= columns {
		return false
	}
	return true
}
func validDigit(digit int8) bool {
	if digit >= 1 && digit <= 9 {
		return true
	}
	return false
}

var (
	ErrBounds = errors.New("Out of bounds")
	ErrDigit  = errors.New("Invalid digit")
)

func (g *Grid) Error() string {
	return "Grid Error"
}

type SudokuError []error

func (se SudokuError) Error() string {
	var s []string
	for _, err := range se {
		s = append(s, err.Error())
	}
	return strings.Join(s, ", ")
}

func main() {
	files, err := ioutil.ReadDir("..")
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

	var g Grid
	err = g.Set(10, 0, 15)
	if err != nil {
		if errs, ok := err.(SudokuError); ok {
			fmt.Println(len(errs))
			for _, e := range errs {
				fmt.Println(e)
			}
		}
		switch err {
		case ErrBounds, ErrDigit: // errors.New uses a memory address so it's comparing addresses and not the text contained
			fmt.Printf("An error occured: %v. \n", err)
			os.Exit(1)
		default:
			fmt.Println(err)
		}
	}

	// url.go
	parse, err := url.Parse("https://w a.com")
	if err != nil {
		fmt.Printf("%#v\n", err)
		if errs, ok := err.(*url.Error); ok {
			fmt.Println("Op", errs.Op)
			fmt.Println("URL", errs.URL)
			fmt.Println("Err", errs.Err)
		}
	}
	fmt.Println(parse)
}
