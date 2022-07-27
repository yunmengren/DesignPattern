package abstract_factory

type IRuleConfigParser interface {
	ParseRule(data []byte)
}

type ISystemConfigParser interface {
	ParseSystem(data []byte)
}

type IConfigParserFactory interface {
	CreateRuleParser() IRuleConfigParser
	CreateSystemParser() ISystemConfigParser
}

type jsonConfigParserFactory struct{}

type jsonRuleConfigParser struct {
}

func (j jsonRuleConfigParser) ParseRule(data []byte) {
	panic("implement me")
}

type jsonSystemConfigParser struct {
}

func (j jsonSystemConfigParser) ParseSystem(data []byte) {
	panic("implement me")
}

func (j jsonConfigParserFactory) CreateRuleParser() IRuleConfigParser {
	return jsonRuleConfigParser{}
}

func (j jsonConfigParserFactory) CreateSystemParser() ISystemConfigParser {
	return jsonSystemConfigParser{}
}

type xmlConfigParserFactory struct{}

type xmlRuleConfigParser struct {
}

func (j xmlRuleConfigParser) ParseRule(data []byte) {
	panic("implement me")
}

type xmlSystemConfigParser struct {
}

func (j xmlSystemConfigParser) ParseSystem(data []byte) {
	panic("implement me")
}

func (j xmlConfigParserFactory) CreateRuleParser() IRuleConfigParser {
	return xmlRuleConfigParser{}
}

func (j xmlConfigParserFactory) CreateSystemParser() ISystemConfigParser {
	return xmlSystemConfigParser{}
}
