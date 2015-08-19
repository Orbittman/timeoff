package data

import (
	"reflect"
	
	"net/http"
	
	"appengine"
	"appengine/datastore"
)

func Put(r *http.Request, dataObject IDataObject) error {
	c := appengine.NewContext(r)
	
	var entity = reflect.TypeOf(dataObject).Elem().Name()
	k := datastore.NewKey(c, entity, dataObject.DataKey(), 0, nil)
	_, err := datastore.Put(c, k, dataObject)
	
	if err != nil {
		c.Infof("Error: %s : %s : %s: %s", err.Error(), reflect.TypeOf(dataObject), entity, dataObject.DataKey())
	}
	
	return err
}