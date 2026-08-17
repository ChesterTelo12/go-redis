package redis

import (
	"context"
	"time"
)

func (c *PubSub) healthCheck() error {
	cn, err := c.conn()
	if err != nil {
		return err
	}

	// Enforce a deadline for the health check PING
	deadline := time.Now().Add(5 * time.Second)
	cn.SetDeadline(deadline)

	err = c.writeCmd(cn, "PING")
	if err != nil {
		cn.Close()
		return err
	}

	// Read the PONG response
	_, err = c.receive(cn)
	if err != nil {
		cn.Close()
		return err
	}

	// Reset deadline
	cn.SetDeadline(time.Time{})
	return nil
}