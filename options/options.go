package options

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/hnakamur/zap-ltsv"
	"go.uber.org/zap"
)

const (
	EnvGoogleAuthJSON = "GOOGLE_AUTH_JSON"

	EnvAccount     = "RESIZER_ACCOUNT"
	EnvBucket      = "RESIZER_BUCKET"
	EnvConnections = "RESIZER_CONNECTIONS"
	EnvDSN         = "RESIZER_DSN"
	EnvHost        = "RESIZER_HOST"
	EnvPort        = "RESIZER_PORT"
	EnvPrefix      = "RESIZER_PREFIX"
	EnvVerbose     = "RESIZER_VERBOSE"
	EnvEnviroment  = "ENVIRONMENT"

	FlagAccount     = "account"
	FlagBucket      = "bucket"
	FlagConnections = "connections"
	FlagDSN         = "dsn"
	FlagHost        = "host"
	FlagPort        = "port"
	FlagPrefix      = "prefix"
	FlagVerbose     = "verbose"
	FlagEnviroment  = "enviroment"
)

var (
	Envs = []string{
		EnvAccount,
		EnvBucket,
		EnvConnections,
		EnvDSN,
		EnvHost,
		EnvPort,
		EnvPrefix,
		EnvVerbose,
		EnvEnviroment,
	}
	Flags = []string{
		FlagAccount,
		FlagBucket,
		FlagConnections,
		FlagDSN,
		FlagHost,
		FlagPort,
		FlagPrefix,
		FlagVerbose,
		FlagEnviroment,
	}
	EnvFlagMap = map[string]string{}
)

func init() {
	for i, env := range Envs {
		EnvFlagMap[env] = Flags[i]
	}

	err := ltsv.RegisterLTSVEncoder()
	if err != nil {
		panic(err)
	}
}

type Options struct {
	ServiceAccount     ServiceAccount
	Bucket             string
	MaxHTTPConnections int
	DataSourceName     string
	AllowedHosts       Hosts
	Port               int
	ObjectPrefix       string
	Verbose            bool
	Enviroment         string

	Logger *zap.Logger
}

// NewOptions Initialize Options
// - args command line arguments
func NewOptions(args []string) (*Options, error) {
	var err error

	o := &Options{}
	err = o.parse(args)
	if err != nil {
		return nil, err
	}

	// Initialize Logger
	var zapConfig zap.Config
	var zapLogger *zap.Logger
	switch o.Enviroment {
	case "test":
	case "development":
		zapConfig = ltsv.NewDevelopmentConfig()
	default:
		zapConfig = ltsv.NewProductionConfig()
	}
	zapLogger, err = zapConfig.Build()
	if err != nil {
		return nil, err
	}
	o.Logger = zapLogger

	return o, nil
}

func (o *Options) parse(args []string) error {
	if v := os.Getenv(EnvGoogleAuthJSON); v != "" {
		b := []byte(v)
		if err := json.Unmarshal(b, &o.ServiceAccount); err != nil {
			return err
		}
		o.ServiceAccount.Path = filepath.Join(os.TempDir(), "resizer-google-auth.json")
		if err := ioutil.WriteFile(o.ServiceAccount.Path, b, 0644); err != nil {
			return err
		}
	}

	fs := flag.NewFlagSet("resizer", flag.ContinueOnError)
	fs.Var(&o.ServiceAccount, "account", "Path to the file of Google service account JSON.")
	fs.StringVar(&o.Bucket, "bucket", "", "Bucket name of Google Cloud Storage to upload the resized image.")
	fs.IntVar(&o.MaxHTTPConnections, "connections", 0, `Max simultaneous connections to be accepted by server.
         When 0 or less is specified, the number of connections isn't limited.`)
	fs.StringVar(&o.DataSourceName, "dsn", "", `Data source name of database to store resizing information.`)
	fs.Var(&o.AllowedHosts, "host", `Hosts of the image that is allowed to resize.
         When this value isn't specified, all hosts are allowed.
         Multiple hosts can be specified with:
             $ resizer -host a.com,b.com
             $ resizer -host a.com -host b.com`)
	fs.IntVar(&o.Port, "port", 80, "Port to be listened.")
	fs.StringVar(&o.ObjectPrefix, "prefix", "", "Object prefix of Google Cloud Storage.")
	fs.BoolVar(&o.Verbose, "verbose", false, "Verbose output.")
	fs.StringVar(&o.Enviroment, "enviroment", "production", "development or production. In default production")

	for _, env := range Envs {
		flag := EnvFlagMap[env]
		if v := os.Getenv(env); v != "" {
			fs.Set(flag, v)
		}
	}
	return fs.Parse(args)
}
