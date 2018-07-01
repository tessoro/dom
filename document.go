package dom

func GetDocument() *Document {
	doc := global.Get("document")
	if !IsValid(doc) {
		return nil
	}
	return &Document{NodeBase{doc}}
}

var _ Node = (*Document)(nil)

type Document struct {
	NodeBase
}

func (d *Document) CreateElement(tag string) *Element {
	v := d.v.Call("createElement", tag)
	return AsElement(v)
}
func (d *Document) CreateElementNS(ns string, tag string) *Element {
	v := d.v.Call("createElementNS", ns, tag)
	return AsElement(v)
}
func (d *Document) GetElementsByTagName(tag string) NodeList {
	v := d.v.Call("getElementsByTagName", tag)
	return AsNodeList(v)
}
func (d *Document) QuerySelector(qu string) *Element {
	v := d.v.Call("querySelector", qu)
	return AsElement(v)
}
func (d *Document) QuerySelectorAll(qu string) NodeList {
	v := d.v.Call("querySelectorAll", qu)
	return AsNodeList(v)
}