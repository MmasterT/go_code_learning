package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"strings"
	"sync"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)
var col string

func main() {
	//excercise_1() //easy was not done
	//excercise_2()
	//excercise_3() //done in the web
	//excercise_4()
	//excercise_5() //done in the web
	//excercise_6() //todo 6-9
	//excercise_10() //excercise 10 and 11 together
	//excercise_12()
}

func excercise_2() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style= 'strke-width= 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j) //corner returns two values (x,y) coordinates
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' stroke= '%v' fill = 'white'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, col)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	// Compute surface height z.
	z := f(x, y)
	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	if math.IsInf(r, 0) || math.IsNaN(r) {
		return 0
	}
	if r > 12 {
		col = "red"
	} else {
		col = "blue"
	}
	return math.Sin(r) / r
}

var mu sync.Mutex
var count int

func excercise_4() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		io.WriteString(w, plot3d())
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// counter echoes the number of calls so far.
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
func plot3d() string {
	var print_out string
	print_out = fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>\n", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j) //corner returns two values (x,y) coordinates
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			print_out = print_out + fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	print_out = print_out + fmt.Sprint("</svg>")
	return print_out
}

/*
func excercise_6() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func Map[T, V any](ts []T, fn func(T) V) []V {
	result := make([]V, len(ts))
	for i, t := range ts {
		result[i] = fn(t)
	}
	return result
}


func mandelbrot(z complex128) color.Color {
	const (
		iterations = 200
		contrast   = 4
	)

	var v complex128
	x := real(z)
	y := float64(imag(z))
	supersampling := []color.Color{}

	for i := 0.0; i < 2; i++ {
		x = x + i/2
		for j := 0.0; j < 2; j++ {
			y = y + i/2
			for n := uint8(0); n < iterations; n++ {
				v = v*v + z
				if cmplx.Abs(v) > 2 {
					supersampling = append(supersampling, color.Gray{255 - contrast*n})
				} else {
					supersampling = append(supersampling, color.Black)
				}
			}
		}
	}
	supersamplingR := Map(supersampling, func(item color.Color) float64 {
		item = uint32(item.R)
		item |= item << 8
	})

	supersamplingG := Map(supersampling, func(item color.Color) float64 {
		r = uint32(c.R)
		r |= r << 8
	})

	supersamplingB := Map(supersampling, func(item color.Color) float64 {
		r = uint32(c.R)
		r |= r << 8
	})
	return color.Black //aca hay que devolver el promedio
}
*/

// comma inserts commas in a non-negative decimal integer string.
func excercise_10() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")

	fmt.Println(comma_1(text))
}

func comma_1(s string) string {
	var buf bytes.Buffer

	if strings.Contains(s, ".") {
		index := strings.Index(s, ".")
		s = s[0:index]
	}

	n := len(s)

	if n <= 3 {
		return s
	}
	for i := 0; i < n; i++ {

		if (n-i)%3 == 0 {
			buf.WriteString(",")
		}
		fmt.Fprintf(&buf, "%v", string(s[i]))
	}
	return buf.String()
}

func excercise_12() {
	s1 := os.Args[1]
	s2 := os.Args[2]
	fmt.Printf("%s , %s\n", s1, s2)
	if anagram(s1, s2) {
		fmt.Println("You are in the presence of an anagram!")
	} else {
		fmt.Println("no anagram but nice dick, bro!")
	}

}

func anagram(s1 string, s2 string) bool {
	s1 = strings.TrimSpace(s1)
	s2 = strings.TrimSpace(s2)

	if len(s1) != len(s2) {
		return false
	}
	r1 := []rune(s1)
	r2 := []rune(s2)
	for i, value_r1 := range s1 {
		for j, value_r2 := range s2 {

			if value_r1 == value_r2 {
				fmt.Printf("i = %v\n", i)
				fmt.Printf("j = %v\n", j)
				r1 = delChar(r1, i)
				r2 = delChar(r2, j)
				fmt.Println("deleting a pair of chars...")
			}
			s1 = string(r1)
			s2 = string(r2)
			fmt.Println("searching anagram...")
			if len(s1) != len(s2) && r1 != nil && r2 != nil {
				fmt.Printf("%s != %s\ndifferent length, aha!\n", s1, s2)
				return false
			}
		}

	}
	return true
}

func delChar(s []rune, index int) []rune {
	if index < len(string(s))-1 && index > 0 {
		return append(s[0:index], s[index+1:]...)

	} else if index == len(string(s))-1 {
		return append(s[0:index])
	} else {
	}
	return s[1:]
}
