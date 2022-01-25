package token_test

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/kazmerdome/go-graphql-starter/pkg/util/token"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestToken(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Token Test Suite")
}

var _ = Describe("JWT Token Suite", func() {
	/*
	 * Not Expired, Valid JWT With Valid Data
	 */
	Context("Not Expired, Valid JWT With Valid Data", func() {
		var tokenString string
		sessionSecret := "it-is-a-session-secret"
		data := map[string]string{
			"lid": "lid-test-data",
			"uid": "uid-test-data",
		}

		When("try to generate jwt token", func() {
			It("should provide a valid token", func() {
				t, err := token.GenerateJWTToken(data, sessionSecret, 100)
				Expect(err).To(BeNil())
				Expect(t).ToNot(BeEmpty())
				tokenString = t
			})
		})

		When("try to read and verify jwt token", func() {
			It("should provide a valid token", func() {
				d, err := token.VerifyJWTToken(tokenString, sessionSecret)
				Expect(err).To(BeNil())
				Expect(d).To(Equal(data))
				Expect(d["lid"]).To(Equal("lid-test-data"))
				Expect(d["uid"]).To(Equal("uid-test-data"))
			})
		})

		When("try to read and verify jwt token with a wrong sessionSecret", func() {
			It("should throw a 'signature is invalid' error", func() {
				d, err := token.VerifyJWTToken(tokenString, "iamawrongsessionsecretlol")
				Expect(err).ToNot(BeNil())
				Expect(err.Error()).To(Equal("signature is invalid"))
				Expect(d).To(BeEmpty())
			})
		})

		When("try to read and verify jwt token with an empty sessionSecret", func() {
			It("should throw a 'signature is invalid' error", func() {
				d, err := token.VerifyJWTToken(tokenString, "")
				Expect(err).ToNot(BeNil())
				Expect(err.Error()).To(Equal("signature is invalid"))
				Expect(d).To(BeEmpty())
			})
		})
	})

	/*
	 * Expired, Valid JWT (With Valid Data)
	 */
	Context("Expired, Valid JWT (With Valid Data)", func() {
		var tokenString string
		sessionSecret := "it-is-a-session-secret"
		data := map[string]string{}

		When("try to generate an expried jwt token", func() {
			It("should provide a valid expried token", func() {
				t, err := token.GenerateJWTToken(data, sessionSecret, -1)
				Expect(err).To(BeNil())
				Expect(t).ToNot(BeEmpty())
				tokenString = t
			})
		})

		When("try to read and verify expried jwt token", func() {
			It("should throw an error", func() {
				time.Sleep(time.Millisecond * 100)
				d, err := token.VerifyJWTToken(tokenString, sessionSecret)
				Expect(err).ToNot(BeNil())
				Expect(err.Error()).To(Equal("Token is expired"))
				Expect(d).To(BeEmpty())
			})
		})
	})

	/*
	 * Not Expired, Valid JWT With Empty (Nil) Data Claim Field
	 */
	Context("Not Expired, Valid JWT With Empty (Nil) Data Claim Field", func() {
		sessionSecret := "it-is-a-session-secret"
		When("try to generate jwt token", func() {
			It("should throw an error and not provide a token", func() {
				t, err := token.GenerateJWTToken(nil, sessionSecret, 100)
				Expect(err).ToNot(BeNil())
				Expect(err.Error()).To(Equal("data is required"))
				Expect(t).To(BeEmpty())
			})
		})
	})

	/*
	 * Not Expired, Valid JWT Without Data claim field
	 */
	Context("Not Expired, Valid JWT Without Data claim field", func() {
		var tokenString string
		sessionSecret := "it-is-a-session-secret"

		When("try to generate a not CustomClaim type jwt token", func() {
			It("should provide a valid token", func() {
				expireToken := time.Now().Add(time.Hour * time.Duration(100)).Unix()
				claims := jwt.StandardClaims{
					ExpiresAt: expireToken,
					IssuedAt:  time.Now().Unix(),
					Id:        string(primitive.NewObjectID().Hex()),
					Issuer:    "issuer",
				}
				token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
				ts, err := token.SignedString([]byte(sessionSecret))
				tokenString = ts
				Expect(err).To(BeNil())
				Expect(tokenString).ToNot(BeEmpty())
			})
		})

		When("try to read and verify jwt token", func() {
			It("should throw an error", func() {
				time.Sleep(time.Millisecond * 100)
				d, err := token.VerifyJWTToken(tokenString, sessionSecret)
				Expect(err).ToNot(BeNil())
				Expect(err.Error()).To(Equal("invalid stored token"))
				Expect(d).To(BeEmpty())
			})
		})
	})

	/*
	 * Not Expired, Valid JWT With Wrong (int32) Type Data Claim Field (Instead Of map[string]string)
	 */
	Context("Not Expired, Valid JWT With Wrong (int32) Type Data Claim Field (Instead Of map[string]string)", func() {
		var tokenString string
		sessionSecret := "it-is-a-session-secret"

		When("try to generate a custom claim token with az int32 typed data field in it", func() {
			It("should provide a valid token", func() {
				expireToken := time.Now().Add(time.Hour * time.Duration(100)).Unix()
				var data int32 = 99
				claims := struct {
					Data int32
					jwt.StandardClaims
				}{
					data,
					jwt.StandardClaims{
						ExpiresAt: expireToken,
						IssuedAt:  time.Now().Unix(),
						Id:        string(primitive.NewObjectID().Hex()),
						Issuer:    "issuer",
					},
				}
				token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
				ts, err := token.SignedString([]byte(sessionSecret))
				tokenString = ts
				Expect(err).To(BeNil())
				Expect(tokenString).ToNot(BeEmpty())
			})
		})

		When("try to read and verify jwt token", func() {
			It("should throw an error", func() {
				time.Sleep(time.Millisecond * 100)
				d, err := token.VerifyJWTToken(tokenString, sessionSecret)
				Expect(err).ToNot(BeNil())
				Expect(err.Error()).To(Equal("invalid stored token"))
				Expect(d).To(BeEmpty())
			})
		})
	})

	/*
	 * Not Expired, Valid JWT but with a wrong (map[string]interface{}) data claim field (instead of map[string]string)
	 */
	Context("Not Expired, Valid JWT With Wrong (int32) Type Data Claim Field (Instead Of map[string]string)", func() {
		var tokenString string
		sessionSecret := "it-is-a-session-secret"

		When("try to generate a custom claim token with an map[string]interface{} typed data field in it", func() {
			It("should provide a valid token", func() {
				expireToken := time.Now().Add(time.Hour * time.Duration(100)).Unix()
				var data map[string]interface{} = map[string]interface{}{"lid": true}
				claims := struct {
					Data map[string]interface{} `json:"data"`
					jwt.StandardClaims
				}{
					data,
					jwt.StandardClaims{
						ExpiresAt: expireToken,
						IssuedAt:  time.Now().Unix(),
						Id:        string(primitive.NewObjectID().Hex()),
						Issuer:    "issuer",
					},
				}
				token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
				ts, err := token.SignedString([]byte(sessionSecret))
				tokenString = ts
				Expect(err).To(BeNil())
				Expect(tokenString).ToNot(BeEmpty())
			})
		})

		When("try to read and verify jwt token", func() {
			It("should throw an error", func() {
				time.Sleep(time.Millisecond * 100)
				d, err := token.VerifyJWTToken(tokenString, sessionSecret)
				Expect(err).ToNot(BeNil())
				Expect(err.Error()).To(Equal("invalid stored token"))
				Expect(d).To(BeEmpty())
			})
		})
	})
})
