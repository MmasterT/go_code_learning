package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	//excercise_1()
	//excercise_2()
	//excercise_4()
	//excercise_5(os.Stdout)
	//excercise_7()
	//excercise_8()
	excercise_9()
}

func excercise_1() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args[0:], " "))
	secs := time.Since(start).Seconds()
	fmt.Println(secs)
}

func excercise_2() {
	start := time.Now()
	s, sep := "", ""
	for index, arg := range os.Args[1:] {
		s += sep + arg + " " + strconv.Itoa(index)
		sep = "\n"
	}
	fmt.Println(s)
	secs := time.Since(start).Seconds()
	fmt.Println(secs)
}

//excercise_3 is add the time module and do a benchmarrk of the functions

// #TODO excersie_4
func excercise_4() {
	counts := make(map[string]int)
	files := os.Args[1:]
	filename := ""
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%s\t%s\t%d\n", filename, line, n)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

var palette = []color.Color{color.Black, color.RGBA{0x00, 0x99, 0x00, 0xff}}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func excercise_5(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 128   // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.05
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

// #TODO excersie_6

func excercise_7() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}

func excercise_8() {
	for _, url := range os.Args[1:] {
		prefix := "http://"
		if strings.HasPrefix(url, prefix) == false {
			url = strings.Join([]string{prefix, url}, "")
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}

func excercise_9() {
	for _, url := range os.Args[1:] {
		prefix := "http://"
		if strings.HasPrefix(url, prefix) == false {
			url = strings.Join([]string{prefix, url}, "")
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		status := resp.Status
		fmt.Printf("%s\n%s\n", b, status)
	}
}

//TODO excercise 10
