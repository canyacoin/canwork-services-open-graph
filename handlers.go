package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	humanize "github.com/dustin/go-humanize"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"gitlab.com/canya-com/shared/data-structures/canwork"
)

var (
	host      = mustGetenv("DEFAULT_HOST")
	projectID = mustGetenv("GCP_PROJECT_ID")
)

func jobHandler(c *gin.Context) {
	var (
		ctx  = context.Background()
		slug = c.Param("slug")
		ua   = c.Request.Header.Get("User-Agent")
		job  canwork.JobDocument
		user canwork.UserDocument
		// default value
		ogTitle       = "Hire professional freelancers from around the globe to complete your project or job"
		ogDescription = "CanWork is a global platform used to hire affordable, quality freelancers from around the globe to do the work you need on demand. It allows on demand remote project management with multiple flexible payments options including crypto payments"
		ogImage       = fmt.Sprintf("%sassets/img/CanYa_OpenGraph_CanWork.png", host)
		// url
		isFromFb    = strings.HasPrefix(ua, "facebookexternalhit")
		ogURL       = host
		redirectURL = host
	)

	jobdata := getSlugSnap(ctx, "public-jobs", slug)
	mapstructure.Decode(jobdata, &job)

	if job.Slug != "" {
		uid := job.ClientID
		userref := firestoreClient.Doc(fmt.Sprintf("users/%s", uid))
		usersnap, err := userref.Get(ctx)
		if err != nil {
			logger.Errorf("unable to retrieve user document for ID: %s error was: %s", uid, err.Error())
		}
		mapstructure.Decode(usersnap.Data(), &user)
	}

	if job.Information.Title != "" {
		ogTitle = job.Information.Title
		ogDescription = job.Information.Description
		ogImage = user.Avatar.URI
		ogURL = fmt.Sprintf("%sjobs/public/%s", host, slug)
		redirectURL = ogURL
		if isFromFb {
			ogURL = fmt.Sprintf("https://open-graph-dot-%s.appspot.com/job/%s", projectID, slug)
		}
	}

	c.HTML(http.StatusOK, "meta.html",
		gin.H{
			"ogTitle":       ogTitle,
			"ogType":        "article",
			"ogDescription": ogDescription,
			"ogImage":       ogImage,
			"ogURL":         ogURL,
			"redirectURL":   redirectURL,
		})
}

func profileHandler(c *gin.Context) {
	var (
		ctx  = context.Background()
		slug = c.Param("slug")
		ua   = c.Request.Header.Get("User-Agent")
		user canwork.UserDocument
		// default value
		ogTitle       = "Hire professional freelancers from around the globe to complete your project or job"
		ogDescription = "CanWork is a global platform used to hire affordable, quality freelancers from around the globe to do the work you need on demand. It allows on demand remote project management with multiple flexible payments options including crypto payments"
		ogImage       = fmt.Sprintf("%sassets/img/CanYa_OpenGraph_CanYa.png", host)
		// url
		isFromFb    = strings.HasPrefix(ua, "facebookexternalhit")
		ogURL       = host
		redirectURL = host
	)

	userdata := getSlugSnap(ctx, "users", slug)
	mapstructure.Decode(userdata, &user)

	if user.Name != "" {
		ogTitle = user.Name
		ogDescription = user.Bio
		ogImage = user.Avatar.URI
		ogURL = fmt.Sprintf("%sprofile/%s", host, slug)
		redirectURL = ogURL
		if isFromFb {
			ogURL = fmt.Sprintf("https://open-graph-dot-%s.appspot.com/profile/%s", projectID, slug)
		}
	}
	c.HTML(http.StatusOK, "meta.html",
		gin.H{
			"ogTitle":       ogTitle,
			"ogType":        "profile",
			"ogDescription": ogDescription,
			"ogImage":       ogImage,
			"ogURL":         ogURL,
			"redirectURL":   redirectURL,
		})
}

func statusHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":         "OK",
		"serviceID":      serviceID,
		"serviceStarted": humanize.Time(startedAt),
	})
}
