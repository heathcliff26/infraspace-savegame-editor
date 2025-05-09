package fyne

// Declare conformity to [CanvasObject]
var _ CanvasObject = (*Container)(nil)

// Container is a [CanvasObject] that contains a collection of child objects.
// The layout of the children is set by the specified Layout.
type Container struct {
	size     Size     // The current size of the Container
	position Position // The current position of the Container
	Hidden   bool     // Is this Container hidden

	Layout  Layout         // The Layout algorithm for arranging child [CanvasObject]s
	Objects []CanvasObject // The set of [CanvasObject]s this container holds
}

// NewContainer returns a new [Container] instance holding the specified [CanvasObject]s.
//
// Deprecated: Use [fyne.io/fyne/v2/container.NewWithoutLayout] to create a container that uses manual layout.
func NewContainer(objects ...CanvasObject) *Container {
	return NewContainerWithoutLayout(objects...)
}

// NewContainerWithoutLayout returns a new [Container] instance holding the specified
// [CanvasObject]s that are manually arranged.
//
// Deprecated: Use [fyne.io/fyne/v2/container.NewWithoutLayout] instead.
func NewContainerWithoutLayout(objects ...CanvasObject) *Container {
	ret := &Container{
		Objects: objects,
	}

	ret.size = ret.MinSize()
	return ret
}

// NewContainerWithLayout returns a new [Container] instance holding the specified
// [CanvasObject]s which will be laid out according to the specified Layout.
//
// Deprecated: Use [fyne.io/fyne/v2/container.New] instead.
func NewContainerWithLayout(layout Layout, objects ...CanvasObject) *Container {
	ret := &Container{
		Objects: objects,
		Layout:  layout,
	}

	ret.size = layout.MinSize(objects)
	ret.layout()
	return ret
}

// Add appends the specified object to the items this container manages.
//
// Since: 1.4
func (c *Container) Add(add CanvasObject) {
	if add == nil {
		return
	}

	c.Objects = append(c.Objects, add)
	c.layout()
}

// AddObject adds another [CanvasObject] to the set this Container holds.
//
// Deprecated: Use [Container.Add] instead.
func (c *Container) AddObject(o CanvasObject) {
	c.Add(o)
}

// Hide sets this container, and all its children, to be not visible.
func (c *Container) Hide() {
	if c.Hidden {
		return
	}

	c.Hidden = true
	repaint(c)
}

// MinSize calculates the minimum size of c.
// This is delegated to the [Container.Layout], if specified, otherwise it will be calculated.
func (c *Container) MinSize() Size {
	if c.Layout != nil {
		return c.Layout.MinSize(c.Objects)
	}

	minSize := NewSize(1, 1)
	for _, child := range c.Objects {
		minSize = minSize.Max(child.MinSize())
	}

	return minSize
}

// Move the container (and all its children) to a new position, relative to its parent.
func (c *Container) Move(pos Position) {
	c.position = pos
	repaint(c)
}

// Position gets the current position of c relative to its parent.
func (c *Container) Position() Position {
	return c.position
}

// Refresh causes this object to be redrawn in its current state
func (c *Container) Refresh() {
	c.layout()

	for _, child := range c.Objects {
		child.Refresh()
	}

	// this is basically just canvas.Refresh(c) without the package loop
	o := CurrentApp().Driver().CanvasForObject(c)
	if o == nil {
		return
	}
	o.Refresh(c)
}

// Remove updates the contents of this container to no longer include the specified object.
// This method is not intended to be used inside a loop, to remove all the elements.
// It is much more efficient to call [Container.RemoveAll) instead.
func (c *Container) Remove(rem CanvasObject) {
	if len(c.Objects) == 0 {
		return
	}

	for i, o := range c.Objects {
		if o != rem {
			continue
		}
		copy(c.Objects[i:], c.Objects[i+1:])
		c.Objects[len(c.Objects)-1] = nil
		c.Objects = c.Objects[:len(c.Objects)-1]
		c.layout()
		return
	}
}

// RemoveAll updates the contents of this container to no longer include any objects.
//
// Since: 2.2
func (c *Container) RemoveAll() {
	c.Objects = nil
	c.layout()
}

// Resize sets a new size for c.
func (c *Container) Resize(size Size) {
	if c.size == size {
		return
	}

	c.size = size
	c.layout()
}

// Show sets this container, and all its children, to be visible.
func (c *Container) Show() {
	if !c.Hidden {
		return
	}

	c.Hidden = false
}

// Size returns the current size c.
func (c *Container) Size() Size {
	return c.size
}

// Visible returns true if the container is currently visible, false otherwise.
func (c *Container) Visible() bool {
	return !c.Hidden
}

func (c *Container) layout() {
	if c.Layout == nil {
		return
	}

	c.Layout.Layout(c.Objects, c.size)
}

// repaint instructs the containing canvas to redraw, even if nothing changed.
// This method is a duplicate of what is in `canvas/canvas.go` to avoid a dependency loop or public API.
func repaint(obj *Container) {
	app := CurrentApp()
	if app == nil || app.Driver() == nil {
		return
	}

	c := app.Driver().CanvasForObject(obj)
	if c != nil {
		if paint, ok := c.(interface{ SetDirty() }); ok {
			paint.SetDirty()
		}
	}
}
