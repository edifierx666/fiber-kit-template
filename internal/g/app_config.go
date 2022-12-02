package g

type AppConfig struct {
  Appname  string      `json:"appname,omitempty"`
  Server   Server      `json:"server,omitempty"`
  Log      LogCfg      `json:"log" yaml:"log"`
  Jwt      Jwt         `json:"jwt"`
  Database []Database  `json:"database"`
  Wx       WxConfig    `json:"wx"`
  Qiniu    QiniuConfig `json:"qiniu" yaml:"qiniu"`
}
type Server struct {
  Port    int    `json:"port,omitempty"`
  Address string `json:"address,omitempty"`
  Prefork bool   `json:"prefork,omitempty"`
}
type LogCfg struct {
  File bool `json:"file,omitempty" yaml:"file"`
}
type Jwt struct {
  Key string `json:"key"`
}

type Database struct {
  Name   string `json:"name"`
  Type   string `json:"type"`
  Link   string `json:"link"`
  Level  string `json:"level"`
  Stdout bool   `json:"stdout"`
  Debug  bool   `json:"debug"`
}
type WxConfig struct {
  Appid  string `json:"appid,omitempty"`
  Secret string `json:"secret,omitempty"`
}

type QiniuConfig struct {
  AccessKey      string `json:"accessKey,omitempty"`
  SecretKey      string `json:"secretKey,omitempty"`
  CallbackDomain string `json:"callbackDomain,omitempty"`
}

func (a *AppConfig) GetDBCfg(names ...string) *Database {
  name := ""
  if len(names) > 0 {
    name = names[0]
  }
  if name == "" {
    name = "default"
  }
  for _, database := range a.Database {
    if database.Name == name {
      return &database
    }
  }
  return nil
}
