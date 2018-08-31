package cfg

type SMTP struct {
	Port     int
	Host     string
	Mail     string
	Password string
}
type SQL struct {
	Password string
	Login    string
	Port     string
	Host     string
	DBName   string
}

var (
	Route                 = "127.0.0.1"
	Port                  = "8002"
	SizeChanForAlarmClock = 10
	SizeChanForResend     = 10
	CountOLXClientWorkers = 10
	Storage               = SQL{
		Password: "sem1920dark",
		Login:    "root",
		Port:     "3306",
		Host:     "127.0.0.1",
		DBName:   "parserms",
	}
	Mail = SMTP{
		Mail:     "bondartsov.igor@gmail.com",
		Password: "ronuunqgywwcwrwm",
		Port:     465,
		Host:     "smtp.gmail.com",
	}

	PublicKey = []byte(`-----BEGIN PUBLIC KEY-----
MIIBCgKCAQEAwEXBRCwisurukRcgKDfTpEHlG0lZOjNgPiS3vDorVv5k8pk6iERM
0Q5Bi9ok9RLEuIuxY10b5ODp5qtIXODhg3a/hNye1gaQ1a2JhixTC0DUxYL0GsaG
lUdGd6I3jYxrSjUGFGCubbcllBFnu4BsLxLcy/3sm/ym5sL3aYgjbjB8j/R5T+RJ
Kn/06FdhhxbjVrOQ+ySCvTzAizF+n7Iu/iiVW+0LrWru5GqnjkDp4h3iF9PQEoea
CLFP+XhMEsNF1cuWpZo4JcZODPyP9uhNOmzXR6C5Fd9nsTfrLm1bggMqZvZTvctO
OiP8d2rkiLV0iPNV8KID/kWiGAWcwJ4bJQIDAQAB
-----END PUBLIC KEY-----
`)
)
