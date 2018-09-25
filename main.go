package main

import (
	"./kmeans"
	"flag"
	"math/rand"
	"os"
	"time"
)

func main() {
	var dataset []kmeans.Point
	rand.Seed(time.Now().UnixNano())

	k := flag.Int("k", 0, "number of clusters")
	size := flag.Int("n", 0, "number of elements")
	flag.Parse()

	if *k == 0 || *size == 0 {
		flag.Usage()
		os.Exit(1)
	}

	// Generate some random points
	for i := 0; i < *size; i++ {
		dataset = append(dataset, kmeans.Point{rand.Float64(), rand.Float64()})
	}

	// Clean charts directory
	os.RemoveAll("./charts")
	os.Mkdir("./charts", os.ModeDir)

	//Runs and outputs charts for k clusters
	kmeans.RunWithDrawing(dataset, *k)
}
