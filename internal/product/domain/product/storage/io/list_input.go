package io

import "github.com/eNViDAT0001/GolangAdventure/external/paging"

type ListProductInput struct {
	ProductIDs  []uint
	CategoryIDs []uint
	Paging      paging.ParamsInput
}
