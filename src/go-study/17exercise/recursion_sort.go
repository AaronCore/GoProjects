package main

import "fmt"

type Item struct {
	Id    int
	Type  string
	Child *Item
}

type ItemClassifier interface {
	IsDoll() bool
}

func (i *Item) IsDoll() bool {
	if i.Type == "doll" {
		return true
	}
	return false
}

func findDiamond(item Item) Item {
	if item.IsDoll() {
		return findDiamond(*item.Child)
	} else {
		return item
	}
}

// 递归排序
func main() {
	doll := Item{
		Id:   1,
		Type: "doll",
		Child: &Item{
			Id:   2,
			Type: "doll",
			Child: &Item{
				Id:   3,
				Type: "doll",
				Child: &Item{
					Id:    4,
					Type:  "diamond",
					Child: nil,
				},
			},
		},
	}

	diamond := findDiamond(doll)
	fmt.Printf("Item %d is diamond\n", diamond.Id)
}
