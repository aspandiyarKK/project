package rest

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Router struct {
	log     *logrus.Entry
	router  *gin.Engine
	app     App
	version string
}

type App interface {
	StoreDate(lastVisit time.Time) error
	GetDate() (time.Time, error)
}

func NewRouter(log *logrus.Logger, version string, app App) *Router {
	r := &Router{
		log:     log.WithField("component", "router"),
		router:  gin.Default(),
		app:     app,
		version: version,
	}
	r.router.GET("/version", r.getVersion)
	r.router.GET("/time", r.getTime)
	r.router.GET("/lastVisit", r.getLastVisit)
	return r
}

func (r *Router) Run(addr string) error {
	return r.router.Run(addr)
}

func (r *Router) getVersion(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"version": r.version,
	})
	if err := r.app.StoreDate(time.Now()); err != nil {
		r.log.Errorf("failed to store date: %v", err)
	}
}

func (r *Router) getTime(c *gin.Context) {
	t := time.Now()
	c.JSON(http.StatusOK, gin.H{
		"currentTime": t.Format(time.RFC3339),
	})
	if err := r.app.StoreDate(time.Now()); err != nil {
		r.log.Errorf("failed to store date: %v", err)
	}
}

func (r *Router) getLastVisit(c *gin.Context) {
	t, err := r.app.GetDate()
	if err != nil {
		r.log.Errorf("failed to get date: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to get date",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"lastRequestTime": t.Format(time.RFC3339),
	})
}
