package osticket_test

import (
	"bytes"
	"testing"

	"github.com/codemonkeysoftware/osticket-go"
	"github.com/stretchr/testify/assert"
)

func TestMessageMarshalJSON(t *testing.T) {
	t.Run("includes content type", func(t *testing.T) {
		for _, cType := range []osticket.ContentType{osticket.ContentTypeHTML, osticket.ContentTypePlain} {
			m := osticket.Message{
				ContentType: cType,
			}
			b, err := m.MarshalJSON()
			assert.NoError(t, err)
			assert.Contains(t, string(b), cType)
		}
	})

	t.Run("includes data header, content type part, and body part", func(t *testing.T) {
		const (
			contentType = osticket.ContentTypePlain
			body        = "testbody"
		)
		m := osticket.Message{
			ContentType: contentType,
			Body:        body,
		}

		b, err := m.MarshalJSON()
		assert.NoError(t, err)
		b = bytes.ReplaceAll(b, []byte(`"`), []byte(``))

		parts := bytes.Split(b, []byte(":"))
		twoParts := assert.Len(t, parts, 2)
		assert.Equal(t, string(parts[0]), "data")

		if twoParts {
			parts = bytes.SplitN(parts[1], []byte(","), 2)
			if assert.Len(t, parts, 2) {
				assert.Equal(t, string(parts[1]), body)
			}
			assert.Equal(t, string(parts[0]), string(contentType))
		}
	})
}

func TestOptionalStateFunctions(t *testing.T) {
	result := osticket.Should()
	assert.True(t, *result)

	result = osticket.ShouldNot()
	assert.False(t, *result)
}
