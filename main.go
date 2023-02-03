package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
)

type order struct {
	EaterID    int
	FoodMenuID int
}

type menuCount struct {
	FoodMenuID int
	Count      int
}

func main() {
	topThreeMenuItems("orders.txt")
}

func topThreeMenuItems(filePath string) ([]int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()

	fi, err := file.Stat()
	if err != nil {
		return nil, fmt.Errorf("failed to stat orders file: %w", err)
	}
	if fi.Size() == 0 {
		return nil, fmt.Errorf("orders file is empty")
	}

	orders := make(map[int]map[int]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var o order
		_, err := fmt.Sscanf(scanner.Text(), "(%d, %d)", &o.EaterID, &o.FoodMenuID)
		if err != nil {
			fmt.Println("Error reading line:", err)
			continue
		}

		if _, exists := orders[o.EaterID]; !exists {
			orders[o.EaterID] = make(map[int]int)
		}
		if _, exists := orders[o.EaterID][o.FoodMenuID]; exists {
			//fmt.Printf("Error: eater %d has already ordered menu %d\n", o.EaterID, o.FoodMenuID)
			return nil, errors.New(fmt.Sprintf("Error: eater %d has already ordered menu %d\n", o.EaterID, o.FoodMenuID))
		}
		orders[o.EaterID][o.FoodMenuID] = 1
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	counts := make(map[int]int)
	for _, v := range orders {
		for k := range v {
			counts[k]++
		}
	}

	var menuCounts []menuCount
	for k, v := range counts {
		menuCounts = append(menuCounts, menuCount{FoodMenuID: k, Count: v})
	}

	sort.Slice(menuCounts, func(i, j int) bool {
		return menuCounts[i].Count > menuCounts[j].Count
	})

	result := make([]int, 0)
	fmt.Println("Top 3 menu items consumed:")
	for i := 0; i < 3 && i < len(menuCounts); i++ {
		fmt.Printf("%d. Menu %d with %d orders\n", i+1, menuCounts[i].FoodMenuID, menuCounts[i].Count)
		result = append(result, menuCounts[i].FoodMenuID)
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})
	return result, nil
}
