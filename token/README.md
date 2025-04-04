# package token

`package token` provides a set of interfaces for service authorization through [JSON Web Tokens](https://jwt.io/).

## Usage

New takes a key function and an expected signing method and returns an
`Sign`.

```go
import (
    "context"
    "time"

    "github.com/skortech/st-kit/token"
)

func main() {
    // Use it as singleton 
    sign := New("secret")

    accessToken, err := sign.Generate(
        context.Background(),
        Claims{
            UserID:   "SL20220601",
            Phone:    "8001001050",
            DeviceID: "SJ13131PGW",
        },
        token.WithDuration(60 * time.Minutes)
    )
    if err != nil {
        // Handle the error
    }
    
    claims,err:=token.Verify(context.Background(),accessToken)
    if err != nil {
        // Handle the error
    }
    print(claims)
}
```
