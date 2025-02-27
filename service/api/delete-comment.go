package api

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) uncommentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var deletedComment Comment
	commentid, err1 := strconv.Atoi(ps.ByName("commentid"))

	if err1 != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	authToken := r.Header.Get("Authorization")
	authToken = strings.Split(authToken, " ")[1]
	bool, err := rt.db.CheckAuthToken(authToken)

	if bool {

		err3 := rt.db.DeleteComment(commentid)

		if errors.Is(err3, database.ErrCommentDoesNotExist) {

			w.WriteHeader(http.StatusNotFound)
			return
		} else if err3 != nil {

			ctx.Logger.WithError(err3).WithField("commentID", deletedComment).Error("can't delete the comment")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)

	} else {
		ctx.Logger.WithError(err).Error("Uncorrect token")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
