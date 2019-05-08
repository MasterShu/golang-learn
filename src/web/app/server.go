package app

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"net/http"
	"path"
	"strconv"
)

type Post struct {
	gorm.Model
	Content string
	Author  string `gorm:"not null"`
}

func HandleRequest(w http.ResponseWriter, req *http.Request) {
	var err error
	switch req.Method {
	case http.MethodGet:
		err = handleGet(w, req)
	case http.MethodPost:
		err = handlePost(w, req)
	case http.MethodPut:
		err = handlePut(w, req)
	case http.MethodDelete:
		err = handleDelete(w, req)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleDelete(w http.ResponseWriter, req *http.Request) (err error) {

	id, err := strconv.Atoi(path.Base(req.URL.Path))
	if err != nil {
		return err
	}
	post, err := retrieve(id)
	if err != nil {
		return err
	}
	err = post.delete()
	if err != nil {
		return err
	}
	w.WriteHeader(http.StatusOK)
	return nil
}

func handlePut(w http.ResponseWriter, req *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(req.URL.Path))
	if err != nil {
		return err
	}
	post, err := retrieve(id)
	if err != nil {
		return err
	}

	lens := req.ContentLength
	body := make([]byte, lens)
	_, _ = req.Body.Read(body)
	_ = json.Unmarshal(body, &post)
	err = post.update()
	if err != nil {
		return err
	}
	w.WriteHeader(http.StatusOK)
	return nil
}

func handlePost(writer http.ResponseWriter, req *http.Request) error {
	var err error
	lens := req.ContentLength
	body := make([]byte, lens)
	_, _ = req.Body.Read(body)
	var post Post
	err = json.Unmarshal(body, &post)
	if err != nil {
		return err
	}
	err = post.create()
	writer.WriteHeader(http.StatusOK)
	return nil
}

func handleGet(writer http.ResponseWriter, req *http.Request) error {
	id, err := strconv.Atoi(path.Base(req.URL.Path))
	if err != nil {
		return err
	}

	post, err := retrieve(id)

	if err != nil {
		return err
	}
	fmt.Printf("%+v", post)
	output, err := json.MarshalIndent(&post, "", "\t\t")
	if err != nil {
		return err
	}
	writer.Header().Set("Content-Type", "application/json")
	_, err = writer.Write(output)
	if err != nil {
		return err
	}
	return nil
}
