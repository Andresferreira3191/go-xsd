//	Auto-generated by the "go-xsd" package located at:
//		github.com/metaleap/go-xsd
//	Comments on types and fields (if any) are from the XSD file located at:
//		kbcafe.com/rss/atom.xsd.xml
package gopkg_KbcafeComRssAtomXsdXml

//	This version of the Atom schema is based on version 1.0 of the format specifications,
//	found here http://www.atomenabled.org/developers/syndication/atom-format-spec.php.
//	An Atom document may have two root elements, feed and entry, as defined in section 2.
import (
	xsdt "github.com/metaleap/go-xsd/types"
	xml "github.com/metaleap/go-xsd/pkg/www.w3.org/2001/03/xml.xsd_gopkg"
)

type XsdGoPkgHasCdata struct { CombinedCharDatas string `xml:",chardata"` }

//	Schema definition for an email address.
type TemailType xsdt.NormalizedString

//	Since TemailType is just a simple String type, this merely sets the current value from the specified string.
func (me *TemailType) SetFromString (s string) { (*xsdt.NormalizedString)(me).SetFromString(s) }

//	Since TemailType is just a simple String type, this merely returns the current string value.
func (me TemailType) String () string { return xsdt.NormalizedString(me).String() }

//	This convenience method just performs a simple type conversion to TemailType's alias type xsdt.NormalizedString
func (me TemailType) ToXsdtNormalizedString () xsdt.NormalizedString { return xsdt.NormalizedString(me) }

type XsdGoPkgHasAtts_CommonAttributes struct {
	xml.XsdGoPkgHasAttr_Base
	xml.XsdGoPkgHasAttr_Lang
}

type TxsdTextTypeType xsdt.Token

//	Since TxsdTextTypeType is just a simple String type, this merely sets the current value from the specified string.
func (me *TxsdTextTypeType) SetFromString (s string) { (*xsdt.Token)(me).SetFromString(s) }

//	Since TxsdTextTypeType is just a simple String type, this merely returns the current string value.
func (me TxsdTextTypeType) String () string { return xsdt.Token(me).String() }

//	This convenience method just performs a simple type conversion to TxsdTextTypeType's alias type xsdt.Token
func (me TxsdTextTypeType) ToXsdtToken () xsdt.Token { return xsdt.Token(me) }

//	Returns true if the value of this enumerated TxsdTextTypeType is "text".
func (me TxsdTextTypeType) IsText () bool { return me == "text" }

//	Returns true if the value of this enumerated TxsdTextTypeType is "html".
func (me TxsdTextTypeType) IsHtml () bool { return me == "html" }

//	Returns true if the value of this enumerated TxsdTextTypeType is "xhtml".
func (me TxsdTextTypeType) IsXhtml () bool { return me == "xhtml" }

type XsdGoPkgHasAttr_Type_TxsdTextTypeType_ struct {
	Type TxsdTextTypeType `xml:"http://www.w3.org/2005/Atom type,attr"`
}

//	The Atom text construct is defined in section 3.1 of the format spec.
type TtextType struct {
	XsdGoPkgHasCdata

	XsdGoPkgHasAtts_CommonAttributes

	XsdGoPkgHasAttr_Type_TxsdTextTypeType_

}

type XsdGoPkgHasElems_Name_XsdtString_ struct {
	Names []xsdt.String `xml:"http://www.w3.org/2005/Atom name"`
}

type XsdGoPkgHasElem_Name_XsdtString_ struct {
	Name xsdt.String `xml:"http://www.w3.org/2005/Atom name"`
}

type XsdGoPkgHasElem_Uri_TuriType_ struct {
	Uri *TuriType `xml:"http://www.w3.org/2005/Atom uri"`
}

type XsdGoPkgHasElems_Uri_TuriType_ struct {
	Uris []*TuriType `xml:"http://www.w3.org/2005/Atom uri"`
}

type XsdGoPkgHasElem_Email_TemailType_ struct {
	Email TemailType `xml:"http://www.w3.org/2005/Atom email"`
}

type XsdGoPkgHasElems_Email_TemailType_ struct {
	Emails []TemailType `xml:"http://www.w3.org/2005/Atom email"`
}

//	The Atom person construct is defined in section 3.2 of the format spec.
type TpersonType struct {
	XsdGoPkgHasElems_Name_XsdtString_

	XsdGoPkgHasElems_Uri_TuriType_

	XsdGoPkgHasElems_Email_TemailType_

	XsdGoPkgHasAtts_CommonAttributes

}

type XsdGoPkgHasElem_Author_TpersonType_ struct {
	Author *TpersonType `xml:"http://www.w3.org/2005/Atom author"`
}

type XsdGoPkgHasElems_Author_TpersonType_ struct {
	Authors []*TpersonType `xml:"http://www.w3.org/2005/Atom author"`
}

type XsdGoPkgHasElem_Category_TcategoryType_ struct {
	Category *TcategoryType `xml:"http://www.w3.org/2005/Atom category"`
}

type XsdGoPkgHasElems_Category_TcategoryType_ struct {
	Categories []*TcategoryType `xml:"http://www.w3.org/2005/Atom category"`
}

type XsdGoPkgHasElem_Contributor_TpersonType_ struct {
	Contributor *TpersonType `xml:"http://www.w3.org/2005/Atom contributor"`
}

type XsdGoPkgHasElems_Contributor_TpersonType_ struct {
	Contributors []*TpersonType `xml:"http://www.w3.org/2005/Atom contributor"`
}

type XsdGoPkgHasElems_Generator_TgeneratorType_ struct {
	Generators []*TgeneratorType `xml:"http://www.w3.org/2005/Atom generator"`
}

type XsdGoPkgHasElem_Generator_TgeneratorType_ struct {
	Generator *TgeneratorType `xml:"http://www.w3.org/2005/Atom generator"`
}

type XsdGoPkgHasElem_Icon_TiconType_ struct {
	Icon *TiconType `xml:"http://www.w3.org/2005/Atom icon"`
}

type XsdGoPkgHasElems_Icon_TiconType_ struct {
	Icons []*TiconType `xml:"http://www.w3.org/2005/Atom icon"`
}

type XsdGoPkgHasElems_Id_TidType_ struct {
	Ids []*TidType `xml:"http://www.w3.org/2005/Atom id"`
}

type XsdGoPkgHasElem_Id_TidType_ struct {
	Id *TidType `xml:"http://www.w3.org/2005/Atom id"`
}

type XsdGoPkgHasElem_Link_TlinkType_ struct {
	Link *TlinkType `xml:"http://www.w3.org/2005/Atom link"`
}

type XsdGoPkgHasElems_Link_TlinkType_ struct {
	Links []*TlinkType `xml:"http://www.w3.org/2005/Atom link"`
}

type XsdGoPkgHasElems_Logo_TlogoType_ struct {
	Logos []*TlogoType `xml:"http://www.w3.org/2005/Atom logo"`
}

type XsdGoPkgHasElem_Logo_TlogoType_ struct {
	Logo *TlogoType `xml:"http://www.w3.org/2005/Atom logo"`
}

type XsdGoPkgHasElem_Rights_TtextType_ struct {
	Rights *TtextType `xml:"http://www.w3.org/2005/Atom rights"`
}

type XsdGoPkgHasElems_Rights_TtextType_ struct {
	Rightses []*TtextType `xml:"http://www.w3.org/2005/Atom rights"`
}

type XsdGoPkgHasElem_Subtitle_TtextType_ struct {
	Subtitle *TtextType `xml:"http://www.w3.org/2005/Atom subtitle"`
}

type XsdGoPkgHasElems_Subtitle_TtextType_ struct {
	Subtitles []*TtextType `xml:"http://www.w3.org/2005/Atom subtitle"`
}

type XsdGoPkgHasElem_Title_TtextType_ struct {
	Title *TtextType `xml:"http://www.w3.org/2005/Atom title"`
}

type XsdGoPkgHasElems_Title_TtextType_ struct {
	Titles []*TtextType `xml:"http://www.w3.org/2005/Atom title"`
}

type XsdGoPkgHasElems_Updated_TdateTimeType_ struct {
	Updateds []*TdateTimeType `xml:"http://www.w3.org/2005/Atom updated"`
}

type XsdGoPkgHasElem_Updated_TdateTimeType_ struct {
	Updated *TdateTimeType `xml:"http://www.w3.org/2005/Atom updated"`
}

type XsdGoPkgHasElems_Entry_TentryType_ struct {
	Entries []*TentryType `xml:"http://www.w3.org/2005/Atom entry"`
}

type XsdGoPkgHasElem_Entry_TentryType_ struct {
	Entry *TentryType `xml:"http://www.w3.org/2005/Atom entry"`
}

//	The Atom feed construct is defined in section 4.1.1 of the format spec.
type TfeedType struct {
	XsdGoPkgHasElems_Author_TpersonType_

	XsdGoPkgHasElems_Category_TcategoryType_

	XsdGoPkgHasElems_Contributor_TpersonType_

	XsdGoPkgHasElems_Generator_TgeneratorType_

	XsdGoPkgHasElems_Icon_TiconType_

	XsdGoPkgHasElems_Id_TidType_

	XsdGoPkgHasElems_Link_TlinkType_

	XsdGoPkgHasElems_Logo_TlogoType_

	XsdGoPkgHasElems_Rights_TtextType_

	XsdGoPkgHasElems_Subtitle_TtextType_

	XsdGoPkgHasElems_Title_TtextType_

	XsdGoPkgHasElems_Updated_TdateTimeType_

	XsdGoPkgHasElems_Entry_TentryType_

	XsdGoPkgHasAtts_CommonAttributes

}

type XsdGoPkgHasElem_Content_TcontentType_ struct {
	Content *TcontentType `xml:"http://www.w3.org/2005/Atom content"`
}

type XsdGoPkgHasElems_Content_TcontentType_ struct {
	Contents []*TcontentType `xml:"http://www.w3.org/2005/Atom content"`
}

type XsdGoPkgHasElems_Published_TdateTimeType_ struct {
	Publisheds []*TdateTimeType `xml:"http://www.w3.org/2005/Atom published"`
}

type XsdGoPkgHasElem_Published_TdateTimeType_ struct {
	Published *TdateTimeType `xml:"http://www.w3.org/2005/Atom published"`
}

type XsdGoPkgHasElems_Source_TtextType_ struct {
	Sources []*TtextType `xml:"http://www.w3.org/2005/Atom source"`
}

type XsdGoPkgHasElem_Source_TtextType_ struct {
	Source *TtextType `xml:"http://www.w3.org/2005/Atom source"`
}

type XsdGoPkgHasElems_Summary_TtextType_ struct {
	Summaries []*TtextType `xml:"http://www.w3.org/2005/Atom summary"`
}

type XsdGoPkgHasElem_Summary_TtextType_ struct {
	Summary *TtextType `xml:"http://www.w3.org/2005/Atom summary"`
}

//	The Atom entry construct is defined in section 4.1.2 of the format spec.
type TentryType struct {
	XsdGoPkgHasElems_Content_TcontentType_

	XsdGoPkgHasElems_Published_TdateTimeType_

	XsdGoPkgHasElems_Source_TtextType_

	XsdGoPkgHasElems_Summary_TtextType_

	XsdGoPkgHasAtts_CommonAttributes

}

type XsdGoPkgHasAttr_Type_XsdtString_ struct {
	Type xsdt.String `xml:"http://www.w3.org/2005/Atom type,attr"`
}

type XsdGoPkgHasAttr_Src_XsdtAnyURI_ struct {
	Src xsdt.AnyURI `xml:"http://www.w3.org/2005/Atom src,attr"`
}

//	The Atom content construct is defined in section 4.1.3 of the format spec.
type TcontentType struct {
	XsdGoPkgHasCdata

	XsdGoPkgHasAtts_CommonAttributes

	XsdGoPkgHasAttr_Src_XsdtAnyURI_

	XsdGoPkgHasAttr_Type_XsdtString_

}

type XsdGoPkgHasAttr_Term_XsdtString_ struct {
	Term xsdt.String `xml:"http://www.w3.org/2005/Atom term,attr"`
}

type XsdGoPkgHasAttr_Scheme_XsdtAnyURI_ struct {
	Scheme xsdt.AnyURI `xml:"http://www.w3.org/2005/Atom scheme,attr"`
}

type XsdGoPkgHasAttr_Label_XsdtString_ struct {
	Label xsdt.String `xml:"http://www.w3.org/2005/Atom label,attr"`
}

//	The Atom cagegory construct is defined in section 4.2.2 of the format spec.
type TcategoryType struct {
	XsdGoPkgHasAtts_CommonAttributes

	XsdGoPkgHasAttr_Term_XsdtString_

	XsdGoPkgHasAttr_Scheme_XsdtAnyURI_

	XsdGoPkgHasAttr_Label_XsdtString_

}

type XsdGoPkgHasAttr_Uri_XsdtAnyURI_ struct {
	Uri xsdt.AnyURI `xml:"http://www.w3.org/2005/Atom uri,attr"`
}

type XsdGoPkgHasAttr_Version_XsdtString_ struct {
	Version xsdt.String `xml:"http://www.w3.org/2005/Atom version,attr"`
}

//	The Atom generator element is defined in section 4.2.4 of the format spec.
type TgeneratorType struct {
	XsdGoPkgValue xsdt.String `xml:",chardata"`

	XsdGoPkgHasAtts_CommonAttributes

	XsdGoPkgHasAttr_Uri_XsdtAnyURI_

	XsdGoPkgHasAttr_Version_XsdtString_

}

//	The Atom icon construct is defined in section 4.2.5 of the format spec.
type TiconType struct {
	XsdGoPkgValue xsdt.AnyURI `xml:",chardata"`

	XsdGoPkgHasAtts_CommonAttributes

}

//	The Atom id construct is defined in section 4.2.6 of the format spec.
type TidType struct {
	XsdGoPkgValue xsdt.AnyURI `xml:",chardata"`

	XsdGoPkgHasAtts_CommonAttributes

}

type XsdGoPkgHasAttr_Href_XsdtAnyURI_ struct {
	Href xsdt.AnyURI `xml:"http://www.w3.org/2005/Atom href,attr"`
}

type XsdGoPkgHasAttr_Rel_XsdtString_ struct {
	Rel xsdt.String `xml:"http://www.w3.org/2005/Atom rel,attr"`
}

type XsdGoPkgHasAttr_Hreflang_XsdtNmtoken_ struct {
	Hreflang xsdt.Nmtoken `xml:"http://www.w3.org/2005/Atom hreflang,attr"`
}

type XsdGoPkgHasAttr_Title_XsdtString_ struct {
	Title xsdt.String `xml:"http://www.w3.org/2005/Atom title,attr"`
}

type XsdGoPkgHasAttr_Length_XsdtPositiveInteger_ struct {
	Length xsdt.PositiveInteger `xml:"http://www.w3.org/2005/Atom length,attr"`
}

//	The Atom link construct is defined in section 3.4 of the format spec.
type TlinkType struct {
	XsdGoPkgHasCdata

	XsdGoPkgHasAtts_CommonAttributes

	XsdGoPkgHasAttr_Title_XsdtString_

	XsdGoPkgHasAttr_Rel_XsdtString_

	XsdGoPkgHasAttr_Length_XsdtPositiveInteger_

	XsdGoPkgHasAttr_Hreflang_XsdtNmtoken_

	XsdGoPkgHasAttr_Href_XsdtAnyURI_

}

//	The Atom logo construct is defined in section 4.2.8 of the format spec.
type TlogoType struct {
	XsdGoPkgValue xsdt.AnyURI `xml:",chardata"`

	XsdGoPkgHasAtts_CommonAttributes

}

//	The Atom source construct is defined in section 4.2.11 of the format spec.
type TsourceType struct {
	XsdGoPkgHasAtts_CommonAttributes

}

type TuriType struct {
	XsdGoPkgValue xsdt.AnyURI `xml:",chardata"`

	XsdGoPkgHasAtts_CommonAttributes

}

type TdateTimeType struct {
	XsdGoPkgValue xsdt.DateTime `xml:",chardata"`

	XsdGoPkgHasAtts_CommonAttributes

}

type XsdGoPkgHasElem_Feed struct {
	Feed *TfeedType `xml:"http://www.w3.org/2005/Atom feed"`
}

type XsdGoPkgHasElems_Feed struct {
	Feeds []*TfeedType `xml:"http://www.w3.org/2005/Atom feed"`
}

type XsdGoPkgHasElems_Entry struct {
	Entries []*TentryType `xml:"http://www.w3.org/2005/Atom entry"`
}

type XsdGoPkgHasElem_Entry struct {
	Entry *TentryType `xml:"http://www.w3.org/2005/Atom entry"`
}
