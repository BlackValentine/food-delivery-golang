package appctx

import (
	"server/skio"

	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	SecretKey() string
	GetRealtimeEngine() skio.RealtimeEngine
}

type appCtx struct {
	db        *gorm.DB
	secretKey string
	rtEngine  skio.RealtimeEngine
}

func NewAppContext(db *gorm.DB, secretKey string) *appCtx {
	return &appCtx{db: db, secretKey: secretKey}
}

func (ctx *appCtx) GetMainDBConnection() *gorm.DB {
	return ctx.db
}

func (ctx *appCtx) SecretKey() string {
	return ctx.secretKey
}

func (ctx *appCtx) GetRealtimeEngine() skio.RealtimeEngine {
	return ctx.rtEngine
}

func (ctx *appCtx) SetRealTimeEngine(rtEngine skio.RealtimeEngine) {
	ctx.rtEngine = rtEngine
}
