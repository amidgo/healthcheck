package healthcheck_test

import (
	"testing"

	healthcheckmocks "github.com/amidgo/healthcheck/mocks"
)

type HandlerTest struct {
	CaseName    string
	ServiceName string

	Mock func(p *healthcheckmocks.MockPinger)

	ExpectedError error
}

func (h *HandlerTest) Name() string {
	return h.CaseName
}

func (h *HandlerTest) Test(t *testing.T) {

}
