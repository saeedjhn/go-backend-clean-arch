package main

// Design a route planning system that supports different algorithms like shortest distance, least traffic, or fastest time.

type RoutePlanningAlgorithm interface {
	FindRoute(source, destination string) []string
}

type ShortestDistance struct{}

func (s *ShortestDistance) FindRoute(source, destination string) []string {
	// Please Implement your own 🥰 shortest distance algorithm
	panic("IMPL ME")
}

type LeastTraffic struct{}

func (l *LeastTraffic) FindRoute(source, destination string) []string {
	// Please Implement your own 🥰 least traffic algorithm
	panic("IMPL ME")
}

type FastestTime struct{}

func (f *FastestTime) FindRoute(source, destination string) []string {
	// Please Implement your own 🥰 fastest time algorithm
	panic("IMPL ME")
}

type RoutePlanner struct {
	routePlanningAlgorithm RoutePlanningAlgorithm
}

func (r *RoutePlanner) SetRoutePlanningAlgorithm(algorithm RoutePlanningAlgorithm) {
	r.routePlanningAlgorithm = algorithm
}

func (r *RoutePlanner) PlanRoute(source, destination string) []string {
	return r.routePlanningAlgorithm.FindRoute(source, destination)
}

// func main() {
// 	routePlanner := &RoutePlanner{}
//
// 	shortestDistance := &ShortestDistance{}
// 	leastTraffic := &LeastTraffic{}
// 	fastestTime := &FastestTime{}
//
// 	source := "A"
// 	destination := "B"
//
// 	routePlanner.SetRoutePlanningAlgorithm(shortestDistance)
// 	fmt.Println("Shortest Distance Route:", routePlanner.PlanRoute(source, destination))
//
// 	routePlanner.SetRoutePlanningAlgorithm(leastTraffic)
// 	fmt.Println("Least Traffic Route:", routePlanner.PlanRoute(source, destination))
//
// 	routePlanner.SetRoutePlanningAlgorithm(fastestTime)
// 	fmt.Println("Fastest Time Route:", routePlanner.PlanRoute(source, destination))
// }
