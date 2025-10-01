package config

import (
	"flag"
	"os"
)

type Config struct {
	URL          string
	Token        string
	Timeout      int
	From         string
	To           string
	Date         string
	Class        string
	Adt          int
	Chd          int
	Count        int
	IsBaggage    bool
	IsDirectOnly bool   // 👉 foydalanuvchi uchun qulay (true/false)
	IsCharter    bool
	Lang         string
	Debug        bool
}

// Load parses CLI flags and environment variables
func Load() *Config {
	cfg := &Config{}

	// URL va Token
	flag.StringVar(&cfg.URL, "url", os.Getenv("API_URL"), "API URL")
	flag.StringVar(&cfg.Token, "token", os.Getenv("API_TOKEN"), "API Token")

	// Umumiy sozlamalar
	flag.IntVar(&cfg.Timeout, "timeout", 30, "request timeout (sec)")

	// Flight parametrlari
	flag.StringVar(&cfg.From, "from", "", "Origin code (IATA)")
	flag.StringVar(&cfg.To, "to", "", "Destination code (IATA)")
	flag.StringVar(&cfg.Date, "date", "", "Flight date (dd.MM.yyyy)")
	flag.StringVar(&cfg.Class, "class", "a", "Booking class (e.g. a, b, c)")
	flag.IntVar(&cfg.Adt, "adt", 1, "Number of adults")
	flag.IntVar(&cfg.Chd, "chd", 0, "Number of children")
	flag.IntVar(&cfg.Count, "count", 20, "Max results")

	// Filterlar
	flag.BoolVar(&cfg.IsBaggage, "baggage", false, "With baggage")
	flag.BoolVar(&cfg.IsDirectOnly, "direct", false, "Direct flights only")
	flag.BoolVar(&cfg.IsCharter, "charter", false, "Include charter flights")

	// Qo‘shimcha
	flag.StringVar(&cfg.Lang, "lang", "en", "Language (en/ru/uz)")
	flag.BoolVar(&cfg.Debug, "debug", false, "Debug mode")

	flag.Parse()

	return cfg
}
