package io

import "github.com/eNViDAT0001/GolangAdventure/internal/product/domain/product/storage/io"

type SpecificationCreateForm struct {
	Specification io.ProductSpecificationCreateForm
	Options       []io.ProductOptionCreateForm
}
