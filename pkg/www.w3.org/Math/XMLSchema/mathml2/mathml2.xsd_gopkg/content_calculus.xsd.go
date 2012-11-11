//	Auto-generated by the "go-xsd" package located at:
//		github.com/metaleap/go-xsd
//	Comments on types and fields (if any) are from the XSD file located at:
//		www.w3.org/Math/XMLSchema/mathml2/content/calculus.xsd
package gopkg_WwwW3OrgMathXMLSchemaMathml2Mathml2Xsd

//	This is an XML Schema module for the calculs operators of content
//	MathML.
//	Author: Stéphane Dalmas, INRIA.
import (
)

type XsdGoPkgHasAtts_IntAttlist struct {
	XsdGoPkgHasAtts_DefinitionAttrib

	XsdGoPkgHasAtts_CommonAttrib

}

type XsdGoPkgHasAtts_DiffAttlist struct {
	XsdGoPkgHasAtts_DefinitionAttrib

	XsdGoPkgHasAtts_CommonAttrib

}

type XsdGoPkgHasAtts_PartialdiffAttlist struct {
	XsdGoPkgHasAtts_DefinitionAttrib

	XsdGoPkgHasAtts_CommonAttrib

}

type XsdGoPkgHasAtts_LimitAttlist struct {
	XsdGoPkgHasAtts_DefinitionAttrib

	XsdGoPkgHasAtts_CommonAttrib

}

type XsdGoPkgHasAtts_LowlimitAttlist struct {
	XsdGoPkgHasAtts_DefinitionAttrib

	XsdGoPkgHasAtts_CommonAttrib

}

type XsdGoPkgHasAtts_UplimitAttlist struct {
	XsdGoPkgHasAtts_DefinitionAttrib

	XsdGoPkgHasAtts_CommonAttrib

}

type XsdGoPkgHasAtts_TendstoAttlist struct {
	XsdGoPkgHasAtts_DefinitionAttrib

	XsdGoPkgHasAtts_CommonAttrib

}

type TintType struct {
	XsdGoPkgHasAtts_IntAttlist

}

type TdiffType struct {
	XsdGoPkgHasAtts_DiffAttlist

}

type TpartialdiffType struct {
	XsdGoPkgHasAtts_PartialdiffAttlist

}

type TlimitType struct {
	XsdGoPkgHasAtts_LimitAttlist

}

type TlowlimitType struct {
	XsdGoPkgHasGroup_LowlimitContent

	XsdGoPkgHasAtts_LowlimitAttlist

}

type TuplimitType struct {
	XsdGoPkgHasGroup_UplimitContent

	XsdGoPkgHasAtts_UplimitAttlist

}

type TtendstoType struct {
	XsdGoPkgHasAtts_TendstoAttlist

}

type XsdGoPkgHasElem_Int struct {
	Int *TintType `xml:"http://www.w3.org/1998/Math/MathML int"`
}

type XsdGoPkgHasElems_Int struct {
	Ints []*TintType `xml:"http://www.w3.org/1998/Math/MathML int"`
}

type XsdGoPkgHasElems_Diff struct {
	Diffs []*TdiffType `xml:"http://www.w3.org/1998/Math/MathML diff"`
}

type XsdGoPkgHasElem_Diff struct {
	Diff *TdiffType `xml:"http://www.w3.org/1998/Math/MathML diff"`
}

type XsdGoPkgHasElems_Partialdiff struct {
	Partialdiffs []*TpartialdiffType `xml:"http://www.w3.org/1998/Math/MathML partialdiff"`
}

type XsdGoPkgHasElem_Partialdiff struct {
	Partialdiff *TpartialdiffType `xml:"http://www.w3.org/1998/Math/MathML partialdiff"`
}

type XsdGoPkgHasElems_Limit struct {
	Limits []*TlimitType `xml:"http://www.w3.org/1998/Math/MathML limit"`
}

type XsdGoPkgHasElem_Limit struct {
	Limit *TlimitType `xml:"http://www.w3.org/1998/Math/MathML limit"`
}

type XsdGoPkgHasElem_Lowlimit struct {
	Lowlimit *TlowlimitType `xml:"http://www.w3.org/1998/Math/MathML lowlimit"`
}

type XsdGoPkgHasElems_Lowlimit struct {
	Lowlimits []*TlowlimitType `xml:"http://www.w3.org/1998/Math/MathML lowlimit"`
}

type XsdGoPkgHasElems_Uplimit struct {
	Uplimits []*TuplimitType `xml:"http://www.w3.org/1998/Math/MathML uplimit"`
}

type XsdGoPkgHasElem_Uplimit struct {
	Uplimit *TuplimitType `xml:"http://www.w3.org/1998/Math/MathML uplimit"`
}

type XsdGoPkgHasElem_Tendsto struct {
	Tendsto *TtendstoType `xml:"http://www.w3.org/1998/Math/MathML tendsto"`
}

type XsdGoPkgHasElems_Tendsto struct {
	Tendstos []*TtendstoType `xml:"http://www.w3.org/1998/Math/MathML tendsto"`
}

type XsdGoPkgHasGroup_LowlimitContent struct {
	XsdGoPkgHasGroup_ContentExprClass

}

type XsdGoPkgHasGroup_UplimitContent struct {
	XsdGoPkgHasGroup_ContentExprClass

}

type XsdGoPkgHasGroup_ContentCalculusClass struct {
	XsdGoPkgHasElems_Int

	XsdGoPkgHasElems_Diff

	XsdGoPkgHasElems_Partialdiff

	XsdGoPkgHasElems_Limit

	XsdGoPkgHasElems_Lowlimit

	XsdGoPkgHasElems_Uplimit

	XsdGoPkgHasElems_Tendsto

}
