package vcloud

type LinkPredicate func(*Link) bool

func byTypeAndRel(tpe, rel string) LinkPredicate {
	if rel == "" {
		rel = RelDown
	}
	return func(lnk *Link) bool {
		return lnk != nil && lnk.Type == tpe && lnk.Rel == rel
	}
}

func byNameTypeAndRel(nme, tpe, rel string) LinkPredicate {
	tpePred := byTypeAndRel(tpe, rel)
	return func(lnk *Link) bool {
		return tpePred(lnk) && lnk.Name == nme
	}
}

type LinkList []*Link

func (l LinkList) Find(predicate LinkPredicate) *Link {
	for _, lnk := range l {
		if predicate(lnk) {
			return lnk
		}
	}
	return nil
}

func (l LinkList) ForType(tpe, rel string) *Link {
	return l.Find(byTypeAndRel(tpe, rel))
}

func (l LinkList) ForName(name, tpe, rel string) *Link {
	return l.Find(byNameTypeAndRel(name, tpe, rel))
}

// Link extends reference type by adding relation attribute. Defines a hyper-link with a relationship, hyper-link reference, and an optional MIME type.
// Type: LinkType
// Namespace: http://www.vmware.com/vcloud/v1.5
// Description: Extends reference type by adding relation attribute. Defines a hyper-link with a relationship, hyper-link reference, and an optional MIME type.
// Since: 0.9
type Link struct {
	HREF string `xml:"href,attr"`
	ID   string `xml:"id,attr,omitempty"`
	Type string `xml:"type,attr,omitempty"`
	Name string `xml:"name,attr,omitempty"`
	Rel  string `xml:"rel,attr"`
}
