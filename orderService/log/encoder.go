package log

import (
	"strings"

	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
)

type EscapeSeqJSONEncoder struct {
	zapcore.Encoder
}

func (enc *EscapeSeqJSONEncoder) EncodeEntry(entry zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	encodedBytes, err := enc.Encoder.EncodeEntry(entry, fields)
	if err != nil {
		return nil, err
	}

	encodedBytesContent := string(encodedBytes.Bytes())

	modifiedBuffer := buffer.NewPool().Get()

	encodedBytesContent = strings.ReplaceAll(encodedBytesContent, "\\n", "\n")
	encodedBytesContent = strings.ReplaceAll(encodedBytesContent, "\\t", "\t")

	_, err = modifiedBuffer.Write([]byte(encodedBytesContent))
	if err != nil {
		return nil, err
	}

	return modifiedBuffer, nil
}
