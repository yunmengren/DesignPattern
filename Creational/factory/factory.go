package factory

type IRuleConfigParser interface {
	Parse(data []byte)
}

type jsonRuleConfigParser struct {
}

func (j jsonRuleConfigParser) Parse(data []byte) {
	panic("implement me")
}

type xmlRuleConfigParser struct {
}

func (x xmlRuleConfigParser) Parse(data []byte) {
	panic("implement me")
}

func NewRuleConfigParser(configFormat string) IRuleConfigParser {
	switch configFormat {
	case "json":
		return jsonRuleConfigParser{}
	case "xml":
		return xmlRuleConfigParser{}
	}
	return nil
}

type IRuleConfigParserFactory interface {
	CreateParser() IRuleConfigParser
}

type jsonRuleConfigParserFactory struct {
}

func (j jsonRuleConfigParserFactory) CreateParser() IRuleConfigParser {
	return jsonRuleConfigParser{}
}

type xmlRuleConfigParserFactory struct {
}

func (x xmlRuleConfigParserFactory) CreateParser() IRuleConfigParser {
	return xmlRuleConfigParser{}
}

// 此处用一个简单工厂封装工厂方法
func NewIRuleConfigParserFactory(configFormat string) IRuleConfigParserFactory {
	switch configFormat {
	case "json":
		return jsonRuleConfigParserFactory{}
	case "xml":
		return xmlRuleConfigParserFactory{}
	}
	return nil
}
