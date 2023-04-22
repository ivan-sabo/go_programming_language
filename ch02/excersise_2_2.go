// This package converts pounds to kilograms and vice versa.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

const poundToKgRate = 0.4536

type kilogram float32
type pound float32

func (kg kilogram) String() string { return fmt.Sprintf("%.4f kg\n", kg) }
func (lbs pound) String() string   { return fmt.Sprintf("%.4f lbs\n", lbs) }

func main() {
	flag.Parse()
	inputs := flag.Args()

	if len(inputs) == 0 {
		input := bufio.NewScanner(os.Stdin)

		for input.Scan() {
			value, err := strconv.Atoi(input.Text())
			if err != nil {
				log.Fatalf("error: %v", err)
				continue
			}

			kgToLbs(kilogram(value))
			lbsTokg(pound(value))
		}
		return
	}

	for _, v := range inputs {
		value, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalf("error: %v", err)
			continue
		}

		kgToLbs(kilogram(value))
		lbsTokg(pound(value))
	}
}

func kgToLbs(kg kilogram) {
	lbs := pound(kg / poundToKgRate)
	fmt.Print(lbs.String())
}

func lbsTokg(lbs pound) {
	kg := kilogram(lbs * poundToKgRate)
	fmt.Print(kg.String())
}

func init() {
	fmt.Println("this is an init function!")
}

func init() {
	fmt.Println("this is second init function!")
}
