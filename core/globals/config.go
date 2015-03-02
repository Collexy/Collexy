package globals

var Conf Config

type Config struct {
    DbName string `json:"db_name,omitempty"`
    DbUser string `json:"db_user,omitempty"`
    DbPassword string `json:"db_password,omitempty"`
    DbHost string `json:"db_host,omitempty"`
    SslMode string `json:"ssl_mode,omitempty"`
}