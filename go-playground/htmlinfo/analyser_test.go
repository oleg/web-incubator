package main

import (
	"reflect"
	"testing"
)

func Test_HtmlVersion(t *testing.T) {
	analyser := RegexPageAnalyser{}

	info, _ := analyser.Analyse(`<!DOCTYPE HTML>`)

	expectedVersion := "HTML5"
	if info.HtmlVersion != expectedVersion {
		t.Errorf("wrong version %q, expected %q", info.HtmlVersion, expectedVersion)
	}
}

func Test_Title(t *testing.T) {
	analyser := RegexPageAnalyser{}

	info, _ := analyser.Analyse(`
<title>Page Analyser</title>
`)

	expectedTitle := "page analyser"
	if info.PageTitle != expectedTitle {
		t.Errorf("wrong version %q, expected %q", info.PageTitle, expectedTitle)
	}
}

func Test_Headers(t *testing.T) {
	analyser := RegexPageAnalyser{}

	info, _ := analyser.Analyse(`
<h1>Title</h1>
<h2>Subtitle 1</h2>
<h2>Subtitle 2</h2>
`)

	expectedHeaders := map[string]int{"h1": 1, "h2": 2, "h3": 0, "h4": 0, "h5": 0, "h6": 0}
	if !reflect.DeepEqual(info.Headers, expectedHeaders) {
		t.Errorf("wrong headers %v, expected %v", info.Headers, expectedHeaders)
	}
}

func Test_Links(t *testing.T) {
	analyser := RegexPageAnalyser{}

	info, _ := analyser.Analyse(`
<a href="/">Home</a>
<a href="#about">About</a>
<a href="http://google.com">Google</a>
<a href="https://www.home24.de/">Home24</a>
<a class="primary-link nav-li-link" href="//www.sony.de/" title="sony sites">
	<span class="fi fonticon-10-square-neg-plus" aria-hidden="true"></span>
	<span>sony sites</span>
</a>
`)

	expectedInternalLinks := 3
	if info.InternalLinks != expectedInternalLinks {
		t.Errorf("wrong number of internal links %v, expected %v", info.InternalLinks, expectedInternalLinks)
	}
	expectedExternalLinks := 2
	if info.ExternalLinks != expectedExternalLinks {
		t.Errorf("wrong number of external links %v, expected %v", info.ExternalLinks, expectedExternalLinks)
	}
}
