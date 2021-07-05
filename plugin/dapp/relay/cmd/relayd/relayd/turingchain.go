// Copyright Turing Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package relayd

import (
	"errors"
	"fmt"
	"io"
	"time"

	log "github.com/turingchain2020/turingchain/common/log/log15"
	"github.com/turingchain2020/turingchain/types"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// Client33 to connect with turingchain
type Client33 struct {
	config     *Turingchain
	isSyncing  bool
	isClosed   bool
	lastHeight int64
	types.TuringchainClient
	closer io.Closer
}

// NewClient33 new client instance
func NewClient33(cfg *Turingchain) *Client33 {
	conn, err := grpc.Dial(cfg.Host, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := types.NewTuringchainClient(conn)
	c := &Client33{
		config:        cfg,
		closer:        conn,
		TuringchainClient: client,
	}
	return c
}

func (c *Client33) heartbeat(ctx context.Context) {
	reconnectAttempts := c.config.ReconnectAttempts
out:
	for {
		log.Info("turingchain heartbeat.......")
		select {
		case <-ctx.Done():
			break out

		case <-time.After(time.Second * 60):
			err := c.ping(ctx)
			if err != nil {
				log.Error("heartbeat", "heartbeat turingchain error: ", err.Error(), "reconnectAttempts: ", reconnectAttempts)
				c.autoReconnect(ctx)
				reconnectAttempts--
			} else {
				reconnectAttempts = c.config.ReconnectAttempts
			}
			// TODO
			if reconnectAttempts <= 0 {
				panic(fmt.Errorf("reconnectAttempts <= %d", reconnectAttempts))
			}
		}
	}
}

// Start begin heartbeat to turingchain
func (c *Client33) Start(ctx context.Context) {
	go c.heartbeat(ctx)
}

func (c *Client33) ping(ctx context.Context) error {
	lastHeader, err := c.GetLastHeader(ctx, &types.ReqNil{})
	if err != nil {
		c.isClosed = false
		return err
	}

	c.isClosed = true
	c.lastHeight = lastHeader.Height
	log.Info("ping", "lastHeight:", c.lastHeight)
	isSync, err := c.IsSync(ctx, &types.ReqNil{})
	if err != nil {
		return err
	}

	if !isSync.IsOk {
		c.isSyncing = !isSync.IsOk
		log.Warn(fmt.Sprintf("node is syncing： %s", isSync.String()))
	}
	c.isSyncing = false
	return nil
}

func (c *Client33) autoReconnect(ctx context.Context) {
	if c.isClosed && !c.config.DisableAutoReconnect {
		c.closer.Close()
		conn, err := grpc.Dial(c.config.Host, grpc.WithInsecure())
		if err != nil {
			panic(err)
		}

		client := types.NewTuringchainClient(conn)
		c.closer = conn
		c.TuringchainClient = client
		c.isClosed = true
		c.Start(ctx)
	}
}

// SendTransaction send tx to turingchain
func (c *Client33) SendTransaction(ctx context.Context, in *types.Transaction) (*types.Reply, error) {
	if c.isSyncing {
		return nil, errors.New("node is syncing")
	}
	return c.TuringchainClient.SendTransaction(ctx, in)
}

// Close turingchain close
func (c *Client33) Close() error {
	return c.closer.Close()
}
