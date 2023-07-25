package beater_test

import (
	"fmt"
	"time"

	beater "github.com/pschou/go_beater"
)

func ExampleNew() {
	start := time.Now()

	b := beater.New(func() {
		fmt.Println(time.Now().Sub(start).Truncate(time.Second / 100))
	}, time.Second)

	time.Sleep(time.Second*4 + time.Second/4)
	b.Stop()
	// Output:
	// 1s
	// 2s
	// 3s
	// 4s
}

func ExampleBeat() {
	start := time.Now()

	b := beater.New(func() {
		fmt.Println(time.Now().Sub(start).Truncate(time.Second / 100))
	}, time.Second)

	// Wait a half second and trigger a new beat
	time.Sleep(time.Second / 2)
	b.Beat()

	time.Sleep(time.Second*4 + time.Second/4)
	b.Stop()
	// Output:
	// 500ms
	// 1.5s
	// 2.5s
	// 3.5s
	// 4.5s
}
