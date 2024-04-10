package shared

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLpaInitMarshalJSON(t *testing.T) {
	expected := `{
"lpaType":"",
"donor":{"uid":"","firstNames":"","lastName":"","address":{"line1":"","town":"","country":""},"dateOfBirth":"","email":""},
"attorneys":null,
"certificateProvider":{"uid":"","firstNames":"","lastName":"","address":{"line1":"","town":"","country":""},"email":"","phone":"","channel":""},
"signedAt":"0001-01-01T00:00:00Z"
}`

	data, _ := json.Marshal(LpaInit{})
	assert.JSONEq(t, expected, string(data))
}
