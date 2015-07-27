package data

import (
	"net/http"
	"appengine"
	"appengine/datastore"
)

func Put(r *http.Request, dataObject IDataObject) error {
	c := appengine.NewContext(r)

	_, err := datastore.Put(c, datastore.NewIncompleteKey(c, dataObject.DataKey(), nil), dataObject)
	
	return err
}