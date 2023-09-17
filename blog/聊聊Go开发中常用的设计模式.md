## 这些Coding套路你不会还不知道吧？

对于一名程序员来说，编码进阶是成为优秀工程师非常重要的一步，它可以让我们更加熟练地掌握编程，深入理解数据结构和算法，从而更好地完成复杂的任务，提高工作效率。而我认为熟练使用设计模式就是编码进阶的最好方式之一，下面就用一篇文章来分享下Go开发中常用的设计模式。

- **:one:全局单一实例：单例模式**

- **:car:获取对象超简单：工厂模式**

- **:clipboard:重复代码太多？试试：模板方法模式**

- **:ghost:接口对应功能不知道怎么维护？这里有：策略模式**

- **:balloon:独特好玩的Functional Options模式**

### What？Why？

从两个问题开始今天的分享：

**设计模式是什么？**

简单来讲，就两个字：**`套路`**，编码的套路，在我们写代码的时候这个套路可以不用，但是不能不知道。

**为什么要使用设计模式？**

**设计模式是一种被反复使用的，针对软件设计中常见问题的可复用解决方案。它们提供了一种经过验证的方式来解决复杂的设计问题，并帮助开发人员编写更加清晰、可维护和可扩展的代码**。

使用设计模式的好处包括：

1. 提高代码的可读性和可维护性。通过使用设计模式，开发人员可以将复杂的问题分解成更小、更易于管理的部分，并且可以将这些部分组织成一致的、易于理解的代码结构。
2. 提高代码的可重用性。设计模式提供了一种标准化的解决方案，可以让相同的功能代码在多个地方重复使用，从而避免了重复编写相同的代码，减少了开发时间和成本。
3. 提高代码的灵活性和扩展性。设计模式允许开发人员在不改变现有代码的情况下，轻松地添加新的功能或修改现有的功能。

### 常见的设计模式实践

#### 1 全局单一实例：单例模式

首先，单例模式一般用于以下业务场景：

（1）整个程序的运行中只允许有一个类的实例。

（2）创建对象时耗时过多或者耗资源过多，但又经常用到的。

（3）需要全局访问并保证线程安全的对象。

（4）需要保持状态的对象。

具体场景：

数据库连接对象。我们在具体的项目中往往每种数据库驱动只会使用它的一个实例，因此在这里我们就能够使用单例模式。

代码：

```go
import "sync"

type MysqlConn struct {
	Addr string
}

var (
	mysqlConn *MysqlConn
	once = sync.Once{}
)

func GetMySQLConn() *MysqlConn {
	once.Do(func() {
		mysqlConn = &MysqlConn{Addr: "127.0.0.1"}
	})
	return mysqlConn
}
```

#### 2 获取对象超简单：工厂模式

工厂模式通常应用于以下场景：

（1）当一个系统需要灵活地配置一组对象，并且需要动态地选择其中的一个时。

（2）当一个类需要频繁地创建和销毁时，可以使用工厂模式来提高性能。

具体场景：

简单来讲就是想要获取一个实例，但是这个实例的属性可能会随着参数的变化而具有不同的形态，还可以用数据库连接来举例，我们有多个服务器可以进行连接，但是每次只能连接一个，而且连接的时候可以填写自己的认证账密。

代码：

```go
type DB struct {
   Addr, UserName, Passwd string
}

func NewMysqlConn(addr, userName, passwd string) *DB {
   return &DB{
      Addr:     addr,
      UserName: userName,
      Passwd:   passwd,
   }
}
```

#### 3 重复代码太多？试试：模板方法模式

模板方法模式通常应用于以下场景：

（1）当一个算法的步骤中有一部分是不变的，而另一部分是需要根据不同的条件进行变化时，可以使用模板方法模式来实现。

（2）当一个类的子类需要实现一个接口的不同版本时，可以使用模板方法模式来实现。

具体场景：

[做水果酸奶](https://mp.weixin.qq.com/s?__biz=MzIxNDc2ODc3MA==&mid=2247485000&idx=1&sn=c1e46cfbf817aea90398b57b4630895d&chksm=97a3cba5a0d442b3fcb9bb2814d38612c540dcdda0cdbd9d18f783b87781e64b7897af708c6d#rd)，水果或其他食材是统一的抽象，而火龙果、芒果、饼干都是具体，因为不管用什么水果或食材来做，做法都是如出一辙的，所以做法也是一个抽象，这就好比一个模板，对，就是模版方法模式。

代码：

```go
import (
   "fmt"
   "testing"
)

type MakeYogurtTemplate interface {
	CreateYogurt() //准备好酸奶
	CutFruit()     //切水果
	Merge()        //放在一起搅拌
	Optimize()     //调制味道
	Eating()       //开吃
	Do()
}

type DragonFruit struct {
}

func (d *DragonFruit) CreateYogurt() {
   fmt.Println("用鲜奶和乳酸菌发酵好酸奶")
}

func (d *DragonFruit) CutFruit() {
   fmt.Println("把火龙果切成块块")
}

func (d *DragonFruit) Merge() {
   fmt.Println("放在一起进行搅拌")
}

func (d *DragonFruit) Optimize() {
   fmt.Println("调制出自己喜欢的味道")
}

func (d *DragonFruit) Eating() {
   fmt.Println("开吃")
}

func (d *DragonFruit) Do() {
   d.CreateYogurt()
   d.CutFruit()
   d.Merge()
   d.Optimize()
   d.Eating()
}

func TestDragonFruit(t *testing.T) {
   d := &DragonFruit{}
   d.Do()
}

type Mango struct {
}

func (m *Mango) CreateYogurt() {
   fmt.Println("用鲜奶和乳酸菌发酵好酸奶")
}

func (m *Mango) CutFruit() {
   fmt.Println("把芒果切成块块")
}

func (m *Mango) Merge() {
   fmt.Println("放在一起进行搅拌")
}

func (m *Mango) Optimize() {
   fmt.Println("调制出自己喜欢的味道")
}

func (m *Mango) Eating() {
   fmt.Println("开吃")
}

func (m *Mango) Do() {
   m.CreateYogurt()
   m.CutFruit()
   m.Merge()
   m.Optimize()
   m.Eating()
}

func TestMango(t *testing.T) {
   m := &Mango{}
   m.Do()
}
```

#### 4 接口对应功能不知道怎么维护？这里有：策略模式

策略模式通常在一个系统中需要多个算法，并且这些算法需要在不同的时间或条件下使用不同的算法时，可以使用策略模式来实现。

具体场景：

对于新手程序员同学经常会有这样的选择，到底是学Go还是学Java，这显然是两种不同的策略选择，但是不管是学习哪一个的基本流程都是差不多的，比如都是先准备学习资料，然后在学习和实践，因此可以基于此使用策略模式来形容。

代码：

策略抽象

```go
type Learn interface {
   PrepareData()
   DoLearn()
   Play()
}

func LearnLang(l Learn) {
   l.PrepareData()
   l.DoLearn()
   l.Play()
}
```

策略1

```go
import "fmt"

type LikeGo struct {
}

func (g *LikeGo) PrepareData() {
   fmt.Println("准备Go资料")
}

func (g *LikeGo) DoLearn() {
   fmt.Println("学习Go")
}

func (g *LikeGo) Play() {
   fmt.Println("玩转Go语言")
}
```

策略2

```go
import "fmt"

type LikeJava struct {
}

func (g *LikeJava) PrepareData() {
   fmt.Println("准备Java资料")
}

func (g *LikeJava) DoLearn() {
   fmt.Println("学习Java")
}

func (g *LikeJava) Play() {
   fmt.Println("玩Java")
}
```

执行策略

```go
import "testing"

func TestLearn(t *testing.T) {
   likeGo := &LikeGo{}
   LearnLang(likeGo)
}
```

#### 5 独特好玩的Functional Options模式

Functional Options模式通常在当一个对象或函数具有多个参数时，这些参数可能会有不同的默认值或取值范围。通过将它们作为函数的选项传递，可以更灵活地控制函数的行为

具体场景：

gRPC服务进行实例化的时候有些参数可以不同填，即选填，之后在源码内部会使用默认的值，于是我们就可以使用以下的方式进行处理。

代码：

```go
import "time"

type RpcServer struct {
   Name    string
   MaxConn int
   Address []string
   TimeOut time.Duration
}

type RpcServerOption func(server *RpcServer)

func WithName(name string) RpcServerOption {
   return func(server *RpcServer) {
      server.Name = name
   }
}

func WithMaxConn(max int) RpcServerOption {
   return func(server *RpcServer) {
      server.MaxConn = max
   }
}

func WithAddress(addr []string) RpcServerOption {
   return func(server *RpcServer) {
      server.Address = addr
   }
}

func WithTimeOut(timeout time.Duration) RpcServerOption {
   return func(server *RpcServer) {
      server.TimeOut = timeout
   }
}

func NewRpcServer(opts ...RpcServerOption) *RpcServer {
   server := &RpcServer{}

   for _, opt := range opts {
      opt(server)
   }

   return server
}
```

实例化：

```go
import (
   "fmt"
   "testing"
   "time"
)

func TestCreateRpcServerByOptions(t *testing.T) {
   rpcServer := NewRpcServer(
      WithAddress([]string{"127.0.0.1"}),
      WithName("rpcServer"),
      WithMaxConn(1),
      WithTimeOut(time.Second),
   )

   fmt.Println(*rpcServer)
}
```

### 小总结

本文主要介绍了Go开发中常用的设计模式，包括全局单一实例：单例模式、工厂模式、模板方法模式、策略模式和Functional Options模式。这些设计模式可以帮助我们更好地组织代码，提高代码的可读性和可重用性。

总之，掌握这些设计模式对于提高Go程序员的编码能力非常有帮助，可以让我们在编写代码时更加得心应手，同时也能提高代码的质量和可维护性。
