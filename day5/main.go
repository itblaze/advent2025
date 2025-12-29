package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	// declare vars over here
	// fileName := "test.txt"
	fileName := "advent.txt"
	filePath := "/Users/samanyanga/Dev/projects/advent2025/day5/data"
	dataArr := make([]string, 0, 10)

	println("Beginning of advent of code day 5...")
	fmt.Printf("Reading file: %v from location %v\n", fileName, filePath)

	file, _ := os.Open(fmt.Sprintf("%v/%v", filePath, fileName))
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		dataArr = append(dataArr, scanner.Text())
	}

	// fmt.Printf("Solution Part 1: %d\n", solution1(dataArr))
	fmt.Printf("Solution Part 2: %d\n", solution2(dataArr))

}

type Vertex struct {
	Lo int64
	Hi int64
}

func solution1(dataArr []string) int {
	ranges := make([]Vertex, 0, 10)
	ids := make([]int64, 0, 10)

	for _, line := range dataArr {
		if line == "" {
			continue
		}
		if strings.Contains(line, "-") {
			// a range
			rangesS := strings.Split(line, "-")
			ranges0, _ := strconv.Atoi(rangesS[0])
			ranges1, _ := strconv.Atoi(rangesS[1])
			ranges = append(ranges, Vertex{int64(ranges0), int64(ranges1)})
		}
		idAsInt, _ := strconv.Atoi(line)
		ids = append(ids, int64(idAsInt))
	}

	count := 0

	for _, val := range ids {
		spoiled := true
		for _, vert := range ranges {
			if val >= vert.Lo && val <= vert.Hi {
				spoiled = false
				break
			}
		}
		if !spoiled {
			count++
		}
	}
	return count
}

func isValid(i, j, iMax, jMax int) bool {
	if i < 0 || i >= iMax || j < 0 || j >= jMax {
		return false
	}
	return true
}

func isAt(testChar rune) bool {
	return testChar == '@'
}

func toRuneDoubleSlice(dataArr []string) *[][]rune {
	var rneArr [][]rune
	rneArr = make([][]rune, 0, 10)

	for _, line := range dataArr {
		runeArray := []rune(line)
		rneArr = append(rneArr, runeArray)
	}
	return &rneArr
}

func solution2(dataArr []string) int64 {
	ranges := make([]Vertex, 0, 10)
	ids := make([]int64, 0, 10)

	for _, line := range dataArr {
		if line == "" {
			continue
		}
		if strings.Contains(line, "-") {
			// a range
			rangesS := strings.Split(line, "-")
			ranges0, _ := strconv.Atoi(rangesS[0])
			ranges1, _ := strconv.Atoi(rangesS[1])
			ranges = append(ranges, Vertex{int64(ranges0), int64(ranges1)})
		}
		idAsInt, _ := strconv.Atoi(line)
		ids = append(ids, int64(idAsInt))
	}
	// flatten the ranges
	rangeList := list.New()

	// fmt.Printf("Size of initial ranges before flattening: %v\n", len(ranges))
	for _, vertex := range ranges {
		if rangeList.Len() == 0 {
			rangeList.PushBack(vertex)
			continue
		}
		combined := false
		for e := rangeList.Front(); e != nil; e = e.Next() {
			val := e.Value.(Vertex)

			if val.Lo < vertex.Lo && val.Hi > vertex.Hi {
				// new vertex can be swallowed and ignored
				combined = true
			}
			if vertex.Lo < val.Lo && vertex.Hi > val.Hi {
				// old vertex can be swallowed; remove old vertex from list
				combined = true
				val.Lo = vertex.Lo
				val.Hi = vertex.Hi
				e.Value = val
			}
			if vertex.Lo > val.Lo && vertex.Lo < val.Hi && vertex.Hi > val.Hi {
				// can be combined to be val.Lo -> vertex.Hi
				combined = true
				val.Hi = vertex.Hi
				e.Value = val
			}
			if val.Lo > vertex.Lo && vertex.Hi > val.Lo && vertex.Hi > val.Hi {
				// can be combined to be vert.Lo -> val.Hi
				combined = true
				val.Lo = vertex.Lo
				e.Value = val
			}
		}
		if !combined {
			// was not coalesced into other ranges, add seperately
			rangeList.PushBack(vertex)
		}
	}

	flattedList := flattenRanges(rangeList)
	printList(flattedList)

	var totalCount int64 = 0

	for e := flattedList.Front(); e != nil; e = e.Next() {
		v := e.Value.(Vertex)
		totalCount += (v.Hi - v.Lo + 1)
	}
	return totalCount
}

func printList(l *list.List) {
	fmt.Print("List contents: ")
	for e := l.Front(); e != nil; e = e.Next() {
		value := e.Value.(Vertex)
		fmt.Printf("%v-%v ", value.Lo, value.Hi)
	}
	fmt.Print("\n")
}

func flattenRangesOld(l *list.List) *list.List {
	// copy list
	cl := list.New()
	rangeList := list.New()
	cl.PushBackList(l)

	oSize := cl.Len()

	for e := cl.Front(); e != nil; e = e.Next() {
		vertex := e.Value.(Vertex)
		if rangeList.Len() == 0 {
			rangeList.PushBack(vertex)
			continue
		}
		combined := false
		for e := rangeList.Front(); e != nil; e = e.Next() {
			val := e.Value.(Vertex)

			if val.Lo <= vertex.Lo && val.Hi >= vertex.Hi {
				// new vertex can be swallowed and ignored
				combined = true
				break
			}
			if vertex.Lo <= val.Lo && vertex.Hi >= val.Hi {
				// old vertex can be swallowed; remove old vertex from list
				combined = true
				val.Lo = vertex.Lo
				val.Hi = vertex.Hi
				e.Value = val

				newVert := Vertex{Lo: vertex.Lo, Hi: vertex.Hi}
				rangeList.PushBack(newVert)
				rangeList.Remove(e)
				break
			}
			if vertex.Lo >= val.Lo && vertex.Lo <= val.Hi && vertex.Hi >= val.Hi {
				// can be combined to be val.Lo -> vertex.Hi
				combined = true
				val.Hi = vertex.Hi
				e.Value = val
				break
			}
			if val.Lo >= vertex.Lo && vertex.Hi >= val.Lo && vertex.Hi >= val.Hi {
				// can be combined to be vert.Lo -> val.Hi
				combined = true
				val.Lo = vertex.Lo
				e.Value = val
				break
			}
		}
		if !combined {
			// was not coalesced into other ranges, add seperately
			rangeList.PushBack(vertex)
		}
	}
	nSize := rangeList.Len()
	if nSize == oSize {
		// no further flattening possible, return
		return rangeList
	}
	return flattenRangesOld(rangeList)
}

func flattenRanges(l *list.List) *list.List {
	// convert to slice
	verts := make([]Vertex, 0, l.Len())
	for e := l.Front(); e != nil; e = e.Next() {
		verts = append(verts, e.Value.(Vertex))
	}
	if len(verts) == 0 {
		return list.New()
	}

	// sort by lower bound
	sort.Slice(verts, func(i, j int) bool { return verts[i].Lo < verts[j].Lo })

	// single-pass merge (merges overlaps and contiguous ranges)
	merged := make([]Vertex, 0, len(verts))
	cur := verts[0]
	for _, v := range verts[1:] {
		if v.Lo <= cur.Hi+1 { // overlap or contiguous
			if v.Hi > cur.Hi {
				cur.Hi = v.Hi
			}
		} else {
			merged = append(merged, cur)
			cur = v
		}
	}
	merged = append(merged, cur)

	// build list from merged slice
	out := list.New()
	for _, mv := range merged {
		out.PushBack(mv)
	}

	// if no change in number of ranges, we're done; otherwise recurse to be safe
	if out.Len() == l.Len() {
		return out
	}
	return flattenRanges(out)
}

func accessTissuesSln2(rneArr *[][]rune) int {
	iLen := len(*rneArr)
	jLen := len((*rneArr)[0])
	accessible := 0
	tp := 0

	for i := range iLen {
		for j := range jLen {
			// is co-ordinate a roll of paper
			if (*rneArr)[i][j] != '@' {
				continue
			}
			var char rune
			// i-1, j
			if isValid(i-1, j, iLen, jLen) {
				char = (*rneArr)[i-1][j]
				if isAt(char) {
					tp++
				}
			}
			// i+1, j
			if isValid(i+1, j, iLen, jLen) {
				char = (*rneArr)[i+1][j]
				if isAt(char) {
					tp++
				}
			}
			// i, j-1
			if isValid(i, j-1, iLen, jLen) {
				char = (*rneArr)[i][j-1]
				if isAt(char) {
					tp++
				}
			}
			// i, j+1
			if isValid(i, j+1, iLen, jLen) {
				char = (*rneArr)[i][j+1]
				if isAt(char) {
					tp++
				}
			}

			// i-1, j-1
			if isValid(i-1, j-1, iLen, jLen) {
				char = (*rneArr)[i-1][j-1]
				if isAt(char) {
					tp++
				}
			}
			// i-1, j+1
			if isValid(i-1, j+1, iLen, jLen) {
				char = (*rneArr)[i-1][j+1]
				if isAt(char) {
					tp++
				}
			}
			// i+1, j-1
			if isValid(i+1, j-1, iLen, jLen) {
				char = (*rneArr)[i+1][j-1]
				if isAt(char) {
					tp++
				}
			}
			// i+1, j+1
			if isValid(i+1, j+1, iLen, jLen) {
				char = (*rneArr)[i+1][j+1]
				if isAt(char) {
					tp++
				}
			}

			if tp < 4 {
				accessible++

				// mark the spot with a .
				(*rneArr)[i][j] = '.'
			}
			tp = 0

		}
	}
	return accessible
}
