package ovf

type Network struct {
	//XMLName xml.Name `xml:""`
	Name        string `xml:"http://schemas.dmtf.org/ovf/envelope/1 name,attr"`
	Description string `xml:"Description"`
}
