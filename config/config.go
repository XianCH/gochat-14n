package config

import "time"

type app struct {
	Desc       string `yaml:"desc"`
	Addr       string `yaml:"addr"`
	ConfigFile string `yaml:"configFile"`
	Version    string `yaml:"version"`
	Env        string `yaml:"env"`
}

type ServerConfig struct {
	App   app   `yaml:"app"`
	Mysql mysql `yaml:"mysql"`
	Log   log   `yaml:"log"`
	Jwt   jwt   `yaml:"jwt"`
	Redis redis `yaml:"redis"`
	// Jwt     jwt     `yaml:"jwt"`
	// Redis   redis   `yaml:"redis"`
	// Elastic elastic `yaml:"elastic"`
}

type mysql struct {
	Host                      string        `yaml:"host"`
	Port                      string        `yaml:"port"`
	User                      string        `yaml:"user"`
	Password                  string        `yaml:"password"`
	Database                  string        `yaml:"database"`
	Charset                   string        `yaml:"charset"`                   //要支持完整的UTF-8编码,需设置成: utf8mb4
	AutoMigrate               bool          `yaml:"autoMigrate"`               // 初始化时调用数据迁移
	ParseTime                 bool          `yaml:"parseTime"`                 //解析time.Time类型
	TimeZone                  string        `yaml:"timeZone"`                  // 时区,若设置 Asia/Shanghai,需写成: Asia%2fShanghai
	DefaultStringSize         uint          `yaml:"defaultStringSize"`         // string 类型字段的默认长度
	DisableDatetimePrecision  bool          `yaml:"disableDatetimePrecision"`  // 禁用 datetime 精度
	SkipInitializeWithVersion bool          `yaml:"skipInitializeWithVersion"` // 根据当前 MySQL 版本自动配置
	Gorm                      gorm          `yaml:"gorm"`
	SlowSql                   time.Duration `yaml:"slowSql"`                   //慢SQL
	LogLevel                  string        `yaml:"logLevel"`                  // 日志记录级别
	IgnoreRecordNotFoundError bool          `yaml:"ignoreRecordNotFoundError"` // 是否忽略ErrRecordNotFound(未查到记录错误)
	Enable                    bool          `yaml:"enable"`
}

type gorm struct {
	SkipDefaultTx   bool   `yaml:"skipDefaultTx"`                            //是否跳过默认事务
	CoverLogger     bool   `yaml:"coverLogger"`                              //是否覆盖默认logger
	PreparedStmt    bool   `yaml:"prepareStmt"`                              // 设置SQL缓存
	CloseForeignKey bool   `yaml:"disableForeignKeyConstraintWhenMigrating"` // 禁用外键约束
	TablePrefix     string `yaml:"tablePrefix"`                              // 表前缀
	SingularTable   bool   `yaml:"singularTable"`
}

type log struct {
	Path       string     `yaml:"path"`
	Level      string     `yaml:"level"`
	FilePrefix string     `yaml:"filePrefix"`
	FileFormat string     `yaml:"fileFormat"`
	OutFormat  string     `yaml:"outFormat"`
	LumberJack lumberJack `yaml:"lumberJack"`
}

// 日志切割
type lumberJack struct {
	MaxSize    int  `yaml:"maxSize"`    //单文件最大容量(单位MB)
	MaxBackups int  `yaml:"maxBackups"` // 保留旧文件的最大数量
	MaxAge     int  `yaml:"maxAge"`     // 旧文件最多保存几天
	Compress   bool `yaml:"compress"`   // 是否压缩/归档旧文件
}

type jwt struct {
	Securt     string        `yaml:"securt"`
	Issure     string        `yaml:"issure"`
	ExpireTime time.Duration `yaml:"expireTime"`
}

type redis struct {
	Addr        string        `yaml:"addr"`
	Password    string        `yaml:"password"`
	DefaultDB   int           `yaml:"defaultDB"`
	DialTimeOut time.Duration `yaml:"deialTimeOut"`
	Enable      bool          `yaml:"enable"`
}
