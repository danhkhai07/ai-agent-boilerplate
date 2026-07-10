package mcp

import (
	"context"
	"errors"
	"time"

	mcp_sdk "github.com/modelcontextprotocol/go-sdk/mcp"
)

type checkTimeInput struct {
	City string `json:"city" jsonschema:"hanoi, nyc or beijing; default if empty: hanoi"`
}

type checkTimeOutput struct {
	Time time.Time `json:"time" jsonschema:"time output of hanoi, nyc or beijing"`
}

func checkTime(ctx context.Context, request *mcp_sdk.CallToolRequest, input checkTimeInput) (
	*mcp_sdk.CallToolResult,
	checkTimeOutput,
	error,
) {
	now := time.Now()
	var loc *time.Location
	var locString string

	switch input.City {
	case "hanoi":
		locString = "Asia/Ho_Chi_Minh"
	case "nyc":
		locString = "America/New_York"
	case "beijing":
		locString = "Asia/Shanghai"
	case "":
		locString = "Asia/Ho_Chi_Minh"
	default:
		return nil, checkTimeOutput{}, errors.New("undetected city")
	}

	loc, err := time.LoadLocation(locString)
	if err != nil {
		return nil, checkTimeOutput{}, err
	}
	now = now.In(loc)

	return nil, checkTimeOutput{ now }, nil
}
