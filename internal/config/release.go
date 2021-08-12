//+build !dev

package config

type Config struct {
	DebugMode bool `env:"DEBUG" envDefault:"false"`
	Verbose   bool `env:"VERBOSE" envDefault:"false"`

	// Web host address
	WebHost string `env:"WEB_HOST" envDefault:"127.0.0.1"`
	WebPort int    `env:"WEB_PORT" envDefault:"8080"`
}
