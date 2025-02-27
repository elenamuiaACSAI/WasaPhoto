package api

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var photo Photo
	file, _, err2 := r.FormFile("photo")

	userid := ps.ByName("userid")

	if err2 != nil {

		w.WriteHeader(http.StatusBadRequest)

		return
	}

	defer file.Close()

	bytes, err3 := io.ReadAll(file)
	if err3 != nil {

		w.WriteHeader(http.StatusBadRequest)

		return
	}

	var base64Encoding string

	// Determine the content type of the image file
	mimeType := http.DetectContentType(bytes)

	// Prepend the appropriate URI scheme header depending
	// on the MIME type
	switch mimeType {
	case "image/jpeg":
		base64Encoding = "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding = "data:image/png;base64,"
	}
	resultEncoding := base64.StdEncoding.EncodeToString(bytes)

	// Append the base64 encoded output
	base64Encoding = base64Encoding + resultEncoding

	// fileBytes is now a []byte containing the contents of the file

	authToken := r.Header.Get("Authorization")
	authToken = strings.Split(authToken, " ")[1]
	bool, err := rt.db.CheckAuthToken(authToken)
	if bool {

		photoRet, err4 := rt.db.PostPhoto(photo.ToDatabasePhoto(userid, base64Encoding))
		if err4 != nil {

			ctx.Logger.WithError(err4).Error("Can't post photo")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(photoRet)
	} else {
		ctx.Logger.WithError(err).Error("Uncorrect token")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
