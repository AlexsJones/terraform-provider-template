This example would work with this kind of tf file as a provider

`go build -o terraform-provider-engineer`

```
provider "engineer" {}

resource "host" "example_host" {
  name = "test"
}
```
