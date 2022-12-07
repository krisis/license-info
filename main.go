package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/minio/pkg/licverifier"
)

var pemBytes = []byte(`-----BEGIN PUBLIC KEY-----
MHYwEAYHKoZIzj0CAQYFK4EEACIDYgAEaK31xujr6/rZ7ZfXZh3SlwovjC+X8wGq
qkltaKyTLRENd4w3IRktYYCRgzpDLPn/nrf7snV/ERO5qcI7fkEES34IVEr+2Uff
JkO2PfyyAYEO/5dBlPh1Undu9WQl6J7B
-----END PUBLIC KEY-----`)

type licInfo licverifier.LicenseInfo

func main() {
	var licpath string
	var verbose bool
	flag.StringVar(&licpath, "file", "", "path to license-key file")
	flag.StringVar(&licpath, "f", "", "path to license-key file")
	flag.BoolVar(&verbose, "v", false, "verbosity")

	flag.Parse()

	if licpath == "" {
		log.Fatalln("Provide the path to the license-key file")
	}

	curdir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	license, err := os.ReadFile(path.Join(curdir, licpath))
	if err != nil {
		log.Fatalln("failed to read license-key file", err)
	}

	var v *licverifier.LicenseVerifier
	v, err = licverifier.NewLicenseVerifier(pemBytes)
	if err != nil {
		log.Fatalln(err)
	}

	var info licverifier.LicenseInfo
	info, err = v.Verify(string(license))
	if err != nil {
		log.Fatalln(err)
	}
	if verbose {
		var b []byte
		b, err = json.MarshalIndent(info, "", " ")
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(string(b))
	} else {
		fmt.Printf("License expires on %s\n", info.ExpiresAt)
	}
}
