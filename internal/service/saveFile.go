package service

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid"
)

const IMAGE_DIR = "./image"

func saveFile(c *gin.Context) (string, error) {
	productId, _ := c.GetPostForm("post_id")
	{
		if len(productId) == 0 {
			productId = "default"
		}
		err := os.MkdirAll(fmt.Sprintf("%s/%s", IMAGE_DIR, productId), os.ModePerm)
		if err != nil {
			return "", err
		}
	}
	path := fmt.Sprintf("%s/%s", IMAGE_DIR, productId)

	form, _ := c.MultipartForm()
	var fileName string
	imgExt := "jpeg"
	for key := range form.File {
		fileName = key
		arr := strings.Split(fileName, ".")
		if len(arr) > 1 {
			imgExt = arr[len(arr)-1]
		}
		continue
	}

	file, _, err := c.Request.FormFile(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	fullFileName := fmt.Sprintf("%s.%s", randomFilename(), imgExt)
	fileOnDisk, err := os.Create(fmt.Sprintf("%s/%s", path, fullFileName))
	if err != nil {
		return "", err
	}
	defer fileOnDisk.Close()

	_, err = fileOnDisk.Write(fileBytes)
	if err != nil {
		return "", err
	}

	return fullFileName, nil
}

func randomFilename() string {
	t := time.Now()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	return strings.ToLower(fmt.Sprintf("%v", ulid.MustNew(ulid.Timestamp(t), entropy)))
}
