package handler

import (
	"github.com/auster-kaki/auster-mono/pkg/app/usecase"
)

type VendorHandler struct {
	vendorUseCase usecase.VendorUseCase
}

func NewVendorHandler(u *usecase.VendorUseCase) []Input {
	//h := &VendorHandler{vendorUseCase: u}
	return []Input{}
}
