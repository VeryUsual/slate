package main

import (
	"fmt"
)

type Panel struct {
	AppID      string
	ID         string
	Title      string
	Workspace  int
}

type WMRoot struct {
	Node       any // WMNode or WMSplit
}

type WMNode struct {
	Panel      Panel
}

type WMSplit struct {
	Type       string
	Nodes      []any // WMNode or WMSplit
}

type DrawnPanel struct {
	X          int
	Y          int
	Width      int
	Height     int
	Panel      Panel
}

func handleSplit(n any, x int, y int, width int, height int) []DrawnPanel {
	drawnPanels := []DrawnPanel{}

	switch n_typed := n.(type) {
	case *WMNode:
		drawnPanels = append(drawnPanels, DrawnPanel{X: x, Y: y, Width: width, Height: height, Panel: n_typed.Panel})
	case *WMSplit:
		for i, node := range n_typed.Nodes {
			drawnPanels = append(drawnPanels, handleSplit(node, x + i*(width / 2), y, width / 2, height)...)
		}
	default:
		fmt.Println("this shouldnt happen")
	}

	return drawnPanels
}

func SplitFocusedPanel(n *any, focusedPanel *string) bool {
	switch n_typed := (*n).(type) {
	case *WMNode:
		if *focusedPanel == n_typed.Panel.AppID + "/" + n_typed.Panel.ID {
			*n = &WMSplit{
				Type: "h",
				Nodes: []any{
					n_typed,
					&WMNode{
						Panel: Panel{
							AppID: "author.app99",
							ID: "main",
							Title: "panel that got split",
						},
					},
				},
			}

			return true
		}
	case *WMSplit:
		for i := range n_typed.Nodes {
			if SplitFocusedPanel(&n_typed.Nodes[i], focusedPanel) {
				return true
			}
		}
	}

	return false
}