package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"./kmeans"
)

/*ModeSync executes the algorithm synchronously*/
const ModeSync = "sync"

/*ModeAsync executes the algorithm asynchronously*/
const ModeAsync = "async"

/*ModeChart executes the algorithm and creates a chart*/
const ModeChart = "chart"

func main() {
	var dataset []kmeans.Point
	rand.Seed(time.Now().UnixNano())

	// t: number of iterations
	var t *int
	t = new(int)
	*t = 0

	k := flag.Int("k", 0, "number of clusters")
	n := flag.Int("n", 0, "number of elements")
	makeGif := flag.Bool("gif", false, "wheter make gif or not")
	static := flag.Bool("static", false, "don't random values")
	mode := flag.String("mode", ModeSync, "Mode to run the program")

	flag.Parse()

	if *static {
		*k = 4
		*n = 800
	}

	// If there is no "k" or "n" the program will exit
	if *k == 0 || *n == 0 {
		flag.Usage()
		os.Exit(1)
	}

	if *static {
		dataset = kmeans.GetStaticPoints()
	} else {
		// Generate some random points
		for i := 0; i < *n; i++ {
			dataset = append(dataset, kmeans.Point{X: rand.Float64(), Y: rand.Float64()})
		}
	}

	// d1 := []byte("")
	// d1 = append(d1, []byte("Point{"+strconv.FormatFloat(rand.Float64(), 'f', 6, 64)+", "+strconv.FormatFloat(rand.Float64(), 'f', 6, 64)+"},\n")...)
	// ioutil.WriteFile("prueba.txt", d1, 0644)
	// os.Exit(1)

	if *mode == ModeSync {
		kmeans.RunSync(dataset, *k, *static)
	} else if *mode == ModeAsync {
		kmeans.RunAsync(dataset, *k, *static)
	} else if *mode == ModeChart {
		// Clean charts directory
		os.RemoveAll("./charts")
		os.Mkdir("./charts", os.ModeDir)

		kmeans.RunWithDrawing(dataset, *k, t, *static)

		// We are going to make a GIF when `gif` flag is true
		if *makeGif {
			kmeans.MakeGif(*t)
		}
	} else {
		fmt.Printf("Mode '%s' not supported!\n", *mode)
		os.Exit(1)
	}
}
