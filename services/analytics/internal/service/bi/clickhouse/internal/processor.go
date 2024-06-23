package internal

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
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
	// 连接ClickHouse
	if addr != "" {
		if conn, err := clickhouse.Open(&clickhouse.Options{
			Addr: strings.Split(addr, ","),
			Auth: clickhouse.Auth{
				Database: dbname,
				Username: uname,
				Password: passwd,
			},
		}); err != nil {
			p.logger.Error("clickhouse init fail", zap.Error(err))
		} else {
			p.conn = conn
			p.logger.Info("clickhouse init", zap.Any("addr", addr), zap.Any("ping", p.conn.Ping(context.TODO())))
		}
	}
	return nil
}

func (p *Processor) Handle(event bi.EventType, properties []byte) error {
	if err := p.deliver(event.String(), properties); err != nil {
		return err
	}
	return nil
}

func (p *Processor) deliver(event string, data []byte) (err error) {

	var params = map[string]any{}
	if err = json.Unmarshal(data, &params); err != nil {
		return
	}
	params["event"] = event
	if data, err = json.Marshal(params); err != nil {
		return
	}

	if er := p.writeLogs(data); er != nil {
		p.logger.Error("clickhouselogs fail", zap.Error(er))
	}

	if p.conn != nil {
		btime := time.Now()
		delete(params, "event")
		if er := p.insertData(event, params); er != nil {
			p.logger.Error("clickhouse insert fail", zap.Error(er), zap.Any("time", time.Since(btime).Milliseconds()))
			return er
		}
	}

	return
}

func (p *Processor) insertData(table string, params map[string]any) (err error) {

	if table == "" {
		return errors.New("event empty")
	}
	if len(params) == 0 {
		return errors.New("params empty")
	}
	var argv []any
	var sql = fmt.Sprintf("INSERT INTO %s (", table)
	for k, v := range params {
		sql += k + ","
		argv = append(argv, v)
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

func (p *Processor) writeLogs(data []byte) (err error) {
	fileName := filepath.Join(p.rootPath, fmt.Sprintf("clickhouse_%s.log", time.Now().Format("2006-01-02")))
	if err = os.MkdirAll(filepath.Dir(fileName), os.ModePerm); err != nil {
		return fmt.Errorf("mkdir fail: %s - %w", fileName, err)
	}
	var file *os.File
	if file, err = os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644); err != nil {
		return fmt.Errorf("open fail fail: %s - %w", fileName, err)
	}
	defer file.Close()
	if _, err = file.WriteString(string(data) + "\n"); err != nil {
		return fmt.Errorf("write fail: %s - %s - %w", fileName, string(data), err)
	}
	return
}
