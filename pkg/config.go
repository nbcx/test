package pkg

type Config struct {
	Hacker   bool     `mapstructure:"hacker"`
	Name     int      `mapstructure:"name"`
	Hobbies  []string `mapstructure:"hobbies"`
	Clothing struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
	} `mapstructure:"clothing"`
}

type Configer interface {
	Construct(Configer)
	Name() string // 配置文件的名字
	// Bind(v interface{})     // 设置绑定解析的映射结构体
	Check(string, string) bool // 检查配置是否正确,不正确将中断更新，不会进行Update和Callback的调用了
	Callback()                 // 完成结构体的解析映射后，回调
}

type figer interface {
	Bind(v interface{})
	Load(...Configer)
	LoadWithPath(string, ...Configer)
}
