package discord

import (
	"fmt"

	"github.com/Karitham/corde/components"
)

// rspErr is a response error sent to discord
// It responds with an ephemeral message to the user
type rspErr string

// InteractionRespData implements the response data interface
func (r rspErr) InteractionRespData() *components.InteractionRespData {
	return components.NewResp().Ephemeral().Content(string(r)).InteractionRespData()
}

// newErrf returns a new rspErr with the given format
func newErrf(format string, args ...any) rspErr {
	return rspErr(fmt.Sprintf(format, args...))
}
