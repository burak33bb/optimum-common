package auth

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

// ParseUnverified extracts claims without validating the signature.
// This mirrors your existing pattern but centralizes the quirks.
func ParseUnverified(tokenString string) (*Claims, error) {
	tok, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// No key: we expect signature-related error.
		return nil, nil
	})
	var mc jwt.MapClaims
	if err != nil {
		// Accept signature errors, still read claims.
		if vErr, ok := err.(*jwt.ValidationError); ok &&
			(vErr.Errors&jwt.ValidationErrorSignatureInvalid != 0 ||
				vErr.Errors&jwt.ValidationErrorUnverifiable != 0) {
			if tok != nil {
				if m, ok := tok.Claims.(jwt.MapClaims); ok {
					mc = m
				}
			}
		} else {
			return nil, fmt.Errorf("%w: %v", ErrParsingToken, err)
		}
	} else {
		var ok bool
		mc, ok = tok.Claims.(jwt.MapClaims)
		if !ok {
			return nil, ErrInvalidClaims
		}
	}
	if mc == nil {
		return nil, ErrInvalidClaims
	}

	return fromMap(mc)
}
