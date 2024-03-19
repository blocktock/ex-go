package util

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/blocktock/go-pkg/log"
	"go.uber.org/zap"
	"io"
	"net/http"
)

func SendRequest(ctx context.Context, method string, url string, data interface{}) (res []byte, err error) {

	logger := log.LoggerWithContext(ctx).With(
		zap.Any("params", data),
		zap.String("url", url),
	)

	requestByte, err := json.Marshal(data)
	if err != nil {
		logger.Error("SendRequest json.Marshal failed",
			zap.Error(err),
		)
		return nil, err
	}

	req, err := http.NewRequest(method, url, bytes.NewReader(requestByte))
	if err != nil {
		logger.Error("SendRequest http.NewRequest failed",
			zap.Error(err),
		)
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		logger.Error("SendRequest client.Do failed",
			zap.Error(err),
		)
		return nil, err
	}

	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	logger.Info("SendRequest",
		zap.Any("response", string(responseBody)),
		zap.Error(err),
	)
	return responseBody, err
}
