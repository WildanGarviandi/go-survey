package controllers

import (
	"encoding/json"
	"fmt"

	"strings"
	"time"

	"github.com/kataras/iris"
	"github.com/satori/go.uuid"
)

type requestLog struct {
	Timestamp     time.Time   `json:"timestamp"`
	CorrelationID string      `json:"correlationId"`
	Method        string      `json:"method"`
	URL           string      `json:"url"`
	Status        int         `json:"status"`
	ResponseTime  float64     `json:"responseTime"`
	RequestBody   interface{} `json:"requestBody"`
}

// LogHandler - log request information
var LogHandler = func(ctx iris.Context) {
	cid := ctx.GetHeader("X-Correlation-Id")
	startTime := time.Now()
	if len(strings.TrimSpace(cid)) == 0 {
		uid, _ := uuid.NewV4()
		cid = fmt.Sprintf("%s", uid)
		ctx.Request().Header.Set("X-Correlation-Id", cid)
	}
	var req interface{}

	err := ctx.ReadJSON(&req)

	bodyJSON, err := json.Marshal(req)

	rl := requestLog{
		Timestamp:     startTime,
		CorrelationID: cid,
		Method:        ctx.Method(),
		URL:           fmt.Sprintf("%s", ctx.Request().URL),
		Status:        ctx.GetStatusCode(),
		ResponseTime:  time.Now().Sub(startTime).Seconds(),
		RequestBody:   string(bodyJSON),
	}
	b, err := json.Marshal(&rl)
	if err == nil {
		fmt.Printf("%+v\n", string(b))
	}
	ctx.Next()
}
