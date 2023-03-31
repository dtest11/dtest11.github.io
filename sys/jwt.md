# jwt(JSON Web Token)

1. #### format

	头部、载荷与签名

	- **Header**

		token的类型（“JWT”）和算法名称（比如：HMAC SHA256或者RSA等等）。
		```json
		{
		"alg": "HS256",
		"typ": "JWT"
		}
		```

	- **Payload**

		Payload 部分也是一个JSON对象,JWT 规定了7个官方字段，供选用。
		```erlang
			iss (issuer)：签发人
			exp (expiration time)：过期时间
			sub (subject)：主题
			aud (audience)：受众
			nbf (Not Before)：生效时间
			iat (Issued At)：签发时间
			jti (JWT ID)：编号
		```

	- **Signature**

		为了得到签名部分，你必须有编码过的header、编码过的payload、一个秘钥（这个秘钥只有服务端知道），签名算法是header中指定的那个，然对它们签名即可。


  		`HMACSHA256(base64UrlEncode(header) + "." + base64UrlEncode(payload), secret)`

		算出签名以后，把 Header、Payload、Signature 三个部分拼成一个字符串，每个部分之间用"点"（.）分隔，就可以返回给用户。需要提醒一下：base64是一种编码方式，并非加密方式。

		```tex
		eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c
		```

2. **好处**

	* 无状态

	* 不需要暴露详细的用户信息

	* 支持跨域

3. **example** 
	1. generate

		```go
			func ExampleNewWithClaims_registeredClaims() {
						mySigningKey := []byte("AllYourBase")
		
						// Create the Claims
						claims := &jwt.RegisteredClaims{
							ExpiresAt: jwt.NewNumericDate(time.Unix(1516239022, 0)),
							Issuer:    "test",
						}
		
						token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
						ss, err := token.SignedString(mySigningKey)
						fmt.Printf("%v %v", ss, err)
						//Output: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ0ZXN0IiwiZXhwIjoxNTE2MjM5MDIyfQ.0XN_1Tpp9FszFOonIBpwha0c_SfnNI22DhTnjMshPg8 <nil>
					}
		```
		
		
		
	2. parse
		
		```go
				func ExampleParseWithClaims_customClaimsType() {
					tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJpc3MiOiJ0ZXN0IiwiYXVkIjoic2luZ2xlIn0.QAWg1vGvnqRuCFTMcPkjZljXHh8U3L_qUjszOtQbeaA"
		
					type MyCustomClaims struct {
						Foo string `json:"foo"`
						jwt.RegisteredClaims
					}
		
					token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
						return []byte("AllYourBase"), nil
					})
		
					if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
						fmt.Printf("%v %v", claims.Foo, claims.RegisteredClaims.Issuer)
					} else {
						fmt.Println(err)
					}
		
					// Output: bar test
				}
		```

4. go-zero 实现参考 
[gggg](https://github.com/zeromicro/go-zero/blob/3f492df74e64c47928ba0d00c64a1304c2e8f749/rest/token/tokenparser.go#L29	)	
```go
package token

import (
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/golang-jwt/jwt/v4/request"
	"github.com/zeromicro/go-zero/core/timex"
)

const claimHistoryResetDuration = time.Hour * 24

type (
	// ParseOption defines the method to customize a TokenParser.
	ParseOption func(parser *TokenParser)

	// A TokenParser is used to parse tokens.
	TokenParser struct {
		resetTime     time.Duration
		resetDuration time.Duration
		history       sync.Map
	}
)

// NewTokenParser returns a TokenParser.
func NewTokenParser(opts ...ParseOption) *TokenParser {
	parser := &TokenParser{
		resetTime:     timex.Now(),
		resetDuration: claimHistoryResetDuration,
	}

	for _, opt := range opts {
		opt(parser)
	}

	return parser
}

// ParseToken parses token from given r, with passed in secret and prevSecret.
func (tp *TokenParser) ParseToken(r *http.Request, secret, prevSecret string) (*jwt.Token, error) {
	var token *jwt.Token
	var err error

	if len(prevSecret) > 0 {
		count := tp.loadCount(secret)
		prevCount := tp.loadCount(prevSecret)

		var first, second string
		if count > prevCount {
			first = secret
			second = prevSecret
		} else {
			first = prevSecret
			second = secret
		}

		token, err = tp.doParseToken(r, first)
		if err != nil {
			token, err = tp.doParseToken(r, second)
			if err != nil {
				return nil, err
			}

			tp.incrementCount(second)
		} else {
			tp.incrementCount(first)
		}
	} else {
		token, err = tp.doParseToken(r, secret)
		if err != nil {
			return nil, err
		}
	}

	return token, nil
}

func (tp *TokenParser) doParseToken(r *http.Request, secret string) (*jwt.Token, error) {
	return request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		}, request.WithParser(newParser()))
}

func (tp *TokenParser) incrementCount(secret string) {
	now := timex.Now()
	if tp.resetTime+tp.resetDuration < now {
		tp.history.Range(func(key, value interface{}) bool {
			tp.history.Delete(key)
			return true
		})
	}

	value, ok := tp.history.Load(secret)
	if ok {
		atomic.AddUint64(value.(*uint64), 1)
	} else {
		var count uint64 = 1
		tp.history.Store(secret, &count)
	}
}

func (tp *TokenParser) loadCount(secret string) uint64 {
	value, ok := tp.history.Load(secret)
	if ok {
		return *value.(*uint64)
	}

	return 0
}

// WithResetDuration returns a func to customize a TokenParser with reset duration.
func WithResetDuration(duration time.Duration) ParseOption {
	return func(parser *TokenParser) {
		parser.resetDuration = duration
	}
}

func newParser() *jwt.Parser {
	return jwt.NewParser(jwt.WithJSONNumber())
}
```
[jwt](https://jwt.io/introduction)

