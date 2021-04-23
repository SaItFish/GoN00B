// Package main
// @file: main.go
// @description:
// @author: SaltFish
// @date: 2020/09/20
package main

import "network-file/api/view"

// POST /api/v1/files/
// GET  /api/v1/files/:name

func main() {
	view.StartServer()
}
