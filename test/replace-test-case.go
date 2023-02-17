package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	chars := map[rune]string{'<': "推し",
		'>': "尊い",
		'+': "待って",
		'-': "しんどい",
		'.': "マジ",
		',': "最高",
		'[': "尊すぎ",
		']': "無理",
	}

	program := ""

	if len(os.Args) < 2 {
		log.Fatalln("Please give the input Brainfuxk file.")
	}

	program_bytes, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalln("The input Brainfuxk file does not exist.")
	}
	program = string(program_bytes)

	for k, e := range chars {
		program = strings.Replace(program, string([]rune{k}), e, -1)
	}

	f, err := os.Create(os.Args[1])
	if err != nil {
		log.Fatalln("The input Brainfuxk file does not exist.")
	}
	f.WriteString(program)
}
