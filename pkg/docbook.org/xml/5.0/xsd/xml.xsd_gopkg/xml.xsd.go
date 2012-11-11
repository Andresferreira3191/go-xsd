//	Auto-generated by the "go-xsd" package located at:
//		github.com/metaleap/go-xsd
//	Comments on types and fields (if any) are from the XSD file located at:
//		docbook.org/xml/5.0/xsd/xml.xsd
package gopkg_DocbookOrgXml50XsdXmlXsd

import (
	xsdt "github.com/metaleap/go-xsd/types"
)

type XsdGoPkgHasCdata struct { CombinedCharDatas string `xml:",chardata"` }

type XsdGoPkgHasAttr_Id struct {
	Id xsdt.Id `xml:"http://www.w3.org/XML/1998/namespace id,attr"`
}

type XsdGoPkgHasAttr_Lang struct {
	Lang xsdt.String `xml:"http://www.w3.org/XML/1998/namespace lang,attr"`
}

type XsdGoPkgHasAttr_Base struct {
	Base xsdt.String `xml:"http://www.w3.org/XML/1998/namespace base,attr"`
}

type TxsdSpace xsdt.Token

//	Since TxsdSpace is just a simple String type, this merely sets the current value from the specified string.
func (me *TxsdSpace) SetFromString (s string) { (*xsdt.Token)(me).SetFromString(s) }

//	Since TxsdSpace is just a simple String type, this merely returns the current string value.
func (me TxsdSpace) String () string { return xsdt.Token(me).String() }

//	This convenience method just performs a simple type conversion to TxsdSpace's alias type xsdt.Token
func (me TxsdSpace) ToXsdtToken () xsdt.Token { return xsdt.Token(me) }

//	Returns true if the value of this enumerated TxsdSpace is "preserve".
func (me TxsdSpace) IsPreserve () bool { return me == "preserve" }

type XsdGoPkgHasAttr_Space struct {
	Space TxsdSpace `xml:"http://www.w3.org/XML/1998/namespace space,attr"`
}
