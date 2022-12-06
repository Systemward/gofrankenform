package main

import (
  "bufio"
  "fmt"
  "io"
  "net/http"
  "os/exec"
  "strings"
)

const port = 8000

func main() {
  // Create a request handler for the HTTP server
  http.HandleFunc("/terraform", func(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
      // Set the response code to 200 (OK)
      w.WriteHeader(http.StatusOK)

      // Read the request body
      body := r.Body
      defer body.Close()
      reader := bufio.NewReader(body)

      // Parse the request body to get the Terraform configuration
      // You may need to modify this to fit your specific needs
      config := make(map[string]string)
      for {
        line, err := reader.ReadString('\n')
        if err != nil {
          if err == io.EOF {
            break
          }
          panic(err)
        }
        parts := strings.Split(line, "=")
        if len(parts) == 2 {
          config[parts[0]] = strings.TrimSpace(parts[1])
        }
      }

      // Create a temporary file to store the Terraform configuration
      file, err := os.Create("terraform.tf")
      if err != nil {
        panic(err)
      }
      defer file.Close()

      // Write the Terraform configuration to the file
      for key, value := range config {
        fmt.Fprintf(file, "%s = %s\n", key, value)
      }

      // Run the "terraform init" command to initialize the working directory
      cmd := exec.Command("terraform", "init")
      out, err := cmd.CombinedOutput()
      if err != nil {
        fmt.Fprintf(w, "<h1>Error: %s</h1>", out)
        return
      }

      // Run the "terraform apply" command to create the resources
      cmd = exec.Command("terraform", "apply", "-auto-approve")
      out, err = cmd.CombinedOutput()
      if err != nil {
        fmt.Fprintf(w, "<h1>Error: %s</h1>", out)
        return
      }

      // Write the response body
      fmt.Fprint(w, "<h1>Resources Created</h1>")
    } else {
      // Set the response code to 404 (Not Found)
      w.WriteHeader(http.StatusNotFound)
      // Write the response body
      fmt.Fprint(w, "<h1>404 Not Found</h1>")
    }
  })

  // Create the HTTP server and start listening on the specified port
  fmt.Printf("Serving HTTP on port %d\n", port)
  err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
  if err != nil {
    panic(err)
  }
}