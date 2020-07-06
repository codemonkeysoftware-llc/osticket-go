package osticket_test

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/codemonkeysoftware/osticket-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
func TestTicketMarshalJSON(t *testing.T) {
	t.Run("correctly marshals attachments", func(t *testing.T) {

		const (
			name     = "filename.jpg"
			mimeType = "image/jpeg"
			content  = "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVR42mMUbz3yHwAEfAJhtYFLfQAAAABJRU5ErkJggg=="
		)
		contentBytes, err := base64.StdEncoding.DecodeString(content)
		require.NoError(t, err)
		m := osticket.Attachment{
			Name:     name,
			MimeType: mimeType,
			Encoding: "base64",
			Data:     contentBytes,
		}

		b, err := m.MarshalJSON()
		assert.NoError(t, err)
		fmt.Printf("%s", b)
	})
}

func TestOptionalStateFunctions(t *testing.T) {
	result := osticket.Should()
	assert.True(t, *result)

	result = osticket.ShouldNot()
	assert.False(t, *result)
}
