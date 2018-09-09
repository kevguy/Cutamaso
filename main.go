package main

import (
  "fmt"
  "net/http"
)

// Parse the configuration file '.env', and establish a connection to the DB
func initEnv() {
  fmt.Println("Initializing Environmental Variables")

  // Read .env
  // then you can use for instance, os.Getenv("S3_BUCKET_NAME") to get the values
  fEnvFile := flag.String("env-file", "", "path to environment file")
  mode = flag.String("mode", "", "dev/production mode")
  flag.Parse()

  if *mode == "dev" {
    if *fEnvFile != "" {
      err := util.LoadEnvFile(*fEnvFile)
      if err != nil {
        log.Fatal(err)
      }
    }
  }

  // Init MongoDB

}


func main() {
  fmt.Printf("Hello world")
}
