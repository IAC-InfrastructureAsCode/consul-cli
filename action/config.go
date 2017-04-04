package action

import (
	"flag"
)

type config struct {
	consul
	output output
	raw raw
	service service
	check check
}

type service struct {
        id          string
        tags        []string
        address     string
        port        int
        overrideTag bool
}

type check struct {
      id string
      http string
      script string
      ttl string
      interval string
      notes string
      dockerId string
      shell string
      deregCrit string
      skipVerify bool
}

var gConfig config
var gFlags *flag.FlagSet

func init() {
	gFlags = flag.NewFlagSet("consul-cli", flag.ExitOnError)
	gFlags.StringVar(&gConfig.address, "consul", "", "Consul address:port")
	gFlags.BoolVar(&gConfig.ssl, "ssl", false, "Use HTTPS when talking to Consul")
	gFlags.BoolVar(&gConfig.sslVerify, "ssl-verify", true, "Verify certificates when connecting via SSL")
	gFlags.StringVar(&gConfig.sslCert, "ssl-cert", "", "Path to an SSL client certificate for authentication")
	gFlags.StringVar(&gConfig.sslKey, "ssl-key", "", "Path to an SSL client certificate key for authentication")
	gFlags.StringVar(&gConfig.sslCaCert, "ssl-ca-cert", "", "Path to a CA certificate file to validate the Consul server")
	gFlags.StringVar(&gConfig.auth, "auth", "", "The HTTP basic authentication username (and optional password) separated by a colon")
	gFlags.StringVar(&gConfig.token, "token", "", "The Consul ACL token")
	gFlags.StringVar(&gConfig.tokenFile, "token-file", "", "Path to file containing Consul ACL token")
}

func (c *config) addStaleFlag(f *flag.FlagSet) {
	f.BoolVar(&c.stale, "stale", false, "Allow any agent to service the request")
}


func (c *config) addOutputFlags(f *flag.FlagSet, kvFlag bool) {
	c.output.addOutputFlags(f, kvFlag)
}

func (c *config) addServiceFlags(f *flag.FlagSet) {
        f.StringVar(&c.service.id, "id", "", "Service id")
        f.Var(newStringSliceValue(&c.service.tags), "tag", "Service tag. Multiple tags are allowed")
        f.StringVar(&c.service.address, "address", "", "Service address")
        f.IntVar(&c.service.port, "port", 0, "Service port")
        f.BoolVar(&c.service.overrideTag, "override-tag", false, "Disable anti-entropy for this service's tags")
}

func (c *config) addCheckFlags(f *flag.FlagSet) {
      f.StringVar(&c.check.id, "id", "", "Service id")
        f.StringVar(&c.check.http, "http", "", "A URL to GET every interval")
        f.StringVar(&c.check.script, "script", "", "A script to run every interval")
        f.StringVar(&c.check.ttl, "ttl", "", "Fail if TTL expires before service checks in")
        f.StringVar(&c.check.interval, "interval", "", "Interval between checks")
        f.StringVar(&c.check.notes, "notes", "", "Description of the check")
        f.StringVar(&c.check.dockerId, "docker-id", "", "Docker container ID")
        f.StringVar(&c.check.shell, "shell", "", "Shell to use inside docker container")
        f.StringVar(&c.check.deregCrit, "deregister-crit", "", "Deregister critical service after this interval")
        f.BoolVar(&c.check.skipVerify, "skip-verify", false, "Skip TLS verification for HTTP checks")
}

func (c *config) addRawFlag(f *flag.FlagSet) {
	f.StringVar(&c.raw.data, "raw", "", "Raw JSON data for upload")
}

func (c *config) Output(v interface{}) error {
	return c.output.output(v)
}

func (c *config) OutputKv(v interface{}) error {
	return c.output.outputKv(v)
}

