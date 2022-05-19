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
		

[jwt]: https://jwt.io/introduction

