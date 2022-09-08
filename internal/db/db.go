package db

import (
	"context"
	"fmt"
	"github.com/UnTea/L0/internal/model"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
)

// DB - database connection structure
type DB struct {
	conn *pgxpool.Pool
	tx   pgx.Tx
}

// OrderData - wrapping for Item structure
type OrderData struct {
	OrderUID string
	Item     model.Items
}

// NewDatabaseInstance - database instance structure constructor
func NewDatabaseInstance(conn *pgxpool.Pool) *DB {
	return &DB{
		conn: conn,
	}
}

// Get - gets data from database for restoring in-memory cache
func (d *DB) Get(ctx context.Context) (data []model.Data, err error) {

	data, err = d.getOrder(ctx)
	if err != nil {
		return nil, err
	}

	items, err := d.getItems(ctx)
	if err != nil {
		return nil, err
	}

	for i, value := range data {
		for j, v := range items {
			if value.OrderUID == v.OrderUID {
				data[i].Items = append(data[i].Items, items[j].Item)
			}
		}
	}

	return data, nil
}

func (d *DB) getOrder(ctx context.Context) ([]model.Data, error) {
	rows, err := d.conn.Query(ctx,
		"SELECT order_uid, track_number, entry, locale, internal_signature, customer_id, "+
			"delivery_service, shardkey, sm_id, date_created, oof_shard, name, phone, zip, city, address, "+
			"region, email, transaction, request_id, currency, provider, payment_dt, bank, delivery_cost, "+
			"goods_total, custom_fee "+
			"FROM get_order();")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []model.Data

	for rows.Next() {
		var row model.Data
		values, err := rows.Values()
		if err != nil {
			return nil, err
		}

		row.OrderUID = values[0].(string)
		row.TrackNumber = values[1].(string)
		row.Entry = values[2].(string)
		row.Locale = values[3].(string)
		row.InternalSignature = values[4].(string)
		row.CustomerID = values[5].(string)
		row.DeliveryService = values[6].(string)
		row.ShardKey = values[7].(string)
		row.SmID = int(values[8].(int32))
		row.DateCreated = values[9].(time.Time).Format("2006-01-02T15:04:05Z0700")
		row.OofShard = values[10].(string)
		row.Delivery.Name = values[11].(string)
		row.Delivery.Phone = values[12].(string)
		row.Delivery.Zip = values[13].(string)
		row.Delivery.City = values[14].(string)
		row.Delivery.Address = values[15].(string)
		row.Delivery.Region = values[16].(string)
		row.Delivery.Email = values[17].(string)
		row.Payment.Transaction = values[18].(string)
		row.Payment.RequestID = values[19].(string)
		row.Payment.Currency = values[20].(string)
		row.Payment.Provider = values[21].(string)
		row.Payment.PaymentDt = int(values[22].(int32))
		row.Payment.Bank = values[23].(string)
		row.Payment.DeliveryCost = values[24].(float64)
		row.Payment.GoodsTotal = values[25].(float64)
		row.Payment.CustomFee = values[26].(float64)
		row.Payment.Amount = row.Payment.DeliveryCost + row.Payment.GoodsTotal + row.Payment.CustomFee
		data = append(data, row)
	}

	return data, nil
}

func (d *DB) getItems(ctx context.Context) ([]OrderData, error) {
	rows, err := d.conn.Query(ctx,
		"SELECT chrt_id, price, rid, name, sale, size, nm_id, brand, status, order_id From get_items();")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []OrderData

	for rows.Next() {
		var row OrderData

		values, err := rows.Values()
		if err != nil {
			return nil, err
		}

		row.Item.ChrtID = int(values[0].(int32))
		row.Item.Price = values[1].(float64)
		row.Item.Rid = values[2].(string)
		row.Item.Name = values[3].(string)
		row.Item.Sale = int(values[4].(int32))
		row.Item.Size = values[5].(string)
		row.Item.NmId = int(values[6].(int32))
		row.Item.Brand = values[7].(string)
		row.Item.Status = int(values[8].(int32))
		row.OrderUID = values[9].(string)
		sale := row.Item.Price * float64(row.Item.Sale) / float64(100)
		row.Item.TotalPrice = float64(int(row.Item.Price - sale))
		items = append(items, row)
	}

	return items, nil
}

// Set - add data to database
func (d *DB) Set(ctx context.Context, data model.Data) (string, error) {

	customer := model.CustomerDB{
		CustomerID: data.CustomerID,
		Name:       data.Delivery.Name,
		Phone:      data.Delivery.Phone,
		Zip:        data.Delivery.Zip,
		City:       data.Delivery.City,
		Address:    data.Delivery.Address,
		Region:     data.Delivery.Region,
		Email:      data.Delivery.Email,
	}

	order := model.OrderDB{
		OrderUID:          data.OrderUID,
		TrackNumber:       data.TrackNumber,
		Entry:             data.Entry,
		CustomerID:        data.CustomerID,
		Locale:            data.Locale,
		DeliveryService:   data.DeliveryService,
		ShardKey:          data.ShardKey,
		SmID:              data.SmID,
		DateCreated:       data.DateCreated,
		OofShard:          data.OofShard,
		InternalSignature: data.InternalSignature,
	}

	orderItems := make([]model.OrderItemsDB, 0, len(data.Items))
	items := make([]model.ItemDB, 0, len(data.Items))
	for _, value := range data.Items {
		item := model.ItemDB{
			ChrtID: value.ChrtID,
			Price:  value.Price,
			Rid:    value.Rid,
			Name:   value.Name,
			Sale:   value.Sale,
			Size:   value.Size,
			NmId:   value.NmId,
			Brand:  value.Brand,
		}
		orderItem := model.OrderItemsDB{
			OrderID: data.OrderUID,
			ItemID:  value.ChrtID,
			Status:  value.Status,
		}
		items = append(items, item)
		orderItems = append(orderItems, orderItem)
	}

	payment := model.PaymentDB{
		Transaction:  data.Payment.Transaction,
		RequestID:    data.Payment.RequestID,
		Currency:     data.Payment.Currency,
		Provider:     data.Payment.Provider,
		PaymentDt:    data.Payment.PaymentDt,
		Bank:         data.Payment.Bank,
		DeliveryCost: data.Payment.DeliveryCost,
		GoodsTotal:   data.Payment.GoodsTotal,
		CustomFee:    data.Payment.CustomFee,
	}

	// starting transaction
	tx, err := d.conn.Begin(ctx)
	if err != nil {
		return "", err
	}
	d.tx = tx

	err = d.addCustomer(ctx, customer)
	if err != nil {
		tx.Rollback(ctx)
		return "", fmt.Errorf("rollback transaction due to error: %v", err)
	}

	err = d.addOrder(ctx, order)
	if err != nil {
		tx.Rollback(ctx)
		return "", fmt.Errorf("rollback transaction due to error: %v", err)
	}

	for _, item := range items {
		err = d.addItems(ctx, item)
		if err != nil {
			tx.Rollback(ctx)
			return "", fmt.Errorf("rollback transaction due to error: %v", err)
		}
	}

	for _, orderItem := range orderItems {
		err = d.addOrderItems(ctx, orderItem)
		if err != nil {
			tx.Rollback(ctx)
			return "", fmt.Errorf("rollback transaction due to error: %v", err)
		}
	}

	err = d.addPayment(ctx, payment)
	if err != nil {
		tx.Rollback(ctx)
		return "", fmt.Errorf("rollback transaction due to error: %v", err)
	}

	// commit transaction
	err = d.tx.Commit(ctx)
	if err != nil {
		return "", err
	}

	return data.OrderUID, nil
}

func (d *DB) addCustomer(ctx context.Context, customer model.CustomerDB) error {
	_, err := d.tx.Exec(ctx,
		"SELECT add_customer($1, $2, $3, $4, $5, $6, $7, $8);",
		customer.CustomerID,
		customer.Name,
		customer.Phone,
		customer.Zip,
		customer.City,
		customer.Address,
		customer.Region,
		customer.Email,
	)
	if err != nil {
		return err
	}
	return nil
}

func (d *DB) addOrder(ctx context.Context, order model.OrderDB) error {
	_, err := d.tx.Exec(ctx,
		"SELECT add_order($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);",
		order.OrderUID,
		order.CustomerID,
		order.TrackNumber,
		order.Entry,
		order.Locale,
		order.DeliveryService,
		order.ShardKey,
		order.SmID,
		order.DateCreated,
		order.OofShard,
		order.InternalSignature,
	)
	if err != nil {
		return err
	}
	return nil
}

func (d *DB) addOrderItems(ctx context.Context, orderItems model.OrderItemsDB) error {
	_, err := d.tx.Exec(ctx,
		"SELECT add_order_items($1, $2, $3);",
		orderItems.OrderID,
		orderItems.ItemID,
		orderItems.Status,
	)
	if err != nil {
		return err
	}
	return nil
}

func (d *DB) addItems(ctx context.Context, items model.ItemDB) error {
	_, err := d.tx.Exec(ctx,
		"SELECT add_items($1, $2, $3, $4, $5, $6, $7, $8);",
		items.ChrtID,
		items.Price,
		items.Rid,
		items.Name,
		items.Sale,
		items.Size,
		items.NmId,
		items.Brand,
	)
	if err != nil {
		return err
	}
	return nil
}

func (d *DB) addPayment(ctx context.Context, payment model.PaymentDB) error {
	_, err := d.tx.Exec(ctx,
		"SELECT add_payment($1, $2, $3, $4, $5, $6, $7, $8, $9);",
		payment.Transaction,
		payment.RequestID,
		payment.Currency,
		payment.Provider,
		payment.PaymentDt,
		payment.Bank,
		payment.DeliveryCost,
		payment.GoodsTotal,
		payment.CustomFee,
	)
	if err != nil {
		return err
	}
	return nil
}
