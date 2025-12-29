##  Go program that reads YAML files, edits values, and writes them back to disk:
* Reads a YAML file from disk using os.ReadFile
* Unmarshals the YAML into a Go struct using the gopkg.in/yaml.v3 package
* Modifies values in the struct
* Marshals the struct back to YAML format
* Writes the updated YAML back to disk