// +build gofuzz
package middleware

//Fuzz runs a fuzz test
func Fuzz(data []byte) int {
	config := CSRFConfig{
		TokenLookup: string(data),
	}
	CSRFWithConfig(config)
	return 0
}
