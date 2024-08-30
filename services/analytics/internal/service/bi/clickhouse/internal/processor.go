package internal

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"go.uber.org/zap"

	"github.com/moke-game/platform/services/analytics/internal/service/bi"
)

type Processor struct {
	logger   *zap.Logger
	rootPath string
	conn     driver.Conn
}

func (p *Processor) Init(
	logger *zap.Logger, rootPath, addr, dbname, uname, passwd string,
) error {
	p.logger = logger.With(zap.String("data_processor", "clickhouse"))
	p.rootPath = rootPath
	if addr == "" {
		p.logger.Error("clickhouse addr empty")
		return nil
	}
	var err error
	if p.conn, err = clickhouse.Open(&clickhouse.Options{
		Addr: strings.Split(addr, ","),
		TLS:  &tls.Config{},
		Auth: clickhouse.Auth{
			Database: dbname,
			Username: uname,
			Password: passwd,
		},
	}); err != nil {
		p.logger.Error("clickhouse init fail", zap.Error(err))
	} else if err := p.conn.Ping(context.Background()); err != nil {
		p.logger.Error("clickhouse ping fail", zap.Error(err))
	} else {
		p.logger.Info("clickhouse init success", zap.Any("addr", addr))
	}
	return nil
}

func (p *Processor) Handle(event bi.EventType, userId, distinct string, properties []byte) error {
	if err := p.deliver(event.String(), properties); err != nil {
		return err
	}
	return nil
}

func (p *Processor) deliver(event string, data []byte) error {
	if p.conn == nil {
		return errors.New("clickhouse conn nil")
	}

	var params = map[string]any{}
	if err := json.Unmarshal(data, &params); err != nil {
		return err
	}

	if isExist, err := p.existsTable(event); err != nil {
		return err
	} else if !isExist {
		if er := p.createTable(event, params); er != nil {
			return er
		}
	}

	if err := p.insertData(event, params); err != nil {
		p.logger.Error("clickhouse insert fail", zap.Error(err))
		return err
	}

	return nil
}

// https://clickhouse.com/docs/en/sql-reference/statements/create/table
// create table if table not exists
func (p *Processor) createTable(table string, params map[string]any) error {
	if p.conn == nil {
		return errors.New("clickhouse conn nil")
	}
	if table == "" {
		return errors.New("event empty")
	}
	if len(params) == 0 {
		return errors.New("params empty")
	}

	sql := "CREATE TABLE IF NOT EXISTS " + table + " ("
	for k, v := range params {
		sql += k + " " + p.getType(v) + ","
	}
	sql = strings.Trim(sql, ",") + ") ENGINE = MergeTree() ORDER BY ("
	for k := range params {
		sql += k + ","
	}
	sql = strings.Trim(sql, ",") + ")"

	if err := p.conn.Exec(context.Background(), sql); err != nil {
		return err
	}
	return nil
}

func (p *Processor) getType(v any) string {
	switch v.(type) {
	case int:
		return "Int32"
	case int64:
		return "Int64"
	case float64:
		return "Float64"
	case string:
		return "String"
	case time.Time:
		return "DateTime"
	default:
		return "String"
	}
}

func (p *Processor) existsTable(table string) (bool, error) {
	if p.conn == nil {
		return false, errors.New("clickhouse conn nil")
	}
	sql := "EXISTS TABLE " + table
	if rows, err := p.conn.Query(context.Background(), sql); err != nil {
		return false, err
	} else if rows.Next() {
		var exists bool
		if err := rows.Scan(&exists); err != nil {
			return false, err
		}
		return exists, nil
	}
	return false, nil
}

func (p *Processor) insertData(table string, params map[string]any) (err error) {
	if table == "" {
		return errors.New("event empty")
	}
	if len(params) == 0 {
		return errors.New("params empty")
	}
	var argv []any
	sql := "INSERT INTO " + table + " ("
	for k, v := range params {
		sql += k + ","
		argv = append(argv, v)
	}
	sql = strings.Trim(sql, ",") + ") VALUES ("
	for range params {
		sql += "?,"
	}
	sql = strings.Trim(sql, ",") + ")"

	if batch, er := p.conn.PrepareBatch(context.TODO(), sql); er != nil {
		return er
	} else if err = batch.Append(argv...); err != nil {
		return err
	} else if err = batch.Send(); err != nil {
		return err
	}
	return nil
}
