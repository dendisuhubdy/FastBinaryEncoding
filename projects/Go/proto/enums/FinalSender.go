// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: enums.fbe
// Version: 1.2.0.0

package enums

import "errors"
import "../fbe"

// Workaround for Go unused imports issue
var _ = fbe.Version

// Fast Binary Encoding enums final sender
type FinalSender struct {
    *fbe.Sender
    enumsModel *EnumsFinalModel
}

// Create a new enums final sender with an empty buffer
func NewFinalSender() *FinalSender {
    return NewFinalSenderWithBuffer(fbe.NewEmptyBuffer())
}

// Create a new enums final sender with the given buffer
func NewFinalSenderWithBuffer(buffer *fbe.Buffer) *FinalSender {
    return &FinalSender{
        fbe.NewSender(buffer, true),
        NewEnumsFinalModel(buffer),
    }
}

// Sender models accessors

func (s *FinalSender) EnumsModel() *EnumsFinalModel { return s.enumsModel }

// Send methods

func (s *FinalSender) Send(value interface{}) (int, error) {
    switch value := value.(type) {
    case *Enums:
        return s.SendEnums(value)
    }
    return 0, nil
}

func (s *FinalSender) SendEnums(value *Enums) (int, error) {
    // Serialize the value into the FBE stream
    serialized, err := s.enumsModel.Serialize(value)
    if serialized <= 0 {
        return 0, errors.New("enums.Enums serialization failed")
    }
    if err != nil {
        return 0, err
    }
    if !s.enumsModel.Verify() {
        return 0, errors.New("enums.Enums validation failed")
    }

    // Log the value
    if s.Logging() {
        message := value.String()
        s.HandlerOnSendLog.OnSendLog(message)
    }

    // Send the serialized value
    return s.SendSerialized(serialized)
}
