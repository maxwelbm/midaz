package command

import (
	"context"
	"errors"
	"fmt"
	"reflect"

	"github.com/LerianStudio/midaz/common"
	"github.com/LerianStudio/midaz/common/mlog"
	"github.com/LerianStudio/midaz/components/ledger/internal/app"
	a "github.com/LerianStudio/midaz/components/ledger/internal/domain/portfolio/account"
	"github.com/google/uuid"
)

// UpdateAccountByID update an account from the repository by given id.
func (uc *UseCase) UpdateAccountByID(ctx context.Context, id string, balance *a.Balance) (*a.Account, error) {
	logger := mlog.NewLoggerFromContext(ctx)
	logger.Infof("Trying to update account by id: %v", id)

	account := &a.Account{
		Balance: *balance,
	}

	accountUpdated, err := uc.AccountRepo.UpdateAccountById(ctx, uuid.MustParse(id), account)
	if err != nil {
		logger.Errorf("Error updating account on repo by id: %v", err)

		if errors.Is(err, app.ErrDatabaseItemNotFound) {
			return nil, common.EntityNotFoundError{
				EntityType: reflect.TypeOf(a.Account{}).Name(),
				Message:    fmt.Sprintf("Account with id %s was not found", id),
				Code:       "ACCOUNT_NOT_FOUND",
				Err:        err,
			}
		}

		return nil, err
	}

	return accountUpdated, nil
}
