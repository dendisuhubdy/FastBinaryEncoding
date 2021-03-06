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

// Fast Binary Encoding array bool field model
type FieldModelArrayBool struct {
    // Field model buffer
    buffer *fbe.Buffer
    // Field model buffer offset
    offset int

    // Base field model value
    model *fbe.FieldModelBool
    // Array size
    size int
}

// Create a new array bool field model
func NewFieldModelArrayBool(model *fbe.FieldModelBool, buffer *fbe.Buffer, offset int, size int) *FieldModelArrayBool {
    fbeResult := FieldModelArrayBool{buffer: buffer, offset: offset}
    fbeResult.model = model
    fbeResult.size = size
    return &fbeResult
}

// Get the field size
func (fm *FieldModelArrayBool) FBESize() int { return fm.size * fm.model.FBESize() }

// Get the field extra size
func (fm *FieldModelArrayBool) FBEExtra() int { return 0 }

// Get the field offset
func (fm *FieldModelArrayBool) FBEOffset() int { return fm.offset }
// Set the field offset
func (fm *FieldModelArrayBool) SetFBEOffset(value int) { fm.offset = value }

// Shift the current field offset
func (fm *FieldModelArrayBool) FBEShift(size int) { fm.offset += size }
// Unshift the current field offset
func (fm *FieldModelArrayBool) FBEUnshift(size int) { fm.offset -= size }

// Get the array offset
func (fm *FieldModelArrayBool) Offset() int { return 0 }
// Get the array size
func (fm *FieldModelArrayBool) Size() int { return fm.size }

// Array index operator
func (fm *FieldModelArrayBool) GetItem(index int) (*fbe.FieldModelBool, error) {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return nil, errors.New("model is broken")
    }
    if index >= fm.size {
        return nil, errors.New("index is out of bounds")
    }

    fm.model.SetFBEOffset(fm.FBEOffset())
    fm.model.FBEShift(index * fm.model.FBESize())
    return fm.model, nil
}

// Check if the array is valid
func (fm *FieldModelArrayBool) Verify() bool {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return false
    }

    fm.model.SetFBEOffset(fm.FBEOffset())
    for i := 0; i < fm.size; i++ {
        if !fm.model.Verify() {
            return false
        }
        fm.model.FBEShift(fm.model.FBESize())
    }

    return true
}

// Get the array
func (fm *FieldModelArrayBool) Get(values []bool) error {
    values = values[:0]

    fbeModel, err := fm.GetItem(0)
    if err != nil {
        return err
    }

    for i := 0; i < fm.size; i++ {
        value, err := fbeModel.Get()
        if err == nil {
            return err
        }
        values = append(values, value)
        fbeModel.FBEShift(fbeModel.FBESize())
    }

    return nil
}

// Set the array
func (fm *FieldModelArrayBool) Set(values []bool) error {
    if (fm.buffer.Offset() + fm.FBEOffset() + fm.FBESize()) > fm.buffer.Size() {
        return errors.New("model is broken")
    }

    fbeModel, err := fm.GetItem(0)
    if err != nil {
        return err
    }

    size := len(values)
    if size > fm.size {
        size = fm.size
    }

    for i := 0; i < size; i++ {
        err := fbeModel.Set(values[i])
        if err == nil {
            return err
        }
        fbeModel.FBEShift(fbeModel.FBESize())
    }

    return nil
}
