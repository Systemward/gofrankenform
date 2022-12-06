package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os/exec"

    "github.com/hashicorp/terraform/terraform"
)

func main() {
    http.HandleFunc("/deploy", deployResources)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func deployResources(w http.ResponseWriter, r *http.Request) {
    // Parse the request body
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Unmarshal the request body into a map
    var data map[string]interface{}
    if err := json.Unmarshal(body, &data); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Get the resources to deploy from the request data
    resources, ok := data["resources"].([]interface{})
    if !ok {
        http.Error(w, "Invalid request data", http.StatusBadRequest)
        return
    }

    // Use the Terraform package to deploy the specified resources
    tf := terraform.NewTerraform()
    tf.Init()
    tf.Apply(resources)

    // Return a success response
    fmt.Fprintln(w, "Resources deployed successfully")
}
