# [CONFIG MANAGEMENT](https://skorlife.atlassian.net/wiki/spaces/EF/pages/58425345/Configuration+Management)

## Installation

To install run the following command inside your project directory

```sh
go get github.com/skortech/st-kit/config
```

## Vendoring to your project (recommended for all projects)

```sh
git clone ssh://git@github.com/skorlife/st-kit.git
```

## Usage Example

### Config Initialization

```go
package main

import (
	"context"

	"github.com/skortech/st-kit/config"
)

func main() {
	conf := config.New()
	value, err := conf.Parameter(context.Background(), "password", config.WithDefault("N/A"))
}
```
