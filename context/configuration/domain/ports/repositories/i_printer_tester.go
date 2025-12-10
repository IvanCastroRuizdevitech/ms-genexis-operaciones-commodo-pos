package irepositories

import "context"

// IPrinterTester defines an outbound port to send a test ticket to the printer microservice.
type IPrinterTester interface {
	PrintTest(ctx context.Context, host string, port int, template []string) error
}
