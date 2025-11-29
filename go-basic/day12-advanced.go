package main

import (
	"errors"
	"fmt"
	"strings"
)

// Struct to hold the pod data for a simulated Pod.
type Pod struct {
	Name      string
	Namespace string
	Status    string
	IP        string
}

// method on the Pod struct
// this method modifies the pod
// we use pointer receiver method so we can change the original pod's data
func (p *Pod) SetStatus(newStatus string) {
	fmt.Printf("Updating status for Pod '%s' ti '%s'", p.Name, newStatus)
	p.Status = newStatus
}

// method on pod struct value receiver
// this method reads the data without changing anything
func (p Pod) IsRunning() bool {
	return p.Status == "Running"
}

// This Client Struct
// This struct will hold the connection data and
// be the receiver for our action methods.
type kubeClient struct {
	ClusterURL string
	// We'll use a map to simulate the pod database.
	// The key is the pod name, and the value is a *pointer*
	// to the Pod, so we can modify the original.
	pods map[string]*Pod
}

// Health check method on kubeClient
// This is out complex method attached to the client
func (c *kubeClient) CheckAllPodsHealth() {
	fmt.Println(strings.Repeat("=", 30))
	fmt.Printf("--- Checking Health for Cluster: %s ---\n", c.ClusterURL)

	if len(c.pods) == 0 {
		fmt.Println("No pod found")
		return
	}

	healthyCount := 0
	for name, pod := range c.pods {
		fmt.Printf("Checking pod: %s (Status: %s)\n", name, pod.Status)
		if pod.IsRunning() {
			fmt.Println("-> HEALTH: OK")
			healthyCount++
		} else {
			fmt.Println("-> HEALTH UNHEALTHY")
		}
	}

	fmt.Printf("--- Summary: %d / %d pods are healthy ---\n ", healthyCount, len(c.pods))
	fmt.Println(strings.Repeat("=", 30))
}

// Method on kubeClient to Find a Pod
func (c *kubeClient) GetPod(name string) (*Pod, error) {
	pod, ok := c.pods[name]
	if !ok {
		// 'errors.New' is a built-in way to create a new error
		return nil, errors.New("Pod not found: " + name)
	}
	return pod, nil // Return the pod pointer and 'nil' for the error
}

func main() {
	// create pod object similar to k8s cluster
	pod1 := &Pod{Name: "api-gateway-1", Namespace: "prod", Status: "Running", IP: "10.0.0.1"}
	pod2 := &Pod{Name: "user-service-1", Namespace: "prod", Status: "Running", IP: "10.0.0.2"}
	pod3 := &Pod{Name: "db-backup-job", Namespace: "default", Status: "Failed", IP: "10.0.0.3"}

	// 2. Create our client and "load" it with the pods
	client := &kubeClient{
		ClusterURL: "https-cluster.my-domain.com",
		pods: map[string]*Pod{
			"api-gateway-1":  pod1,
			"user-service-1": pod2,
			"db-backup-job":  pod3,
		},
	}

	// 3. Run our "complex" method
	client.CheckAllPodsHealth()

	// 4. Let's find a pod and change its status
	fmt.Println("\n=== Making a change... ===")
	targetPod, err := client.GetPod("user-service-1")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		targetPod.SetStatus("Pending")
	}

	client.CheckAllPodsHealth()
}
