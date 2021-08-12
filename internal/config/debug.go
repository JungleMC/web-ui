//+build dev

package config

type Config struct {
	DebugMode bool `env:"DEBUG" envDefault:"true"`
	Verbose   bool `env:"VERBOSE" envDefault:"true"`

	// Web host address
	WebHost string `env:"WEB_HOST" envDefault:"127.0.0.1"`
	WebPort int    `env:"WEB_PORT" envDefault:"8080"`
}
