package options

import "flag"

type Options struct {
	AllowedHosts       Hosts
	Bucket             string
	DataSourceName     string
	MaxHTTPConnections int
	ObjectPrefix       string
	Port               int
	ServiceAccount     ServiceAccount
	Verbose            bool
}

func Parse(args []string) (Options, error) {
	o := Options{}

	fs := flag.NewFlagSet("resizer", flag.ContinueOnError)
	fs.Var(&o.AllowedHosts, "host", `Hosts of the image that is allowed to resize.
         When this value isn't specified, all hosts are allowed.
         Multiple hosts can be specified with:
             $ resizer -host a.com,b.com
             $ resizer -host a.com -host b.com`)
	fs.StringVar(&o.Bucket, "bucket", "", `Bucket name of Google Cloud Storage to upload the resized image.`)
	fs.StringVar(&o.DataSourceName, "dsn", "", `Data source name of database to store resizing information.`)
	fs.IntVar(&o.MaxHTTPConnections, "connections", 0, `Max simultaneous connections to be accepted by server.
         When 0 or less is specified, the number of connections isn't limited.
         `)
	fs.StringVar(&o.ObjectPrefix, "prefix", "resized/", ``)
	fs.IntVar(&o.Port, "port", 80, `Port to be listened.
         `)
	fs.Var(&o.ServiceAccount, "account", `Path to the file of Google service account JSON.`)
	fs.BoolVar(&o.Verbose, "verbose", false, `Verbose output.
         `)

	return o, fs.Parse(args)
}
