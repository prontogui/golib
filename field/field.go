package field

import "github.com/prontogui/golib/key"

type Field interface {
	PrepareForUpdates(fkey key.FKey, pkey key.PKey, onset key.OnSetFunction)
	EgestValue() any
	IngestValue(value any) error
}
