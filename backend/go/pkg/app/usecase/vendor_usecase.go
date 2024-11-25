package usecase

import (
	"github.com/auster-kaki/auster-mono/pkg/app/repository"
)

type VendorUseCase struct {
	repository repository.Repository
}

func NewVendorUseCase(r repository.Repository) *VendorUseCase {
	return &VendorUseCase{repository: r}
}
