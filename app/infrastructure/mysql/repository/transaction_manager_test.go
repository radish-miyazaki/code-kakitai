package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/radish-miyazaki/code-kakitai/application/transaction"
	"github.com/radish-miyazaki/code-kakitai/infrastructure/mysql/db"
	"github.com/radish-miyazaki/code-kakitai/infrastructure/mysql/db/db_gen"
)

type transactionManager struct{}

func NewTransactionManager() *transaction.TransactionManager {
	return &transactionManager{}
}

func (tm *transactionManager) RunInTransaction(ctx context.Context, fn func(context.Context) error) error {
	dbCon := db.GetDBConn()
	tx, err := dbCon.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := db_gen.New(tx)
	ctxWithQueries := db.WithQueries(ctx, q)

	err = fn(ctxWithQueries)
	if err != nil {
		log.Printf("db rollback: %v\n", err)
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit failed: %w", err)
	}

	return nil
}
