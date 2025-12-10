package printer

import (
	"context"
	"fmt"
	"net"
	"time"
)

// Client handles TCP writes to a network printer.
type Client struct {
	Timeout time.Duration
}

// Send connects to host:port with the configured timeout and writes the payload bytes.
func (c *Client) Send(ctx context.Context, host string, port int, payload []byte) (int, error) {
	timeout := c.Timeout
	if timeout <= 0 {
		timeout = 3 * time.Second
	}

	ctxTimeout, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	address := fmt.Sprintf("%s:%d", host, port)
	dialer := &net.Dialer{}
	conn, err := dialer.DialContext(ctxTimeout, "tcp", address)
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	_ = conn.SetDeadline(time.Now().Add(timeout))

	total := 0
	for len(payload) > 0 {
		n, err := conn.Write(payload)
		total += n
		if err != nil {
			return total, err
		}
		payload = payload[n:]
	}

	return total, nil
}
