// Package util handles generating random data
package util

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

// alphabet Contains all the letters of the alphabet to generate a random string with
const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var siteSuffix = []string{"com", "net", "org", "edu", "gov", "io"}
var imageFileTypes = []string{"jpeg", "jpg", "png", "vector", "ai"}

/* Init is the initializer function that gets ran when the package is initialized
   @params NONE
   @return NONE
*/
func init() {
	rand.Seed(time.Now().UnixNano()) // Generate the random seed
}

/*RandomInt Generates a random int64 between the min and max
@param min int64 The minimum range for a random int64
@param max int64 The maximum range for a random int64
@return int64 The random int64 that was generated
*/
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

/*RandomString Generates a random string of a specified length
@param n int length of the generated string
@return string Returns the generated string
*/
func RandomString(n int) string {
	var stringBuilder strings.Builder

	length := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(length)] // Grab a random character
		stringBuilder.WriteByte(c)
	}

	return stringBuilder.String() // return the string
}

/*RandomName generates a random name
@params NONE
@return string Returns the random name generated
*/
func RandomName() string {
	return RandomString(10)
}

/*RandomAddress Generates a random address
@params NONE
@return Returns the randomly generated address
*/
func RandomAddress() string {
	return fmt.Sprintf("%d %s %s", RandomInt(10, 999), RandomString(15), RandomString(15))
}

/*RandomEmail Generates a random email address
@params NONE
@return The randomly generated email address
*/
func RandomEmail() string {
	return fmt.Sprintf("%s@%s.%s", RandomString(10), RandomString(5), siteSuffix[rand.Intn(len(siteSuffix))])
}

/*RandomPostal Generates a random postal code
@params NONE
@return string The random postal code generated
*/
func RandomPostal() string {
	return fmt.Sprintf("%d", RandomInt(10000, 99999))
}

/*RandomSKU generates a random SKU
@param none
@return string Generated SKU
*/
func RandomSKU() string {
	return fmt.Sprintf("%d", RandomInt(1000000000000000, 9999999999999999))
}

/*RandomFloat2 Generates a random float rounded to 2 decimal places
@params NONE
@return float64 The generated float
*/
func RandomFloat2() float64 {
	return math.Round(rand.Float64()*100) / 100
}

/*RandomImage Generates a random image file name
@param NONE
@return string Generated image file name
*/
func RandomImage() string {
	return fmt.Sprintf("%s.%s", RandomString(4), imageFileTypes[rand.Intn(len(imageFileTypes))])
}

/*RandomBool Generates a random Bool
@params NONE
@return bool The generated bool
*/
func RandomBool() bool {
	n := RandomInt(1, 20)
	if n%2 == 0 {
		return true
	}

	return false
}
