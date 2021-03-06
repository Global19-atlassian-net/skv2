// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/skv2/api/core/v1/core.proto

package v1

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/golang/protobuf/proto"
	equality "github.com/solo-io/protoc-gen-ext/pkg/equality"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = bytes.Compare
	_ = strings.Compare
	_ = equality.Equalizer(nil)
	_ = proto.Message(nil)
)

// Equal function
func (m *ObjectRef) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ObjectRef)
	if !ok {
		that2, ok := that.(ObjectRef)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetName(), target.GetName()) != 0 {
		return false
	}

	if strings.Compare(m.GetNamespace(), target.GetNamespace()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *ClusterObjectRef) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ClusterObjectRef)
	if !ok {
		that2, ok := that.(ClusterObjectRef)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetName(), target.GetName()) != 0 {
		return false
	}

	if strings.Compare(m.GetNamespace(), target.GetNamespace()) != 0 {
		return false
	}

	if strings.Compare(m.GetClusterName(), target.GetClusterName()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *TypedObjectRef) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*TypedObjectRef)
	if !ok {
		that2, ok := that.(TypedObjectRef)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetApiGroup()).(equality.Equalizer); ok {
		if !h.Equal(target.GetApiGroup()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetApiGroup(), target.GetApiGroup()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetKind()).(equality.Equalizer); ok {
		if !h.Equal(target.GetKind()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetKind(), target.GetKind()) {
			return false
		}
	}

	if strings.Compare(m.GetName(), target.GetName()) != 0 {
		return false
	}

	if strings.Compare(m.GetNamespace(), target.GetNamespace()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *TypedClusterObjectRef) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*TypedClusterObjectRef)
	if !ok {
		that2, ok := that.(TypedClusterObjectRef)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetApiGroup()).(equality.Equalizer); ok {
		if !h.Equal(target.GetApiGroup()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetApiGroup(), target.GetApiGroup()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetKind()).(equality.Equalizer); ok {
		if !h.Equal(target.GetKind()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetKind(), target.GetKind()) {
			return false
		}
	}

	if strings.Compare(m.GetName(), target.GetName()) != 0 {
		return false
	}

	if strings.Compare(m.GetNamespace(), target.GetNamespace()) != 0 {
		return false
	}

	if strings.Compare(m.GetClusterName(), target.GetClusterName()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *Status) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Status)
	if !ok {
		that2, ok := that.(Status)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if m.GetState() != target.GetState() {
		return false
	}

	if strings.Compare(m.GetMessage(), target.GetMessage()) != 0 {
		return false
	}

	if m.GetObservedGeneration() != target.GetObservedGeneration() {
		return false
	}

	if h, ok := interface{}(m.GetProcessingTime()).(equality.Equalizer); ok {
		if !h.Equal(target.GetProcessingTime()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetProcessingTime(), target.GetProcessingTime()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetOwner()).(equality.Equalizer); ok {
		if !h.Equal(target.GetOwner()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetOwner(), target.GetOwner()) {
			return false
		}
	}

	return true
}
