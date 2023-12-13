package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Node struct {
	value int
	left  *Node
	right *Node
}

type list struct {
	root  *Node
	end   *Node
	count int
}

func (queue *list) New(data int) *Node {
	var block *Node = new(Node)

	block.value = data
	block.left = nil
	block.right = nil

	return block
}

func (queue *list) add(data int) {
	var block *Node = queue.New(data)

	if queue.count == 0 {
		queue.root = block
		queue.end = queue.root
	} else {
		block.left = queue.end
		queue.end.right = block
		queue.end = block

	}
	queue.count++
}

func (queue *list) delet() int {
	var data int

	if queue.root != nil {
		data = queue.root.value

		if queue.count == 1 {
			queue.root = nil
			queue.end = queue.root

		} else {
			queue.root = queue.root.right

		}

		queue.count--

	}

	return data
}
func (queue *list) addMiddle(data, value int) {
	var newBlock *Node = queue.New(data)
	flag := false

	for block := queue.root; block != queue.end; block = block.right {
		if block.value == value {
			newBlock.left = block
			newBlock.right = block.right
			block.right = newBlock
			queue.count++
			flag = true
			break
		}

	}
	if flag == false {
		queue.add(data)
	}
}

func (queue *list) printList() {
	for block := queue.root; block != nil; block = block.right {
		fmt.Print(block.value+1, " ")
	}
}

func bubbleSort(arr *[][3]float64) {
	n := len(*arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if (*arr)[j][0] > (*arr)[j+1][0] {
				(*arr)[j][0], (*arr)[j+1][0] = (*arr)[j+1][0], (*arr)[j][0]
				(*arr)[j][1], (*arr)[j+1][1] = (*arr)[j+1][1], (*arr)[j][1]
				(*arr)[j][2], (*arr)[j+1][2] = (*arr)[j+1][2], (*arr)[j][2]
			}
		}
	}
}
func roulette(matrix [][]int, manyCities []int, currentCity []int) (int, int) {
	var (
		tapeMeasureSize       float64
		countCity             int
		nextCity              int
		previousCity          int
		nextCityProbabilities float64
		i                     int
	)

	for jndex, _ := range currentCity {
		for index, _ := range matrix[currentCity[jndex]] {
			if manyCities[index] == 0 && matrix[currentCity[jndex]][index] != 0 {
				fmt.Println(currentCity[jndex]+1, "->", index+1, "=", 1.0/float64(matrix[currentCity[jndex]][index]))
				tapeMeasureSize += 1.0 / float64(matrix[currentCity[jndex]][index])
				countCity++

			}
		}
	}
	probabilities := make([][3]float64, countCity)
	fmt.Println("tape measure size:", tapeMeasureSize)

	for jndex, _ := range currentCity {
		for index, _ := range matrix[currentCity[jndex]] {
			if manyCities[index] == 0 && matrix[currentCity[jndex]][index] != 0 {
				probabilities[i][0] = (1.0 / float64(matrix[currentCity[jndex]][index])) / tapeMeasureSize
				probabilities[i][1] = float64(index)
				probabilities[i][2] = float64(currentCity[jndex])
				fmt.Println(currentCity[jndex]+1, "->", index+1, "=", probabilities[i][0])
				i++
			}
		}
	}

	bubbleSort(&probabilities)

	rand.Seed(time.Now().UnixNano())
	randomSector := rand.Float64() * tapeMeasureSize
	fmt.Println("random sector:", randomSector)

	for index := 0; nextCityProbabilities <= randomSector; index++ {
		if index == countCity {
			break
		}
		nextCityProbabilities = probabilities[index][0]
		nextCity = int(probabilities[index][1])
		previousCity = int(probabilities[index][2])

	}

	return previousCity, nextCity
}

func nearestNeighbor(transitionMatrix [][]int, visitingCity int) {
	var (
		distance   int
		pathLength int
		city       int
	)

	currentCity := make([]int, 1)
	manyCities := make([]int, len(transitionMatrix[0]))
	res := make([]int, len(transitionMatrix[0])+1)

	res[0] = visitingCity + 1
	manyCities[visitingCity] = 1

	for index := 0; index < len(transitionMatrix[0])-1; index++ {

		fmt.Println("i:", index)
		fmt.Print("currentPath:")
		for j, _ := range res {
			if res[j] != 0 {
				fmt.Print(res[j], " ")
			}
		}
		fmt.Println()

		currentCity[0] = visitingCity
		_, city = roulette(transitionMatrix, manyCities, currentCity)
		distance = transitionMatrix[visitingCity][city]
		fmt.Println("visitingCite:", city+1)
		fmt.Println(visitingCity+1, "->", city+1, "=", distance)

		visitingCity = city
		pathLength += distance

		res[index+1] = visitingCity + 1
		manyCities[visitingCity] = 1

		fmt.Println("pathLength:", pathLength)

	}
	res[len(res)-1] = res[0]
	pathLength += transitionMatrix[visitingCity][res[0]-1]
	fmt.Println("\nresult:")
	fmt.Print("Path:", res, "\nPathLength:", pathLength)
}

func nearestCity(transitionMatrix [][]int, visitingCity int) {
	var (
		distance   int
		pathLength int
		firstCity  int
		secondCity int
	)

	manyCities := make([]int, len(transitionMatrix[0]))
	res := new(list)

	res.add(visitingCity)
	manyCities[visitingCity] = 1

	for index := 0; index < len(transitionMatrix[0])-1; index++ {
		currentCity := make([]int, 0)

		fmt.Println("i:", index)
		fmt.Print("currentPath:")

		res.printList()
		fmt.Println()

		distance = 0

		for jndex, _ := range manyCities {
			if manyCities[jndex] == 1 {
				currentCity = append(currentCity, jndex)
			}
		}

		firstCity, secondCity = roulette(transitionMatrix, manyCities, currentCity)
		distance = transitionMatrix[firstCity][secondCity]

		fmt.Println("\nvisitingCite:", secondCity+1)
		fmt.Println(firstCity+1, "->", secondCity+1, "=", distance)

		visitingCity = secondCity
		pathLength += distance

		res.addMiddle(secondCity, firstCity)
		manyCities[visitingCity] = 1

		fmt.Println("pathLength:", pathLength)

	}
	pathLength += transitionMatrix[res.end.value][res.root.value]
	res.add(res.root.value)
	fmt.Println("\nresult:")
	fmt.Print("Path:")
	res.printList()
	fmt.Println("\nPathLength:", pathLength)
}
func main() {

	var (
		visitingCity int
		algorithm    int
	)

	transitionMatrix := [][]int{
		{0, 3, 10, 8, 12},
		{3, 0, 11, 5, 9},
		{10, 11, 0, 6, 10},
		{8, 5, 6, 0, 4},
		{12, 9, 10, 4, 0},
	}
	fmt.Println("<Traveling salesman problem>")
	fmt.Println("Choose a greedy algorithm")
	fmt.Println("1 - Nearest neighbor")
	fmt.Println("2 - Nearest city")

	fmt.Print("-->")
	fmt.Scan(&algorithm)

	rand.Seed(time.Now().UnixNano())
	visitingCity = rand.Intn(5)

	switch algorithm {
	case 1:
		nearestNeighbor(transitionMatrix, visitingCity)
		break
	case 2:
		nearestCity(transitionMatrix, visitingCity)
		break
	default:
		fmt.Println("invalid command")
	}

}
