package model

import (
	"context"
	"database/sql"

	"github.com/ngavinsir/api-supply-demand-covid19/models"
	"github.com/segmentio/ksuid"
	"github.com/volatiletech/sqlboiler/boil"
	. "github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/types"
)

// HasGetAllStock handles get stock data.
type HasGetAllStock interface {
	GetAllStock(ctx context.Context, offset int, limit int) ([]*StockData, int64, error)
}

// HasCreateOrUpdateStock creates or updates stock.
type HasCreateOrUpdateStock interface {
	CreateOrUpdateStock(ctx context.Context, stock *models.Stock, tx *sql.Tx) (*models.Stock, error)
}

// HasIsStockAvailable interface
type HasIsStockAvailable interface {
	IsStockAvailable(ctx context.Context, itemID string, unitID string, quantity types.Decimal) (bool, error)
}

// StockDataStore holds db information.
type StockDataStore struct {
	*sql.DB
}

// GetAllStock returns stocks
func (db *StockDataStore) GetAllStock(ctx context.Context, offset int, limit int) ([]*StockData, int64, error) {
	stocksCount, err := models.Stocks().Count(ctx, db)
	if err != nil {
		return nil, 0, err
	}

	stocks, err := models.Stocks(
		Offset(offset),
		Limit(limit),
		Load(models.StockRels.Item),
		Load(models.StockRels.Unit),
	).All(ctx, db)
	if err != nil {
		return nil, 0, err
	}

	stockData := []*StockData{}
	for _, stock := range stocks {
		data := &StockData{
			ID:       stock.ID,
			Name:     stock.R.Item.Name,
			Unit:     stock.R.Unit.Name,
			Quantity: stock.Quantity,
		}
		stockData = append(stockData, data)
	}

	return stockData, stocksCount, err
}

// CreateOrUpdateStock creates a new stock if given item id in given unit doesn't exist or otherwise
// add given quantity to current stock.
func (db *StockDataStore) CreateOrUpdateStock(ctx context.Context, data *models.Stock, parentTx *sql.Tx) (*models.Stock, error) {
	var err error
	var tx *sql.Tx
	if parentTx == nil {
		tx, err = db.BeginTx(ctx, &sql.TxOptions{
			Isolation: sql.LevelSerializable,
		})
		if err != nil {
			return nil, err
		}
	} else {
		tx = parentTx
	}

	stock := &models.Stock{
		ID:       ksuid.New().String(),
		ItemID:   data.ItemID,
		UnitID:   data.UnitID,
		Quantity: data.Quantity,
	}

	if data.ID != "" {
		stock, err = models.FindStock(ctx, tx, data.ID)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		stock.Quantity.Add(stock.Quantity.Big, data.Quantity.Big)
	} else {
		stocks, err := models.Stocks(
			models.StockWhere.ItemID.EQ(data.ItemID),
			models.StockWhere.UnitID.EQ(data.UnitID),
		).All(ctx, tx)
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		if len(stocks) > 0 {
			stock = stocks[0]
			stock.Quantity.Add(stock.Quantity.Big, data.Quantity.Big)
		}
	}

	if err := stock.Upsert(ctx, tx, true, []string{"id"}, boil.Infer(), boil.Infer()); err != nil {
		tx.Rollback()
		return nil, err
	}

	if parentTx == nil {
		if err := tx.Commit(); err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	return stock, nil
}

// IsStockAvailable returns true if given item id, unit id, and quantity exists in stock.
func (db *StockDataStore) IsStockAvailable(ctx context.Context, itemID string, unitID string, quantity types.Decimal) (bool, error) {
	isExist, err := models.Stocks(
		models.StockWhere.ItemID.EQ(itemID),
		models.StockWhere.UnitID.EQ(unitID),
		models.StockWhere.Quantity.GTE(quantity),
	).Exists(ctx, db)
	if err != nil {
		return false, err
	}

	return isExist, nil
}

// StockData struct
type StockData struct {
	ID       string        `boil:"id" json:"id"`
	Name     string        `boil:"name" json:"name"`
	Unit     string        `boil:"unit" json:"unit"`
	Quantity types.Decimal `boil:"quantity" json:"quantity"`
}
