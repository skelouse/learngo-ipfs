package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	shell "github.com/ipfs/go-ipfs-api"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type webPlusSH struct {
	web *echo.Echo
	sh  *shell.Shell
}

type getRequest struct {
	Password string `json:"password" form:"password" query:"password"`
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
func allowOrigin(origin string) (bool, error) {
	return true, nil
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
	web.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOriginFunc: allowOrigin,
		AllowMethods:    []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowHeaders:    []string{echo.MIMEMultipartForm},
	}))
	web.POST("/add", engine.post)

	// Start server
	web.Logger.Fatal(web.Start(":1323"))
}

// GET Handler
func (w *webPlusSH) get(c echo.Context) error {
	cid := c.Param("CID")
	resp := verifyCID(cid)
	if resp != "" {
		return echo.NewHTTPError(http.StatusConflict, []byte(resp))
	}
	request := new(getRequest)
	if err := c.Bind(request); err != nil {
		return echo.NewHTTPError(http.StatusConflict, err)
	}

	w.sh.Get(cid, OUTFILE)
	data := decrypt(OUTFILE, request.Password)
	err := ioutil.WriteFile("./temp.txt", data, 0644)
	check(err)
	return c.File("./temp.txt")
}

// POST Handler
func (w *webPlusSH) post(c echo.Context) error {
	file := c.FormValue("file")
	password := c.FormValue("password")
	send_data := encrypt([]byte(file), password)
	reader := bytes.NewReader(send_data)
	cid, err := w.sh.Add(reader)
	check(err)

	return c.String(http.StatusOK, cid)
}

func encrypt(data []byte, password string) []byte {
	key_empty := make([]byte, 32)
	copy(key_empty[:len(password)], []byte(password))

	block, err := aes.NewCipher(key_empty)
	check(err)

	gcm, err := cipher.NewGCM(block)
	check(err)

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		log.Fatal(err)
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)

	return ciphertext

}

func decrypt(outfile string, password string) []byte {
	key_empty := make([]byte, 32)
	copy(key_empty[:len(password)], []byte(password))

	ciphertext, err := ioutil.ReadFile(outfile)
	check(err)

	// println(string(ciphertext))  // Shows the encrypted text

	block, err := aes.NewCipher(key_empty)
	check(err)

	gcm, err := cipher.NewGCM(block)
	check(err)

	nonce := ciphertext[:gcm.NonceSize()]
	ciphertext = ciphertext[gcm.NonceSize():]

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	check(err)

	return plaintext

}
