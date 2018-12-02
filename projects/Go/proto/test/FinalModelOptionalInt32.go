// Automatically generated by the Fast Binary Encoding compiler, do not modify!
// https://github.com/chronoxor/FastBinaryEncoding
// Source: fbe
// Version: 1.1.0.0

package test

import "errors"
import "../fbe"
import "../proto"

// Workaround for Go unused imports issue
var _ = fbe.Version
var _ = proto.Version

// Fast Binary Encoding optional int32 final model
type FinalModelOptionalInt32 struct {
    // Final model buffer
    buffer *fbe.Buffer
    // Final model buffer offset
    offset int

    // Base final model value
    value *fbe.FinalModelInt32
}

// Create a new optional int32 final model
func NewFinalModelOptionalInt32(buffer *fbe.Buffer, offset int) *FinalModelOptionalInt32 {
    fbeResult := FinalModelOptionalInt32{buffer: buffer, offset: offset}
    fbeResult.value = fbe.NewFinalModelInt32(buffer, 0)
    return &fbeResult
}

// Get the optional final model value
func (fm *FinalModelOptionalInt32) Value() *fbe.FinalModelInt32 { return fm.value }

// Get the allocation size
func (fm *FinalModelOptionalInt32) FBEAllocationSize(fbeValue *int32) int {
    if fbeValue != nil {
        return 1 + fm.value.FBEAllocationSize(*fbeValue)
    } else {
        return 1
    }
}

// Get the field size
func (fm *FinalModelOptionalInt32) FBESize() int { return 0 }

// Get the field extra size
func (fm *FinalModelOptionalInt32) FBEExtra() int { return 0 }

// Get the field offset
func (fm *FinalModelOptionalInt32) FBEOffset() int { return fm.offset }
// Set the field offset
func (fm *FinalModelOptionalInt32) SetFBEOffset(value int) { fm.offset = value }

// Shift the current field offset
func (fm *FinalModelOptionalInt32) FBEShift(size int) { fm.offset += size }
// Unshift the current field offset
func (fm *FinalModelOptionalInt32) FBEUnshift(size int) { fm.offset -= size }

// Check if the object contains a value
func (fm *FinalModelOptionalInt32) HasValue() bool {
    if (fm.buffer.Offset() + fm.FBEOffset() + 1) > fm.buffer.Size() {
        return false
    }

    fbeHasValue := fbe.ReadUInt8(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset())
    return fbeHasValue != 0
}

// Check if the optional value is valid
func (fm *FinalModelOptionalInt32) Verify() int {
    if (fm.buffer.Offset() + fm.FBEOffset() + 1) > fm.buffer.Size() {
        return fbe.MaxInt
    }

    fbeHasValue := fbe.ReadUInt8(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset())
    if fbeHasValue == 0 {
        return 1
    }

    fm.buffer.Shift(fm.FBEOffset() + 1)
    fbeResult := fm.value.Verify()
    fm.buffer.Unshift(fm.FBEOffset() + 1)
    return fbeResult
}

// Get the optional value
func (fm *FinalModelOptionalInt32) Get() (*int32, int, error) {
    fbeResult := fbe.OptionalInt32(0)
    fbeSize, err := fm.GetValue(fbeResult)
    return fbeResult, fbeSize, err
}

// Get the optional value by the given pointer
func (fm *FinalModelOptionalInt32) GetValue(fbeValue *int32) (int, error) {
    if (fm.buffer.Offset() + fm.FBEOffset() + 1) > fm.buffer.Size() {
        return 0, errors.New("model is broken")
    }

    if !fm.HasValue() {
        return 1, nil
    }

    var fbeResult int
    var err error

    fm.buffer.Shift(fm.FBEOffset() + 1)
    *fbeValue, fbeResult, err = fm.value.Get()
    fm.buffer.Unshift(fm.FBEOffset() + 1)
    return fbeResult, err
}

// Set the optional value
func (fm *FinalModelOptionalInt32) Set(fbeValue *int32) (int, error) {
    if (fm.buffer.Offset() + fm.FBEOffset() + 1) > fm.buffer.Size() {
        return 0, errors.New("model is broken")
    }

    fbeHasValue := uint8(0)
    if fbeValue != nil {
        fbeHasValue = uint8(1)
    }
    fbe.WriteUInt8(fm.buffer.Data(), fm.buffer.Offset() + fm.FBEOffset(), fbeHasValue)
    if fbeHasValue == 0 {
        return 1, nil
    }

    fm.buffer.Shift(fm.FBEOffset() + 1)
    fbeResult, err := fm.value.Set(*fbeValue)
    fm.buffer.Unshift(fm.FBEOffset() + 1)
    return 1 + fbeResult, err
}
