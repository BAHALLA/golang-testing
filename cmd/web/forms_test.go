package main

import (
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestFormHas(t *testing.T) {

	form := NewForm(nil)

	has := form.Has("some")

	if has {
		t.Errorf("From show has field when it shouldn't")
	}

	postedData := url.Values{}

	postedData.Add("a", "a")

	form = NewForm(postedData)

	has = form.Has("a")

	if !has {
		t.Error("Form show it doesn't have when it should")
	}

}

func TestFormRequired(t *testing.T) {

	r := httptest.NewRequest("POST", "/some", nil)

	form := NewForm(r.PostForm)

	form.Required("a", "b", "c")

	if form.Valid() {
		t.Error("Form shows valid when messing required fields ")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "b")
	postedData.Add("c", "c")

	r = httptest.NewRequest("POST", "/some", nil)

	r.PostForm = postedData
	form = NewForm(r.PostForm)

	form.Required("a", "b", "c")

	if !form.Valid() {
		t.Error("Form shows not valid when it has all required fields")
	}
}

func TestFormCheck(t *testing.T) {
	form := NewForm(nil)

	form.Check(false, "password", "password is required")

	if form.Valid() {
		t.Error("valid return false, when it should return true when calling check")
	}
}

func TestFormGet(t *testing.T) {
	form := NewForm(nil)
	form.Check(false, "password", "password required")

	s := form.Errors.Get("password")

	if len(s) == 0 {
		t.Error("should have an error returned from Get, but it has not!")
	}

	s = form.Errors.Get("some")

	if len(s) != 0 {
		t.Error("should not have an error , but got one")
	}

}
