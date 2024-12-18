package query

import (
	"context"
	"errors"
	"go.uber.org/mock/gomock"
	"testing"

	"github.com/LerianStudio/midaz/components/ledger/internal/adapters/postgres/ledger"
	"github.com/LerianStudio/midaz/pkg"
	"github.com/LerianStudio/midaz/pkg/mmodel"

	"github.com/stretchr/testify/assert"
)

// TestGetAllLedgersError is responsible to test GetAllLedgers with success and error
func TestGetAllLedgers(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockLedgerRepo := ledger.NewMockRepository(ctrl)
	organizationID := pkg.GenerateUUIDv7()
	limit := 10
	page := 1

	uc := UseCase{
		LedgerRepo: mockLedgerRepo,
	}

	t.Run("Success", func(t *testing.T) {
		ledgers := []*mmodel.Ledger{{}}
		mockLedgerRepo.
			EXPECT().
			FindAll(gomock.Any(), organizationID, page, limit).
			Return(ledgers, nil).
			Times(1)
		res, err := uc.LedgerRepo.FindAll(context.TODO(), organizationID, page, limit)

		assert.NoError(t, err)
		assert.Len(t, res, 1)
	})

	t.Run("Error", func(t *testing.T) {
		errMsg := "errDatabaseItemNotFound"
		mockLedgerRepo.
			EXPECT().
			FindAll(gomock.Any(), organizationID, page, limit).
			Return(nil, errors.New(errMsg)).
			Times(1)
		res, err := uc.LedgerRepo.FindAll(context.TODO(), organizationID, page, limit)

		assert.EqualError(t, err, errMsg)
		assert.Nil(t, res)
	})
}
