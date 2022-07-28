## 单例模式
单例设计模式（Singleton Design Pattern)。一个类只允许创建一个对象（或者实例），那这个类就是一个单例类

#### 应用场景
1. 处理资源访问冲突
例如往文件中打印日志的 Logger 类，日志文件就是一个临界资源
2. 表示全局唯一类
有些数据在系统中只应保存一份，例如配置信息类。只有一个配置文件，当配置文件被加载到内存之后，以对象的形式存在，理所应当止只有一份。

#### 实现一个单例
我们需要关注下面几点：
- 构造函数需要是 private 访问权限的，这样才能避免外部通过 new 创建实例
- 考虑对象创建时的线程安全问题
- 考虑是否支持延迟加载
- 考虑 getInstance() 性能是否高（是否加锁）
#### 1.饿汉式
在类加载的时候，instance 静态实例就已经创建并初始化好了，所以，实例的创建过程时线程安全的。不过这样的实现方式不支持延迟加载（即在真正用到的时候，再创建实例）
```go
package idgenerator

import "sync/atomic"

type idEagerGenerator struct {
	id uint64
}

var iEG *idEagerGenerator

func init() {
	iEG = &idEagerGenerator{
		id: 0,
	}
}

func GetEagerInstance() *idEagerGenerator {
	return iEG
}

func (self *idEagerGenerator) GetId() uint64 {
	return atomic.AddUint64(&self.id, 1)
}
```
#### 2.懒汉式
在使用的时候， getInstance() 如果静态实例对象还没有被创建，则创建，每次需要判断，创建实例时要考虑线程安全问题
```go
package idgenerator

import (
	"sync"
	"sync/atomic"
)

type idLazyGenerator struct {
	id uint64
}

var (
	IdLazyGenerator *idLazyGenerator
	once            = &sync.Once{}
)

func GetLazyInstance() *idLazyGenerator {
	if IdLazyGenerator == nil {
		once.Do(func() {
			IdLazyGenerator = &idLazyGenerator{
				id: 0,
			}
		})
	}
	return IdLazyGenerator
}

func (self *idLazyGenerator) GetId() uint64 {
	return atomic.AddUint64(&self.id, 1)
}
```
#### 3.双重检测
#### 4.静态内部类
利用 Java 的静态内部类来实现单例。
#### 5.枚举
通过 Java 枚举类型本身的特性，保证实例创建的线程安全性和实例的唯一性。

#### 单例模式存在的问题
大部分情况，在项目中使用单例，都是用它来表示一些全局唯一类，比如配置信息类、连接池类、ID生成器类。
单例模式书写简洁、使用方便，在代码中，我们不需要创建对象，直接通过类似 IdGenerator.getInstance().GetId() 这样的方式调用就可以了。
##### 1.单例对 OOP 特性支持不友好
单例这种实现方式违背了基于接口而非实现编程的设计原则，如果未来某一天，我们希望针对不同的业务采用不同的 ID 生成算法，那么所有用到 IdGenerator 类的地方，都要进行修改。
##### 2.单例会隐藏类之间的依赖关系
单例类不需要显示创建、不需要依赖参数传递。
##### 3.单例对代码的扩展性不友好
如果单例类需要创建两个实例或多个实例，那么就要对代码有比较大的改动。类似于数据库连接池、线程池这类的资源池，最好不要设计成单例类。
##### 4.单例对代码的可测试性不友好
如果单例类依赖比较重的外部资源，比如 DB，我们在写单元测试的时候希望通过 mock 的方式将它替换掉，而单例类这种硬编码式的使用方式，导致无法实现 mock 替换。
此外，如果单例持有成员变量，那么它实际上是相当于一种全局变量，被所有的代码共享。
##### 5.单例不支持有参数的构造函数

#### 单例有什么替代解决方案
为了保证全局唯一，除了使用单例，还可以用静态方法来实现。
如果要完全解决问题，要从根上寻找其他方式来实现全局唯一类。比如，通过工厂模式，IOC 容器来保证或者由程序员自己来保证。

#### 如何理解单例模式中的唯一性
一般来说，唯一性指的是进程内只允许创建一个对象。但也有线程唯一的单例。
#### 如何实现线程唯一的单例
通过 HashMap 来存储对象，其中 key 是线程 ID, value 是对象，这样我们就可以做到不同的线程对应不同的对象。
#### 如何理解集群环境下的单例
集群相当于多个进程构成的一个集合。我们可以把单例对象序列化并存储到外部共享存储区(比如文件)，进程在使用这个对象的时候，需要先从外部共享存储区读取到内存，然后再使用，使用完后再存储回外部存储区。为了保证单例对象的唯一和安全性，一个进程在获取到对象后，需要对对象加锁，使用完后，释放改对象的加锁。
#### 如何实现一个多例模式
多例指的是一个类可以创建多个对象，但是个数是有限制的，比如只能创建 3 个对象，可以理解为同一类型的只能创建一个对象，不同类型的可以创建多个对象。
这种多例模式的理解方式有点类似于工厂模式。它跟工厂模式不同的是，多例模式创建的对象都是同一个类的对象，而工厂模式创建的是不同子类的对象。


## 工厂模式(Factory Design Pattern)
一般情况下，工厂模式分为三种更加细分的类型：简单工厂、工厂方法、抽象工厂。前两种比较简单，实际项目中也比较常用，而抽象工厂的原理稍微复杂点，在实际项目中也相对不常用。
我们举一个例子，根据配置文件的后缀（json, xml, yaml, properties),选择不同德解析器(JsonRuleConfigParser, XmlRuleConfigParser...),将存储在文件中的配置解析成内存对象 RuleConfig
#### 简单工厂(Simple Factory)
简单工厂模式，就是我们创建一个独立的类，让这个类来负责对象的创建
```go
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
```
#### 工厂方法
```go
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
```
这样当我们新增一种 parser 的时候，只需要新增一个实现了 IRuleConfigParserFactory 接口的 Factory 类即可。所以**工厂方法模式比起简单工厂模式更加符合开闭原则**
NewIRuleConfigParserFactory 相对当于**为工厂类再创建一个简单工厂，也就是工厂的工厂，用来创建工厂类对象**。
#### 什么时候该用工厂方法模式，而非简单工厂模式呢？
当对象的创建逻辑比较复杂，不只是简单的 new 一下就可以，而是要组合其他类对象，做各种初始化操作的时候，我们推荐使用工厂方法模式，将复杂的逻辑拆分到多个工厂类中，让每个工厂类都不至于过于复杂。而使用简单工厂模式，将所有的创建逻辑都放到一个工厂类中，会导致这个工厂类变得很负责。
除此之外，在某些场景下，如果对象不可服用，那工厂类每次都要返回不同的对象。如果我们使用简单工厂模式来实现，就只能选择第一种包含 if 或 switch 分支逻辑的实现方式。如果我们还想避免烦人的 if-else 分支逻辑，这个时候，我们就推荐只用工厂方法模式。
#### 抽象工厂(Abstract Factory)
在规则配置解析的那个例子中，解析器类只会根据配置文件格式(Json, xml, yaml...)来分类。但是如果有两种分类方式，比如，我们既可以按照配置文件格式来分类，也可以按照解析的对象（Rule规则配置还是System系统配置）来分类，那么就会对应下面者8个parser类。
```
针对规则配置的解析器：基于接口IRuleConfigParser
JsonRuleConfigParser
XmlRuleConfigParser
YamlRuleConfigParser
PropertiesRuleConfigParser
针对系统配置的解析器：基于接口ISystemConfigParser
JsonSystemConfigParser
XmlSystemConfigParser
YamlSystemConfigParser
PropertiesSystemConfigParser
```
针对这样的场景，如果还是继续使用工厂方法来实现的话，会有过多的类，让系统难以维护。
我们可以使用抽象工厂，让一个工厂负责创建多个不同类型的对象(IRuleConfigParser, ISystemConfigParser..)
```go
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
```

现在我们上升一个思维层面来看工厂模式，它的作用无外乎下面这四个。这也是判断亚欧吧要使用工厂模式做本质的**参考标准**
- 封装变化：创建逻辑有可能变化，封装成工厂类之后，创建逻辑的变更对调用者透明。
- 代码复用：创建代码抽离到独立的工厂类之后可以复用。
- 隔离复杂性：封装复杂的创建逻辑，调用者无需了解如何创建对象。
- 控制复杂度：将创建代码抽离出来，让原本的函数或类职责单一，代码更简洁。

## DI
依赖注入容器(Dependency Injection Container)，简称 DI 容器
#### 工厂模式和 DI 容器有何区别
实际上，DI 容器底层最基本的设计思路还是基于工厂模式的。DI 容器相当于一个大的工厂类，负责在程序启动的时候，根据配置（要创建哪些类对象，每个类对象的创建需要依赖哪些其他类对象）事先创建好对象。当应用程序需要使用某个类对象的时候，直接从容器中获取即可。正是因为它持有一堆对象，所以这个框架才被称为“容器”。
#### DI容器核心功能
- 配置解析
- 对象创建
- 对象生命周期管理
##### 配置解析
容器读取配置文件，根据配置文件提供的信息创建对象。
##### 对象创建
将所有类对象的创建都放到一个工厂类中完成就可以了，比如 BeansFactory。
利用反射机制，在程序运行的过程中，动态地加载类、创建对象，不需要事先在代码中写死要创建哪些对象。所以不管是创建一个对象，还是十个对象，BeansFactory 工厂类代码都是一样的。
##### 对象的生命周期管理
在简单工厂模式中，有两种实现方式，一种是每次都返回新创建的对象，另一种是每次都返回同一个事先创建好的对象，也就是所谓的单例对象。在 Spring 框架中，可以通过设置 scope 属性，来区分两种不同类型的对象。scope=prototype 表示返回新创建的对象，scope=singleton 表示返回单例对象。
除此之外，还可以配置对象是否支持懒加载。如果 lazy-init=true,对象在真正被使用到的时候（比如：BeanFactory.getBean("userService")）才被创建；如果 lazy-init=false,对象在应用启动的时候就事先创建好。
不仅如此，还可以配置对象的 init-method 和 destory-method 方法。DI 容器在创建好对象之后，会主动调用 init-method 属性指定的方法来初始化对象。在对象最终销毁之前，DI 容器也会主动调用 destory-method 属性指定的方法来做一些清理工作，比如释放数据库连接池、关闭文件。
#### 如何事先一个简单的 DI 容器？
可以通过反射机制来实现一个简单的 DI 容器，核心逻辑只需要包括这两个部分：配置文件解析、根据配置文件通过“反射”语法来创建对象。
##### 1.最小原型设计
##### 2.提供执行入口
##### 3.配置文件解析
##### 4.核心工厂类设计

## 建造者模式
Builder 模式，中文翻译为**建造者模式**或者**构建者模式**，也有人叫它**生成器模式**。
建造者模式的原理和实现比较简单，重点是掌握应用场景，避免过度使用。
如果一个类中有很多属性，为了避免构造函数的参数列表过长，影响代码的可读性和易用性，我们可以直接通过构造函数配合 set() 方法来解决。但是，**如果存在下面情况中的任意一种，我们就要考虑使用建造者模式了**。
- 我们把类的必填属性放到构造函数中，强制创建对象的时候设置。如果必填属性有很多，把这些必填属性都放到构造函数中设置，那构造函数就又会出现参数列表很长的问题。如果我们把必填属性通过 set() 方法设置，那校验这些必填属性是否已经填写的逻辑就无处安放了。
- 如果类的属性之间有一定的依赖关系或者约束条件，我们继续使用构造函数配合 set() 方法的设计思路，那些依赖关系或约束条件的校验逻辑就无处安放了。
- 如果我们希望创建不可变对象，也就是说，对象在创建好之后，就不能再修改内部的属性值，要实现这个功能，我们就不能在类中暴露 set() 方法。构造函数配合 set() 方法来设置属性值的方式就不适用了。
#### 与工厂模式的区别
- 工厂模式是用来创建不同但是相关类型的对象（继承同一父类或者接口的一组子类），由给定的参数来决定创建哪种类型的对象。
- 建造者模式是用来创建一种类型的复杂对象，可以通过设置不同的可选参数，“定制化”地创建不同的对象。
#### golang 实现
在 golang 中我们通常使用，参数传递方法来实现建造者模式
```go
type ResourcePoolConfigOptFunc func(option *ResourcePoolConfigOption)

func NewResourcePoolConfig(name string, opts ...ResourcePoolConfigOptFunc) (*ResourcePoolConfig, error) {
    if name == "" {
        return nil, fmt.Errorf("name can not be empty")
    }
	for _, opt := range opts {
        opt(option)
    }
}
```

## 原型模式
如果对象的创建成本比较大，而用一个类的不同对象之间差别不大（大部分字段都相同），在这种情况下，我们可以利用已有对象（原型）进行复制（或者叫拷贝）的方式来创建新对象，以达到节省创建时间的目的。这种基于原型来创建对象的方式就叫作**原型设计模式*Prototype Design Pattern)**，简称**原型模式**。
#### 原型模式的两种实现方法
有两种实现方法，深拷贝和浅拷贝。浅拷贝只会复制对象中基本数据类型和引用对象的内存地址，不会递归地复制引用对象，以及引用对象的引用对象...而深拷贝得到的是一份完完全全独立的对象。所以，深拷贝比起浅拷贝来说，更加耗时，更加耗内存空间。
如果要拷贝的对象是不可变对象，浅拷贝共享不可变对象是没有问题的，但对于可变对象来说，浅拷贝得到的对象和原始对象会共享部分数据，就有可能出现数据被修改的风险，也就变得复杂多了。
#### 如何实现深拷贝呢？
1. 第一种，递归拷贝对象、对象的引用以及引用对象的引用对象...直到要拷贝的对象只包含基本数据类型，没有引用对象位置。
2. 第二种，先将对象序列化，然后再反序列化成新的对象。
```go
// Keyword 搜索关键字
type Keyword struct {
    word      string
    visit     int
    UpdatedAt *time.Time
}

// Clone 这里使用序列化与反序列化的方式深拷贝
func (k *Keyword) Clone() *Keyword {
    var newKeyword Keyword
    b, _ := json.Marshal(k)
    json.Unmarshal(b, &newKeyword)
    return &newKeyword
}
```
无论哪种实现方式，深拷贝都要比浅拷贝耗时、耗内存空间。

#### 代码地址
https://github.com/yunmengren/DesignPattern
