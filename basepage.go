package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"
)

func serveBaseHTML(template *template.Template, w http.ResponseWriter, r *http.Request) error {
	files, err := ioutil.ReadDir("webroot/configs")
	if err != nil {
		return err
	}

	links := []link{}
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".js" {
			config := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))
			title := strings.Title(strings.Replace(strings.Replace(config, "-", " ", -1), "_", " ", -1))
			links = append(links, link{
				URL:   "?config=" + config,
				Title: title,
			})
		}
	}

	return template.Execute(w, links)
}

func parseBasePageTemplate() (*template.Template, error) {
	baseHTML := `
			<!DOCTYPE html>
			<html>
			<head>
			    <title>Flowbro</title>
			</head>
			<body>
			<div class="box">
			  <div class="row header">
			    Flowbro
			  </div>
			  <div class="row content">
			      <ul>
				{{range $i, $e := .}}<li>
				    <a href="{{.Url}}">{{.Title}}</a>
				</li>{{end}}
			      </ul>
			  </div>
			  <div class="row footer">
			    <p><b>footer</b> (fixed height)</p>
			  </div>
			</div>
			<style>
				html,
				body {
				  background-color: #FFF;
			          margin: 0;
			          padding: 0;
			          line-height: 1;
			          font-family: 'Open Sans', 'Verdana', 'sans-serif';
       			          color: white;
				  height: 100%;
				}

				.box {
				  display: flex;
				  flex-flow: column;
				  height: 100%;
				}

				.box .row {
				  flex: 0 1 30px;
				}

				.box .row.header {
				  flex: 0 1 50px;
				  line-height: 50px;
				  font-size: 26px;
				  background-color: #0091EA;
				  padding: 20px;
				}

				.box .row.content {
				  flex: 1 1 auto;
				  color: black;
			    	  background-color: transparent;
			          box-shadow: inset 0px 3px 3px 1px rgba(0,0,0,0.3);
				}

				.box .row.footer {
				  flex: 0 1 40px;
				}

				ul {
				}

				li {
				  list-style: none;
				  padding-left:0;
				  padding: 20px;
				  font-size: 20px;
				}

				a {
				  text-decoration: none;
				  color: rgb(50, 50, 150);
				}

				a:hover {
				  color: rgb(50, 50, 200);
				}
			</style>
			</body>
			</html>
	`

	return template.New("base").Parse(baseHTML)
}