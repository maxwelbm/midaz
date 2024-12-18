package command

import (
	"context"
	"errors"
	"go.uber.org/mock/gomock"
	"testing"

	"github.com/LerianStudio/midaz/components/ledger/internal/adapters/postgres/product"
	"github.com/LerianStudio/midaz/pkg"
	"github.com/LerianStudio/midaz/pkg/mmodel"

	"github.com/stretchr/testify/assert"
)

// TestCreateProductSuccess is responsible to test CreateProduct with success
func TestCreateProductSuccess(t *testing.T) {
	p := &mmodel.Product{
		ID:             pkg.GenerateUUIDv7().String(),
		OrganizationID: pkg.GenerateUUIDv7().String(),
		LedgerID:       pkg.GenerateUUIDv7().String(),
	}

	uc := UseCase{
		ProductRepo: product.NewMockRepository(gomock.NewController(t)),
	}

	uc.ProductRepo.(*product.MockRepository).
		EXPECT().
		Create(gomock.Any(), p).
		Return(p, nil).
		Times(1)
	res, err := uc.ProductRepo.Create(context.TODO(), p)

	assert.Equal(t, p, res)
	assert.Nil(t, err)
}

// TestCreateProductError is responsible to test CreateProduct with error
func TestCreateProductError(t *testing.T) {
	errMSG := "err to create product on database"
	p := &mmodel.Product{
		ID:             pkg.GenerateUUIDv7().String(),
		OrganizationID: pkg.GenerateUUIDv7().String(),
		LedgerID:       pkg.GenerateUUIDv7().String(),
	}

	uc := UseCase{
		ProductRepo: product.NewMockRepository(gomock.NewController(t)),
	}

	uc.ProductRepo.(*product.MockRepository).
		EXPECT().
		Create(gomock.Any(), p).
		Return(nil, errors.New(errMSG)).
		Times(1)
	res, err := uc.ProductRepo.Create(context.TODO(), p)

	assert.NotEmpty(t, err)
	assert.Equal(t, err.Error(), errMSG)
	assert.Nil(t, res)
}
