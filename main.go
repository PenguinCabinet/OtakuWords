package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

/*
<,推し
>,尊い
+,待って
-,しんどい
.,マジ
,,最高
[,尊すぎ
],無理
*/
var chars = map[rune]string{
	'<': "推し",
	'>': "尊い",
	'+': "待って",
	'-': "しんどい",
	'.': "マジ",
	',': "最高",
	'[': "尊すぎ",
	']': "無理",
}

func program_split(program []rune) []string {
	var A []string
	var stack_temp string
	for len(program) > 0 {
		check_match := false
		for _, e := range chars {
			if len(program) >= len([]rune(e)) {
				if string(program[len(program)-len([]rune(e)):]) == e {
					if stack_temp != "" {
						A = append(A, stack_temp)
					}
					A = append(A, string(program[len(program)-len([]rune(e)):]))
					program = program[:len(program)-len([]rune(e))]
					check_match = true
					stack_temp = ""
				}
			}
		}
		if !check_match {
			stack_temp += string([]rune{program[len(program)-1]})
			program = program[:len(program)-1]
		}
	}
	if stack_temp != "" {
		A = append(A, stack_temp)
	}
	for i := 0; i < len(A)/2; i++ {
		A[i], A[len(A)-i-1] = A[len(A)-i-1], A[i]
	}
	return A
}

func program_check(program []string) []string {
	var A []string
	A_index := 0
	nest := 0
	for _, e := range program {
		if e == chars['<'] ||
			e == chars['>'] ||
			e == chars['+'] ||
			e == chars['-'] ||
			e == chars['.'] ||
			e == chars[','] ||
			e == chars['['] ||
			e == chars[']'] {
			if e == chars['['] {
				nest++
			}
			if e == chars[']'] {
				nest--
			}

			A = append(A, e)
			A_index++
		}
		/*else{
		    writefln("ERROR:There is an illegal character in the source code.")
		    exit(-1)
		}*/
		if nest < 0 {
			log.Fatalln("ERROR:Incorrect block nesting.")
			os.Exit(-1)
		}
	}

	if nest != 0 {
		log.Fatalln("ERROR:Incorrect block nesting.")
	}

	return A
}

func getchar() rune {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	return []rune(string([]byte(input)[0]))[0]
}

func program_run(program []string) {

	var mem [1024]byte
	mem_count := 0

	for program_count := 0; program_count < len(program); program_count++ {
		switch program[program_count] {
		case chars['>']:
			mem_count++
		case chars['<']:
			mem_count--
		case chars['+']:
			mem[mem_count]++
		case chars['-']:
			mem[mem_count]--
		case chars['.']:
			fmt.Printf("%c", rune(mem[mem_count]))
		case chars[',']:
			temp := int(getchar())
			mem[mem_count] = byte(temp)
		case chars['[']:
			if mem[mem_count] == 0 {
				nest := 0
				for true {
					program_count++
					if program[program_count] == chars['['] {
						nest++
					}
					if program[program_count] == chars[']'] && nest == 0 {
						break
					}
					if program[program_count] == chars[']'] {
						nest--
					}
				}
			}
		case chars[']']:
			if mem[mem_count] != 0 {
				nest := 0
				for true {
					program_count--
					if program[program_count] == chars[']'] {
						nest++
					}
					if program[program_count] == chars['['] && nest == 0 {
						break
					}
					if program[program_count] == chars['['] {
						nest--
					}
				}
			}
		}
	}
}

func interpreter_main(args []string) {

	if len(args) < 2 {
		log.Fatalln("Please give the input Brainfuxk file.")
	}
	program := ""
	var splited_program []string
	program_bytes, err := ioutil.ReadFile(args[1])
	if err != nil {
		log.Fatalln("The input Brainfuxk file does not exist.")
	}
	program = string(program_bytes)

	splited_program = program_split([]rune(program))

	splited_program = program_check(splited_program)

	program_run(splited_program)
}

func main() {

	interpreter_main(os.Args)
}
