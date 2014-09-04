
package rest

import (
	"github.com/gin-gonic/gin"
//	"github.com/whyrusleeping/askeecs/server/kvstore"
	"labix.org/v2/mgo/bson"
)

type UserService struct {
	db *Database
}

type User struct {
	ID bson.ObjectId `json:"_id,omitempty"`
	Username string `json:"username"`
	Password string `json:"-"`
	Public   string `json:"public"`
}

func (this *User) GetID() bson.ObjectId {
	return this.ID
}

func (this *User) New() I {
	return new(User)
}

func (p *UserService) Bind (app *gin.Engine) {
	p.db.Collection("Users", new(User))
	app.GET("/users", p.ListUsers)
	app.POST("/users", p.CreateUser)
}

func (p *UserService) ListUsers (c *gin.Context) {
	list := p.db.collections["Users"].FindWhere(bson.M{})
	if list == nil {
		c.JSON(404, gin.H{"message": "no records found"})
		return
	}

	c.JSON(200, list)
}

func (p *UserService) CreateUser(c *gin.Context) {
	var user User
	var err error

	if c.Bind(&user) {
		user.ID = bson.NewObjectId()
		err = p.db.collections["Users"].Save(&user)

		if err != nil {
			panic(err)
			c.JSON(500, gin.H{"message": "error making user"})
			return
		}

		c.JSON(200, user)

	}
}


