package values

import (
	"regexp"
)

const (
	// @opt <opt> <desc>
	optRe  = `^\s*##\s\@opt\s(?P<opt>\w{2,64})\s(?P<desc>.+)$`
	optTpl = "| $opt | `$desc` |\n"

	// @mand <opt> <desc>
	mandOptRe  = `^\s*##\s\@mand\s(?P<opt>\w{2,64})\s(?P<desc>.+)$`
	mandOptTpl = "| **$opt** | `$desc` |\n"

	// @obj <objectName> <desc>
	objNameRe  = `^\s*##\s\@obj\s(?P<objName>\w{2,64})\s(?P<desc>.+)$`
	objNameTpl = "$objName $desc"

	// @obj_mand <objectName> <desc>
	mandObjNameRe  = `^\s*##\s\@obj_mand\s(?P<objName>\w{2,64})\s(?P<desc>.+)$`
	mandObjNameTpl = "$objName $desc"

	// @obj_mand <objectName>.<objOpt> <desc>
	mandObjOptRe  = `^\s*##\s\@obj_mand\s(?P<objName>\w{2,64})\.(?P<objOpt>\w{2,64})\s(?P<desc>.+)$`
	mandObjOptTpl = "$objName $objOpt $desc"

	// @obj <objectName>.<objOpt> <desc>
	objRe  = `^\s*##\s\@obj\s(?P<objName>\w{2,64})\.(?P<objOpt>\w{2,64})\s(P<desc>.+)$`
	objTpl = "$objName $objOpt $desc"

	// @section <secName>
	sectionRe  = `^\s*##\s\@section\s(?P<secName>\w{2,64})$`
	sectionTpl = "\n**$secName:**\n"

	tableHeader = "| name | description |\n| --- | --- |\n"
)

var (
	optRE         *regexp.Regexp
	mandOptRE     *regexp.Regexp
	objNameRE     *regexp.Regexp
	mandObjNameRE *regexp.Regexp
	objRE         *regexp.Regexp
	mandObjOptRE  *regexp.Regexp
	sectionRE     *regexp.Regexp
)

func init() {
	optRE = regexp.MustCompile(optRe)
	mandOptRE = regexp.MustCompile(mandOptRe)
	objNameRE = regexp.MustCompile(objNameRe)
	mandObjNameRE = regexp.MustCompile(mandObjNameRe)
	objRE = regexp.MustCompile(objRe)
	mandObjOptRE = regexp.MustCompile(mandObjOptRe)
	sectionRE = regexp.MustCompile(sectionRe)
}

func isOpt(d []byte) bool {
	return optRE.Match(d)
}

func renderOpt(d []byte) []byte {
	data := make([]byte, 4096)
	data = optRE.Expand(data, []byte(optTpl), d, optRE.FindSubmatchIndex(d))
	return data
}

func isMandOpt(d []byte) bool {
	return mandOptRE.Match(d)
}

func renderMandOpt(d []byte) []byte {
	data := make([]byte, 4096)
	data = mandOptRE.Expand(data, []byte(mandOptTpl), d, mandOptRE.FindSubmatchIndex(d))
	return data
}

func isObjName(d []byte) bool {
	return objNameRE.Match(d)
}

func renderObjName(d []byte) []byte {
	data := make([]byte, 4096)
	data = objNameRE.Expand(data, []byte(objNameTpl), d, objNameRE.FindSubmatchIndex(d))
	return data
}

func isMandObjName(d []byte) bool {
	return mandObjNameRE.Match(d)
}

func renderMandObjName(d []byte) []byte {
	data := make([]byte, 4096)
	data = mandObjNameRE.Expand(data, []byte(mandObjNameTpl), d, mandObjNameRE.FindSubmatchIndex(d))
	return data
}

func isObjOpt(d []byte) bool {
	return objRE.Match(d)
}

func renderObjOpt(d []byte) []byte {
	data := make([]byte, 4096)
	data = objRE.Expand(data, []byte(objTpl), d, objRE.FindSubmatchIndex(d))
	return data
}

func isMandObjOpt(d []byte) bool {
	return mandObjOptRE.Match(d)
}

func renderMandObjOpt(d []byte) []byte {
	data := make([]byte, 4096)
	data = mandObjOptRE.Expand(data, []byte(mandObjOptTpl), d, mandObjOptRE.FindSubmatchIndex(d))
	return data
}

func isSection(d []byte) bool {
	return sectionRE.Match(d)
}

func renderSection(d []byte) []byte {
	data := make([]byte, 4096)
	data = sectionRE.Expand(data, []byte(sectionTpl), d, sectionRE.FindSubmatchIndex(d))
	return data
}
