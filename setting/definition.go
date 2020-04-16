package setting

type ETCDConfig interface {
	GetEnable() bool
	GetPort() int
	GetHost() string
}

type StorageConfig interface {
	GetEnable() bool
	GetType() string
	GetURL() string
}

type LoggerConfig interface {
	GetEnable() bool
	GetPath() string
	GetLevel() int
}

type MailConfig interface {
	GetEnable() bool
	GetHost() string
	GetPort() int
	GetEmail() string
	GetPassword() string
}

type RedisConfig interface {
	GetEnable() bool
	GetPort() int
	GetHost() string
}

type defaultETCDConfig struct {
	Enable bool   `json:"enable"`
	Host   string `json:"host"`
	Port   int    `json:"port"`
}
type defaultStorageConfig struct {
	Enable bool   `json:"enable"`
	Type   string `json:"type"`
	Url    string `json:"url"`
}
type defaultLoggerConfig struct {
	Enable bool   `json:"enable"`
	Path   string `json:"path"`
	Level  int    `json:"level"`
}
type defaultEmailConfig struct {
	Enable   bool   `json:"enable"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type defaultRedisConfig struct {
	Enable bool   `json:"enable"`
	Host   string `json:"host"`
	Port   int    `json:"port"`
}

func (c defaultETCDConfig) GetEnable() bool { return c.Enable }
func (c defaultETCDConfig) GetPort() int    { return c.Port }
func (c defaultETCDConfig) GetHost() string { return c.Host }

func (c defaultStorageConfig) GetEnable() bool { return c.Enable }
func (c defaultStorageConfig) GetType() string { return c.Type }
func (c defaultStorageConfig) GetURL() string  { return c.Url }

func (c defaultLoggerConfig) GetEnable() bool { return c.Enable }
func (c defaultLoggerConfig) GetPath() string { return c.Path }
func (c defaultLoggerConfig) GetLevel() int   { return c.Level }

func (c defaultEmailConfig) GetEnable() bool     { return c.Enable }
func (c defaultEmailConfig) GetHost() string     { return c.Host }
func (c defaultEmailConfig) GetPort() int        { return c.Port }
func (c defaultEmailConfig) GetEmail() string    { return c.Email }
func (c defaultEmailConfig) GetPassword() string { return c.Password }

func (c defaultRedisConfig) GetEnable() bool { return c.Enable }
func (c defaultRedisConfig) GetPort() int    { return c.Port }
func (c defaultRedisConfig) GetHost() string { return c.Host }
