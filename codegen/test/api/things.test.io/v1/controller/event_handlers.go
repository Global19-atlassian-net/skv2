// Definitions for the Kubernetes Controllers
package controller

import (
	"context"

	things_test_io_v1 "github.com/solo-io/skv2/codegen/test/api/things.test.io/v1"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/events"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Handle events for the Paint Resource
// DEPRECATED: Prefer reconciler pattern.
type PaintEventHandler interface {
	CreatePaint(obj *things_test_io_v1.Paint) error
	UpdatePaint(old, new *things_test_io_v1.Paint) error
	DeletePaint(obj *things_test_io_v1.Paint) error
	GenericPaint(obj *things_test_io_v1.Paint) error
}

type PaintEventHandlerFuncs struct {
	OnCreate  func(obj *things_test_io_v1.Paint) error
	OnUpdate  func(old, new *things_test_io_v1.Paint) error
	OnDelete  func(obj *things_test_io_v1.Paint) error
	OnGeneric func(obj *things_test_io_v1.Paint) error
}

func (f *PaintEventHandlerFuncs) CreatePaint(obj *things_test_io_v1.Paint) error {
	if f.OnCreate == nil {
		return nil
	}
	return f.OnCreate(obj)
}

func (f *PaintEventHandlerFuncs) DeletePaint(obj *things_test_io_v1.Paint) error {
	if f.OnDelete == nil {
		return nil
	}
	return f.OnDelete(obj)
}

func (f *PaintEventHandlerFuncs) UpdatePaint(objOld, objNew *things_test_io_v1.Paint) error {
	if f.OnUpdate == nil {
		return nil
	}
	return f.OnUpdate(objOld, objNew)
}

func (f *PaintEventHandlerFuncs) GenericPaint(obj *things_test_io_v1.Paint) error {
	if f.OnGeneric == nil {
		return nil
	}
	return f.OnGeneric(obj)
}

type PaintEventWatcher interface {
	AddEventHandler(ctx context.Context, h PaintEventHandler, predicates ...predicate.Predicate) error
}

type paintEventWatcher struct {
	watcher events.EventWatcher
}

func NewPaintEventWatcher(name string, mgr manager.Manager) PaintEventWatcher {
	return &paintEventWatcher{
		watcher: events.NewWatcher(name, mgr, &things_test_io_v1.Paint{}),
	}
}

func (c *paintEventWatcher) AddEventHandler(ctx context.Context, h PaintEventHandler, predicates ...predicate.Predicate) error {
	handler := genericPaintHandler{handler: h}
	if err := c.watcher.Watch(ctx, handler, predicates...); err != nil {
		return err
	}
	return nil
}

// genericPaintHandler implements a generic events.EventHandler
type genericPaintHandler struct {
	handler PaintEventHandler
}

func (h genericPaintHandler) Create(object runtime.Object) error {
	obj, ok := object.(*things_test_io_v1.Paint)
	if !ok {
		return errors.Errorf("internal error: Paint handler received event for %T", object)
	}
	return h.handler.CreatePaint(obj)
}

func (h genericPaintHandler) Delete(object runtime.Object) error {
	obj, ok := object.(*things_test_io_v1.Paint)
	if !ok {
		return errors.Errorf("internal error: Paint handler received event for %T", object)
	}
	return h.handler.DeletePaint(obj)
}

func (h genericPaintHandler) Update(old, new runtime.Object) error {
	objOld, ok := old.(*things_test_io_v1.Paint)
	if !ok {
		return errors.Errorf("internal error: Paint handler received event for %T", old)
	}
	objNew, ok := new.(*things_test_io_v1.Paint)
	if !ok {
		return errors.Errorf("internal error: Paint handler received event for %T", new)
	}
	return h.handler.UpdatePaint(objOld, objNew)
}

func (h genericPaintHandler) Generic(object runtime.Object) error {
	obj, ok := object.(*things_test_io_v1.Paint)
	if !ok {
		return errors.Errorf("internal error: Paint handler received event for %T", object)
	}
	return h.handler.GenericPaint(obj)
}
