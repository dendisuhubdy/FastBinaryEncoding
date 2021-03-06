// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: test.fbe
// Version: 1.2.0.0

package test

import "errors"
import "../fbe"
import "../proto"

// Workaround for Go unused imports issue
var _ = fbe.Version
var _ = proto.Version

// Fast Binary Encoding test final sender
type FinalSender struct {
    *fbe.Sender
    protoSender *proto.FinalSender
    structSimpleModel *StructSimpleFinalModel
    structOptionalModel *StructOptionalFinalModel
    structNestedModel *StructNestedFinalModel
    structBytesModel *StructBytesFinalModel
    structArrayModel *StructArrayFinalModel
    structVectorModel *StructVectorFinalModel
    structListModel *StructListFinalModel
    structSetModel *StructSetFinalModel
    structMapModel *StructMapFinalModel
    structHashModel *StructHashFinalModel
    structHashExModel *StructHashExFinalModel
}

// Create a new test final sender with an empty buffer
func NewFinalSender() *FinalSender {
    return NewFinalSenderWithBuffer(fbe.NewEmptyBuffer())
}

// Create a new test final sender with the given buffer
func NewFinalSenderWithBuffer(buffer *fbe.Buffer) *FinalSender {
    return &FinalSender{
        fbe.NewSender(buffer, true),
        proto.NewFinalSenderWithBuffer(buffer),
        NewStructSimpleFinalModel(buffer),
        NewStructOptionalFinalModel(buffer),
        NewStructNestedFinalModel(buffer),
        NewStructBytesFinalModel(buffer),
        NewStructArrayFinalModel(buffer),
        NewStructVectorFinalModel(buffer),
        NewStructListFinalModel(buffer),
        NewStructSetFinalModel(buffer),
        NewStructMapFinalModel(buffer),
        NewStructHashFinalModel(buffer),
        NewStructHashExFinalModel(buffer),
    }
}

// Imported senders

func (s *FinalSender) ProtoSender() *proto.FinalSender { return s.protoSender }

// Sender models accessors

func (s *FinalSender) StructSimpleModel() *StructSimpleFinalModel { return s.structSimpleModel }
func (s *FinalSender) StructOptionalModel() *StructOptionalFinalModel { return s.structOptionalModel }
func (s *FinalSender) StructNestedModel() *StructNestedFinalModel { return s.structNestedModel }
func (s *FinalSender) StructBytesModel() *StructBytesFinalModel { return s.structBytesModel }
func (s *FinalSender) StructArrayModel() *StructArrayFinalModel { return s.structArrayModel }
func (s *FinalSender) StructVectorModel() *StructVectorFinalModel { return s.structVectorModel }
func (s *FinalSender) StructListModel() *StructListFinalModel { return s.structListModel }
func (s *FinalSender) StructSetModel() *StructSetFinalModel { return s.structSetModel }
func (s *FinalSender) StructMapModel() *StructMapFinalModel { return s.structMapModel }
func (s *FinalSender) StructHashModel() *StructHashFinalModel { return s.structHashModel }
func (s *FinalSender) StructHashExModel() *StructHashExFinalModel { return s.structHashExModel }

// Send methods

func (s *FinalSender) Send(value interface{}) (int, error) {
    switch value := value.(type) {
    case *StructSimple:
        return s.SendStructSimple(value)
    case *StructOptional:
        return s.SendStructOptional(value)
    case *StructNested:
        return s.SendStructNested(value)
    case *StructBytes:
        return s.SendStructBytes(value)
    case *StructArray:
        return s.SendStructArray(value)
    case *StructVector:
        return s.SendStructVector(value)
    case *StructList:
        return s.SendStructList(value)
    case *StructSet:
        return s.SendStructSet(value)
    case *StructMap:
        return s.SendStructMap(value)
    case *StructHash:
        return s.SendStructHash(value)
    case *StructHashEx:
        return s.SendStructHashEx(value)
    }
    if result, err := s.protoSender.Send(value); (result > 0) || (err != nil) {
        return result, err
    }
    return 0, nil
}

func (s *FinalSender) SendStructSimple(value *StructSimple) (int, error) {
    // Serialize the value into the FBE stream
    serialized, err := s.structSimpleModel.Serialize(value)
    if serialized <= 0 {
        return 0, errors.New("test.StructSimple serialization failed")
    }
    if err != nil {
        return 0, err
    }
    if !s.structSimpleModel.Verify() {
        return 0, errors.New("test.StructSimple validation failed")
    }

    // Log the value
    if s.Logging() {
        message := value.String()
        s.HandlerOnSendLog.OnSendLog(message)
    }

    // Send the serialized value
    return s.SendSerialized(serialized)
}

func (s *FinalSender) SendStructOptional(value *StructOptional) (int, error) {
    // Serialize the value into the FBE stream
    serialized, err := s.structOptionalModel.Serialize(value)
    if serialized <= 0 {
        return 0, errors.New("test.StructOptional serialization failed")
    }
    if err != nil {
        return 0, err
    }
    if !s.structOptionalModel.Verify() {
        return 0, errors.New("test.StructOptional validation failed")
    }

    // Log the value
    if s.Logging() {
        message := value.String()
        s.HandlerOnSendLog.OnSendLog(message)
    }

    // Send the serialized value
    return s.SendSerialized(serialized)
}

func (s *FinalSender) SendStructNested(value *StructNested) (int, error) {
    // Serialize the value into the FBE stream
    serialized, err := s.structNestedModel.Serialize(value)
    if serialized <= 0 {
        return 0, errors.New("test.StructNested serialization failed")
    }
    if err != nil {
        return 0, err
    }
    if !s.structNestedModel.Verify() {
        return 0, errors.New("test.StructNested validation failed")
    }

    // Log the value
    if s.Logging() {
        message := value.String()
        s.HandlerOnSendLog.OnSendLog(message)
    }

    // Send the serialized value
    return s.SendSerialized(serialized)
}

func (s *FinalSender) SendStructBytes(value *StructBytes) (int, error) {
    // Serialize the value into the FBE stream
    serialized, err := s.structBytesModel.Serialize(value)
    if serialized <= 0 {
        return 0, errors.New("test.StructBytes serialization failed")
    }
    if err != nil {
        return 0, err
    }
    if !s.structBytesModel.Verify() {
        return 0, errors.New("test.StructBytes validation failed")
    }

    // Log the value
    if s.Logging() {
        message := value.String()
        s.HandlerOnSendLog.OnSendLog(message)
    }

    // Send the serialized value
    return s.SendSerialized(serialized)
}

func (s *FinalSender) SendStructArray(value *StructArray) (int, error) {
    // Serialize the value into the FBE stream
    serialized, err := s.structArrayModel.Serialize(value)
    if serialized <= 0 {
        return 0, errors.New("test.StructArray serialization failed")
    }
    if err != nil {
        return 0, err
    }
    if !s.structArrayModel.Verify() {
        return 0, errors.New("test.StructArray validation failed")
    }

    // Log the value
    if s.Logging() {
        message := value.String()
        s.HandlerOnSendLog.OnSendLog(message)
    }

    // Send the serialized value
    return s.SendSerialized(serialized)
}

func (s *FinalSender) SendStructVector(value *StructVector) (int, error) {
    // Serialize the value into the FBE stream
    serialized, err := s.structVectorModel.Serialize(value)
    if serialized <= 0 {
        return 0, errors.New("test.StructVector serialization failed")
    }
    if err != nil {
        return 0, err
    }
    if !s.structVectorModel.Verify() {
        return 0, errors.New("test.StructVector validation failed")
    }

    // Log the value
    if s.Logging() {
        message := value.String()
        s.HandlerOnSendLog.OnSendLog(message)
    }

    // Send the serialized value
    return s.SendSerialized(serialized)
}

func (s *FinalSender) SendStructList(value *StructList) (int, error) {
    // Serialize the value into the FBE stream
    serialized, err := s.structListModel.Serialize(value)
    if serialized <= 0 {
        return 0, errors.New("test.StructList serialization failed")
    }
    if err != nil {
        return 0, err
    }
    if !s.structListModel.Verify() {
        return 0, errors.New("test.StructList validation failed")
    }

    // Log the value
    if s.Logging() {
        message := value.String()
        s.HandlerOnSendLog.OnSendLog(message)
    }

    // Send the serialized value
    return s.SendSerialized(serialized)
}

func (s *FinalSender) SendStructSet(value *StructSet) (int, error) {
    // Serialize the value into the FBE stream
    serialized, err := s.structSetModel.Serialize(value)
    if serialized <= 0 {
        return 0, errors.New("test.StructSet serialization failed")
    }
    if err != nil {
        return 0, err
    }
    if !s.structSetModel.Verify() {
        return 0, errors.New("test.StructSet validation failed")
    }

    // Log the value
    if s.Logging() {
        message := value.String()
        s.HandlerOnSendLog.OnSendLog(message)
    }

    // Send the serialized value
    return s.SendSerialized(serialized)
}

func (s *FinalSender) SendStructMap(value *StructMap) (int, error) {
    // Serialize the value into the FBE stream
    serialized, err := s.structMapModel.Serialize(value)
    if serialized <= 0 {
        return 0, errors.New("test.StructMap serialization failed")
    }
    if err != nil {
        return 0, err
    }
    if !s.structMapModel.Verify() {
        return 0, errors.New("test.StructMap validation failed")
    }

    // Log the value
    if s.Logging() {
        message := value.String()
        s.HandlerOnSendLog.OnSendLog(message)
    }

    // Send the serialized value
    return s.SendSerialized(serialized)
}

func (s *FinalSender) SendStructHash(value *StructHash) (int, error) {
    // Serialize the value into the FBE stream
    serialized, err := s.structHashModel.Serialize(value)
    if serialized <= 0 {
        return 0, errors.New("test.StructHash serialization failed")
    }
    if err != nil {
        return 0, err
    }
    if !s.structHashModel.Verify() {
        return 0, errors.New("test.StructHash validation failed")
    }

    // Log the value
    if s.Logging() {
        message := value.String()
        s.HandlerOnSendLog.OnSendLog(message)
    }

    // Send the serialized value
    return s.SendSerialized(serialized)
}

func (s *FinalSender) SendStructHashEx(value *StructHashEx) (int, error) {
    // Serialize the value into the FBE stream
    serialized, err := s.structHashExModel.Serialize(value)
    if serialized <= 0 {
        return 0, errors.New("test.StructHashEx serialization failed")
    }
    if err != nil {
        return 0, err
    }
    if !s.structHashExModel.Verify() {
        return 0, errors.New("test.StructHashEx validation failed")
    }

    // Log the value
    if s.Logging() {
        message := value.String()
        s.HandlerOnSendLog.OnSendLog(message)
    }

    // Send the serialized value
    return s.SendSerialized(serialized)
}
