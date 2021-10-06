package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	shell "github.com/ipfs/go-ipfs-api"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type webPlusSH struct {
	web *echo.Echo
	sh  *shell.Shell
}

func check(e error) {
	if e != nil {
		println("Uhh ohh")
		panic(e)
	}
}

var KEYPATH string = "./keys"
var OUTFILE string = "./outfile.txt"
var HOST string = "localhost:5001"
var CRYPTKEY string = "This is a bad thing to do"

func newEngine() *webPlusSH {
	web := echo.New()
	sh := shell.NewShell(HOST)
	// _ = os.Mkdir(KEYPATH, os.ModePerm)
	println("Started ipfs shell on %s", HOST)
	return &webPlusSH{web, sh}
}

func verifyCID(CID string) string {
	var resp string = ""
	if len(CID) != 46 {
		resp += (CID + " is not 46 characters\n")
	}
	if CID[:2] != "Qm" {
		resp += (CID + " does not start with Qm")
	}
	return resp
}

func main() {

	engine := newEngine()

	web := *engine.web
	//sh := *engine.sh

	// Middleware
	web.Use(middleware.Logger())
	web.Use(middleware.Recover())

	// Routes
	web.GET("/get/:CID", engine.get, middleware.RemoveTrailingSlash())
	web.POST("/add", engine.post)

	// Start server
	web.Logger.Fatal(web.Start(":1323"))
}

// GET Handler
func (w *webPlusSH) get(c echo.Context) error {
	cid := c.Param("CID")
	password := c.Param("password")
	resp := verifyCID(cid)
	if resp != "" {
		return c.String(http.StatusConflict, resp)
	}
	w.sh.Get(cid, OUTFILE)
	data := decrypt(OUTFILE, password)
	return c.String(http.StatusOK, data)
}

// POST Handler
func (w *webPlusSH) post(c echo.Context) error {
	password := c.Param("password")
	data := c.Param("data")
	send_data := encrypt(data, password)
	cid, err := w.sh.Add(strings.NewReader(send_data))
	check(err)

	return c.String(http.StatusOK, cid)

}

func encrypt(data string, password string) string {
	plaintext, err := ioutil.ReadFile(data)
	check(err)

	block, err := aes.NewCipher([]byte(password))
	check(err)

	gcm, err := cipher.NewGCM(block)
	check(err)

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		log.Fatal(err)
	}
	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)
	encoded := base64.StdEncoding.EncodeToString(ciphertext)

	return string(encoded)

}

func decrypt(outfile string, password string) string {

	ciphertext, err := ioutil.ReadFile(outfile)
	check(err)

	block, err := aes.NewCipher([]byte(password))
	check(err)

	gcm, err := cipher.NewGCM(block)
	check(err)

	nonce := ciphertext[:gcm.NonceSize()]
	ciphertext = ciphertext[gcm.NonceSize():]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	check(err)
	println(plaintext)
	decoded, err := base64.StdEncoding.DecodeString(string(plaintext))
	check(err)

	return string(decoded)

}

// // NewContext returns a Context instance.
// func (e *Echo) NewContext(r *http.Request, w http.ResponseWriter) Context {
// 	return &context{
// 		request:  r,
// 		response: NewResponse(w, e),
// 		store:    make(Map),
// 		echo:     e,
// 		pvalues:  make([]string, *e.maxParam),
// 		handler:  NotFoundHandler,
// 	}
// }
