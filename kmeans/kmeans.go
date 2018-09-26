package kmeans

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func centersX(clusters []Cluster) (centersX []float64) {
	for i := 0; i < len(clusters); i++ {
		centersX = append(centersX, clusters[i].Center.X)
	}
	return
}

func centersY(clusters []Cluster) (centersY []float64) {
	for i := 0; i < len(clusters); i++ {
		centersY = append(centersY, clusters[i].Center.Y)
	}
	return
}

func getClusters(k int, static bool) []Cluster {
	var actualLen int

	if static {
		actualLen = 4
	} else {
		actualLen = k
	}
	clusters := make([]Cluster, actualLen, actualLen)

	if static {
		clusters[0] = Cluster{Point{0.915438, 0.760661}, []Point{}}
		clusters[1] = Cluster{Point{0.363143, 0.758002}, []Point{}}
		clusters[2] = Cluster{Point{0.863991, 0.823953}, []Point{}}
		clusters[3] = Cluster{Point{0.924190, 0.716877}, []Point{}}
		return clusters
	}

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < k; i++ {
		clusters[i] = Cluster{Point{rand.Float64(), rand.Float64()}, []Point{}}
	}
	return clusters
}

func logClusters(clusters []Cluster) {
	for i := 0; i < len(clusters); i++ {
		fmt.Printf("Center %d -> X: %f, Y: %f\n", i+1, clusters[i].Center.X, clusters[i].Center.Y)
	}
}

func repositionCenters(clusters []Cluster) {
	for i := 0; i < len(clusters); i++ {
		clusters[i].repositionCenter()
	}
}

/*RunSync runs synchronously */
func RunSync(dataset []Point, k int, static bool) []Cluster {
	start := time.Now()

	t := 0
	pointCenterIsDifferent := true

	pointsClusterIndex := make([]int, len(dataset))
	clusters := getClusters(k, static)

	// Just a dumb loop
	for ; pointCenterIsDifferent; t++ {
		pointCenterIsDifferent = false

		// We loop through all the points
		for i := 0; i < len(dataset); i++ {
			var minDist float64
			var updatedClusterIndex int

			// Dummy loop just to check which center is the nearest
			// to the current point
			for j := 0; j < len(clusters); j++ {
				tmpDist := dataset[i].Distance(clusters[j].Center)
				if minDist == 0 || tmpDist < minDist {
					minDist = tmpDist
					updatedClusterIndex = j
				}
			}

			clusters[updatedClusterIndex].Points = append(clusters[updatedClusterIndex].Points, dataset[i])

			// Continue condition: if the new index is different than the previous we continue
			if pointsClusterIndex[i] != updatedClusterIndex {
				pointsClusterIndex[i] = updatedClusterIndex
				pointCenterIsDifferent = true
			}
		}

		if pointCenterIsDifferent {
			// Reposition each center to the its mean
			repositionCenters(clusters)
		}
	}

	logClusters(clusters)
	elapsed := time.Since(start)
	log.Printf("Sync algorithm with %s iterations took %s", strconv.Itoa(t), elapsed)
	dodraw(clusters, "charts/sync.png")
	return clusters
}

var wg sync.WaitGroup

/*RunAsync runs asynchronously */
func RunAsync(dataset []Point, k int, static bool) []Cluster {
	start := time.Now()

	t := 0
	pointCenterIsDifferent := true

	pointsClusterIndex := make([]int, len(dataset))
	clusters := getClusters(k, static)
	numberOfPoints := len(dataset)

	// Just a dumb loop
	for ; pointCenterIsDifferent; t++ {
		pointCenterIsDifferent = false

		wg.Add(numberOfPoints)
		// We loop through all the points
		for i := 0; i < numberOfPoints; i++ {
			go func(iC int) {
				var minDist float64
				var updatedClusterIndex int

				// Dummy loop just to check which center is the nearest
				// to the current point
				for j := 0; j < len(clusters); j++ {
					tmpDist := dataset[iC].Distance(clusters[j].Center)
					if minDist == 0 || tmpDist < minDist {
						minDist = tmpDist
						updatedClusterIndex = j
					}
				}

				clusters[updatedClusterIndex].Points = append(clusters[updatedClusterIndex].Points, dataset[iC])

				// Continue condition: if the new index is different than the previous we continue
				if pointsClusterIndex[iC] != updatedClusterIndex {
					pointsClusterIndex[iC] = updatedClusterIndex
					pointCenterIsDifferent = true
				}
				wg.Done()
			}(i)
		}

		wg.Wait()

		if pointCenterIsDifferent {
			// Reposition each center to the its mean
			repositionCenters(clusters)
		}
	}

	logClusters(clusters)
	elapsed := time.Since(start)
	log.Printf("Async algorithm with %s iterations took %s", strconv.Itoa(t), elapsed)
	dodraw(clusters, "charts/async.png")
	return clusters
}

/*RunWithDrawing runs the k-means algorithm given an array of coordinates and a specific k*/
func RunWithDrawing(dataset []Point, k int, t *int, static bool) []Cluster {
	pointsClusterIndex := make([]int, len(dataset))
	clusters := getClusters(k, static)
	pointCenterIsDifferent := true

	for *t = 0; pointCenterIsDifferent; *t++ {
		pointCenterIsDifferent = false
		for i := 0; i < len(dataset); i++ {
			var minDist float64
			var updatedClusterIndex int
			for j := 0; j < len(clusters); j++ {
				tmpDist := dataset[i].Distance(clusters[j].Center)
				if minDist == 0 || tmpDist < minDist {
					minDist = tmpDist
					updatedClusterIndex = j
				}
			}
			clusters[updatedClusterIndex].Points = append(clusters[updatedClusterIndex].Points, dataset[i])
			if pointsClusterIndex[i] != updatedClusterIndex {
				pointsClusterIndex[i] = updatedClusterIndex
				pointCenterIsDifferent = true
			}
		}
		dodraw(clusters, "charts/"+strconv.Itoa(*t)+".png")
		if pointCenterIsDifferent {
			repositionCenters(clusters)
		}
	}

	return clusters
}
