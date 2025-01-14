package hypurrutils

import (
	"fmt"

	hp "github.com/s1dorfake/HypurrFun/pb"
)

// NewStaticClient creates a new gRPC client to communicate with Hypurr.fun's static servers
func NewStaticClient(userAgent string) (hp.StaticClient, error) {
	conn, err := createGRPCConnection(userAgent)
	if err != nil {
		return nil, fmt.Errorf("createGRPCConnection error: %w", err)
	}
	return hp.NewStaticClient(conn), nil
}

// NewTelegramClient creates a new gRPC client to communicate with Hypurr.fun's telegram servers
func NewTelegramClient(userAgent string) (hp.TelegramClient, error) {
	conn, err := createGRPCConnection(userAgent)
	if err != nil {
		return nil, fmt.Errorf("createGRPCConnection error: %w", err)
	}
	return hp.NewTelegramClient(conn), nil
}
