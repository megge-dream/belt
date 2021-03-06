package belt

import (
	"context"
)

type Event interface{}

type Item interface {
	Event() Event
	Handler() Handler
	MakeChild(Handler) (Item, error)
	Context() context.Context
	SetContext(ctx context.Context)
}

type Handler interface {
	Handle(context.Context) ([]Handler, error)
}

type Middleware interface {
	Handle(context.Context, Item) ([]Handler, error)
}

type Slot interface {
	Middleware() Middleware
	Reset(state Middleware)
	AddItem(Item) error
	RemoveItem(Item) error
}

type Sorter interface {
	Sort(context.Context, Event) (Slot, Item, error)
}

type Worker interface {
	Work(ctx context.Context, items <-chan Event) error
}

type Canceler interface {
	Cancel()
}
