package model

import (
	"encoding/xml"
)

type CardPeek struct {
	XMLName xml.Name `xml:"cardpeek"`
	Text    string   `xml:",chardata"`
	Version string   `xml:"version"`
	Node    struct {
		Text string `xml:",chardata"`
		Attr []struct {
			Text string `xml:",chardata"`
			Name string `xml:"name,attr"`
		} `xml:"attr"`
		Node []struct {
			Text string `xml:",chardata"`
			Attr []struct {
				Text     string `xml:",chardata"`
				Name     string `xml:"name,attr"`
				Encoding string `xml:"encoding,attr"`
			} `xml:"attr"`
			Node []struct {
				Text string `xml:",chardata"`
				Attr []struct {
					Text string `xml:",chardata"`
					Name string `xml:"name,attr"`
				} `xml:"attr"`
				Node []struct {
					Text string `xml:",chardata"`
					Attr []struct {
						Text     string `xml:",chardata"`
						Name     string `xml:"name,attr"`
						Encoding string `xml:"encoding,attr"`
					} `xml:"attr"`
					Node []struct {
						Text string `xml:",chardata"`
						Attr []struct {
							Text     string `xml:",chardata"`
							Name     string `xml:"name,attr"`
							Encoding string `xml:"encoding,attr"`
						} `xml:"attr"`
						Node []struct {
							Text string `xml:",chardata"`
							Attr []struct {
								Text     string `xml:",chardata"`
								Name     string `xml:"name,attr"`
								Encoding string `xml:"encoding,attr"`
							} `xml:"attr"`
							Node []struct {
								Text string `xml:",chardata"`
								Attr []struct {
									Text     string `xml:",chardata"`
									Name     string `xml:"name,attr"`
									Encoding string `xml:"encoding,attr"`
								} `xml:"attr"`
								Node struct {
									Text string `xml:",chardata"`
									Attr []struct {
										Text     string `xml:",chardata"`
										Name     string `xml:"name,attr"`
										Encoding string `xml:"encoding,attr"`
									} `xml:"attr"`
								} `xml:"node"`
							} `xml:"node"`
						} `xml:"node"`
					} `xml:"node"`
				} `xml:"node"`
			} `xml:"node"`
		} `xml:"node"`
	} `xml:"node"`
}
