package main

import (
	"flag"
	"math/rand"
	"os"
	"time"

	"./kmeans"
)

func main() {
	var dataset []kmeans.Point
	rand.Seed(time.Now().UnixNano())

	var t *int
	t = new(int)
	*t = 0

	k := flag.Int("k", 0, "number of clusters")
	size := flag.Int("n", 0, "number of elements")
	makeGif := flag.Int("gif", 0, "wheter make gif or not")
	flag.Parse()

	if *k == 0 || *size == 0 {
		flag.Usage()
		os.Exit(1)
	}

	// Generate some random points
	for i := 0; i < *size; i++ {
		dataset = append(dataset, kmeans.Point{X: rand.Float64(), Y: rand.Float64()})
	}

	// Clean charts directory
	os.RemoveAll("./charts")
	os.Mkdir("./charts", os.ModeDir)

	//Runs and outputs charts for k clusters
	kmeans.RunWithDrawing(dataset, *k, t)

	if *makeGif == 1 {
		kmeans.MakeGif(*t)
	}
}
