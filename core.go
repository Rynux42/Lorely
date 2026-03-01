package main
import (
  "flag"
  "fmt"
)

func setup() (int, string, string, bool) {
  port := flag.Int("port", 3040, "Port for API Server")
  saves_path := flag.String("saves", "./saves/", "Worlds saves path")
  providers_url := flag.String("providers", "http://localhost:3041", "URL Address for providers microservices")
  debug_mode := flag.Bool("debug", false, "Detailed debug info")

  flag.Parse()

  fmt.Println("==== Lorely Engine Ready ====")
  fmt.Printf("API Port: %d\n", *port)
  fmt.Printf("Providers URL: %s\n", *providers_url)
  fmt.Printf("World saves path: %s\n", *saves_path)

  if *debug_mode {
    fmt.Println("[INF] Debug mode: ON")
  }

  return *port, *saves_path, *providers_url, *debug_mode
}

func main() {
  port, saves_path, providers_url, debug_mode := setup()

}
