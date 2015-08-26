// generated by ginger from go generate -- DO NOT EDIT
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/mingderwang/pepper/jsonToGo"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type OnionResource struct {
	db gorm.DB
}

// @Title CreateOnion
// @Description get string by ID
// @Accept  json
// @Param   some_id     path    int     true        "Some ID"
// @Success 201 {object} string
// @Failure 400 {object} APIError "problem decoding body"
// @Router /onion/ [post]
func (tr *OnionResource) CreateOnion(c *gin.Context) {
	var onion Onion

	if c.Bind(&onion) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "problem decoding body"})
		return
	}
	//onion.Status = OnionStatus
	onion.Ginger_Created = int32(time.Now().Unix())

	tr.db.Save(&onion)

	c.JSON(http.StatusCreated, onion)
}

func (tr *OnionResource) GetAllOnions(c *gin.Context) {
	var onions []Onion

	tr.db.Order("ginger__created desc").Find(&onions)

	c.JSON(http.StatusOK, onions)
}

func (tr *OnionResource) GetOnion(c *gin.Context) {
	id, err := tr.getId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "problem decoding id sent"})
		return
	}

	var onion Onion

	if tr.db.First(&onion, id).RecordNotFound() {
		c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
	} else {
		c.JSON(http.StatusOK, onion)
	}
}

func (tr *OnionResource) BuildOnion(c *gin.Context) {
	id, err := tr.getId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "problem decoding id sent"})
		return
	}

	var onion Onion

	if tr.db.First(&onion, id).RecordNotFound() {
		c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
	} else {
		genJsonToGo(onion)
		c.JSON(http.StatusOK, onion)
	}
}

func genJsonToGo(obj Onion) {
	output, _ := jsonToGo.Gen(obj.JsonSchema, obj.TypeName)
	writeFile("./workspace/"+obj.DomainName+"."+obj.TypeName, "onion.go", output)
}

func writeFile(path string, fileName string, stream string) {
	//check path
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, os.ModePerm)
	}

	f, err := os.OpenFile(path+"/"+fileName, os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	log.Print(err)
	n3, err := f.WriteString(stream)
	fmt.Printf("wrote %d bytes\n", n3)
	f.Sync()
}

func (tr *OnionResource) UpdateOnion(c *gin.Context) {
	id, err := tr.getId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "problem decoding id sent"})
		return
	}

	var onion Onion

	if c.Bind(&onion) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "problem decoding body"})
		return
	}
	onion.Ginger_Id = int32(id)

	var existing Onion

	if tr.db.First(&existing, id).RecordNotFound() {
		c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
	} else {
		tr.db.Save(&onion)
		c.JSON(http.StatusOK, onion)
	}

}

func (tr *OnionResource) PatchOnion(c *gin.Context) {
	id, err := tr.getId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "problem decoding id sent"})
		return
	}

	// this is a hack because Gin falsely claims my unmarshalled obj is invalid.
	// recovering from the panic and using my object that already has the json body bound to it.
	var json []Patch

	r := c.Bind(&json)
	if r != nil {
		fmt.Println(r)
	} else {
		if json[0].Op != "replace" && json[0].Path != "/status" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "PATCH support is limited and can only replace the /status path"})
			return
		}
		var onion Onion

		if tr.db.First(&onion, id).RecordNotFound() {
			c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
		} else {
			//onion.Status = json[0].Value

			tr.db.Save(&onion)
			c.JSON(http.StatusOK, onion)
		}
	}
}

func (tr *OnionResource) DeleteOnion(c *gin.Context) {
	id, err := tr.getId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "problem decoding id sent"})
		return
	}

	var onion Onion

	if tr.db.First(&onion, id).RecordNotFound() {
		c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
	} else {
		tr.db.Delete(&onion)
		c.Data(http.StatusNoContent, "application/json", make([]byte, 0))
	}
}

func (tr *OnionResource) getId(c *gin.Context) (int32, error) {
	idStr := c.Params.ByName("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Print(err)
		return 0, err
	}
	return int32(id), nil
}

/**
* on patching: http://williamdurand.fr/2014/02/14/please-do-not-patch-like-an-idiot/
 *
  * patch specification https://tools.ietf.org/html/rfc5789
   * json definition http://tools.ietf.org/html/rfc6902
*/

type Patch struct {
	Op    string `json:"op" binding:"required"`
	From  string `json:"from"`
	Path  string `json:"path"`
	Value string `json:"value"`
}