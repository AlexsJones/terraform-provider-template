This example would work with this kind of tf file as a provider

`go build`

```
provider "template" {}

resource "template_host" "example_host" {
  name = "test"
}
```
