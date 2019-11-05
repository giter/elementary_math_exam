package main

import (
	"fmt"
	"math/rand"
	"time"
)

func E100Plus() string {

	for {

		a := rand.Intn(100)
		b := rand.Intn(100)

		if a+b < 100 {
			return fmt.Sprintf("%d + %d = ", a, b)
		}
	}
}

func E100Minus() string {

	for {

		a := rand.Intn(100)
		b := rand.Intn(100)

		if a > b {
			return fmt.Sprintf("%d - %d = ", a, b)
		}
	}
}

func E100Mul() string {

	for {

		a := rand.Intn(100)
		b := rand.Intn(100)

		if a > 1 && b > 1 && a*b < 100 && a != 1 && b != 1 {
			return fmt.Sprintf("%d x %d = ", a, b)
		}
	}
}

func E100Div() string {

	for {

		a := rand.Intn(81)
		b := rand.Intn(10)

		if b > 1 && a%b == 0 && a > 1 {
			return fmt.Sprintf("%d ÷ %d = ", a, b)
		}
	}
}

func E1004Z1() string {

	for {

		a := rand.Intn(100)
		b := rand.Intn(100)
		c := rand.Intn(100)

		if a+b*c < 100 && b > 1 && c > 1 {

			return fmt.Sprintf("%d + %d x %d = ", a, b, c)
		}
	}
}

func E1004Z2() string {

	for {

		a := rand.Intn(100)
		b := rand.Intn(100)
		c := rand.Intn(100)

		if c > 1 && b%c == 0 && a+b/c < 100 && b > 0 {

			return fmt.Sprintf("%d + %d ÷ %d = ", a, b, c)
		}
	}
}

func E1004Z3() string {

	for {

		a := rand.Intn(100)
		b := rand.Intn(100)
		c := rand.Intn(100)

		if c > 1 && (a+b)%c == 0 && a > 1 && b > 1 {

			return fmt.Sprintf("(%d + %d) ÷ %d = ", a, b, c)
		}
	}
}

func E1004Z4() string {

	for {

		a := rand.Intn(100)
		b := rand.Intn(100)
		c := rand.Intn(100)

		if c > 1 && a-b > 0 && (a-b)%c == 0 && a > 1 && b > 1 {

			return fmt.Sprintf("(%d - %d) ÷ %d = ", a, b, c)
		}
	}
}

func E1004Z5() string {

	for {

		a := rand.Intn(100)
		b := rand.Intn(100)
		c := rand.Intn(100)

		if (a+b)*c < 100 && c > 1 && a > 1 && b > 1 {

			return fmt.Sprintf("(%d + %d) x %d = ", a, b, c)
		}
	}
}

func E1004Z6() string {

	for {

		a := rand.Intn(100)
		b := rand.Intn(100)
		c := rand.Intn(100)

		if (a-b)*c < 100 && (a-b)*c > 0 && a > 1 && b > 1 {

			return fmt.Sprintf("(%d - %d) x %d = ", a, b, c)
		}
	}
}

func E1004Z7() string {

	for {

		a := rand.Intn(1000)
		b := rand.Intn(1000)
		c := rand.Intn(1000)

		if a-b-c > 0 {
			return fmt.Sprintf("%d - %d - %d = ", a, b, c)
		}
	}
}

func E1004Z8() string {

	for {

		a := rand.Intn(1000)
		b := rand.Intn(1000)
		c := rand.Intn(1000)

		if a+b+c < 1000 {
			return fmt.Sprintf("%d + %d + %d = ", a, b, c)
		}
	}
}

func E1004Z9() string {

	for {

		a := rand.Intn(1000)
		b := rand.Intn(1000)
		c := rand.Intn(1000)

		if a+b-c < 1000 && a+b-c > 0 {
			return fmt.Sprintf("%d + %d - %d = ", a, b, c)
		}
	}
}

func main() {

	rand.Seed(time.Now().Unix())

	fs := []func() string{
		E100Plus, E100Minus, E100Mul, E100Div, E1004Z1, E1004Z2, E1004Z3, E1004Z4, E1004Z5, E1004Z6, E1004Z7, E1004Z8, E1004Z8, E1004Z9,
	}

	for i := 0; i < 100; i++ {
		fmt.Println(fs[rand.Intn(len(fs))]())
	}
}
