// Package entity representation of api objects
package entity

type City struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Borders []int  `json:"borders"`
}
