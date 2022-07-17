<div align="center" id="top"> 
  <img src="./.github/app.gif" alt="Annotation" />

&#xa0;

  <!-- <a href="https://annotation.netlify.app">Demo</a> -->
</div>

<h1 align="center">Annotation</h1>

<p align="center">
  <img alt="Github top language" src="https://img.shields.io/github/languages/top/ugurkorkmaz/annotation?color=56BEB8">

  <img alt="Github language count" src="https://img.shields.io/github/languages/count/ugurkorkmaz/annotation?color=56BEB8">

  <img alt="Repository size" src="https://img.shields.io/github/repo-size/ugurkorkmaz/annotation?color=56BEB8">

  <img alt="License" src="https://img.shields.io/github/license/ugurkorkmaz/annotation?color=56BEB8">

 <img alt="Github issues" src="https://img.shields.io/github/issues/ugurkorkmaz/annotation?color=56BEB8" />

  <img alt="Github forks" src="https://img.shields.io/github/forks/ugurkorkmaz/annotation?color=56BEB8" />

  <img alt="Github stars" src="https://img.shields.io/github/stars/ugurkorkmaz/annotation?color=56BEB8" />
</p>

<!-- Status -->

<!-- <h4 align="center">
	🚧  Annotation 🚀 Under construction...  🚧
</h4>

<hr> -->

<p align="center">
  <a href="#dart-about">About</a> &#xa0; | &#xa0; 
  <a href="#sparkles-features">Features</a> &#xa0; | &#xa0;
  <a href="#rocket-technologies">Technologies</a> &#xa0; | &#xa0;
  <a href="#white_check_mark-requirements">Requirements</a> &#xa0; | &#xa0;
  <a href="#checkered_flag-starting">Starting</a> &#xa0; | &#xa0;
  <a href="#memo-license">License</a> &#xa0; | &#xa0;
  <a href="https://github.com/ugurkorkmaz" target="_blank">Author</a>
</p>

<br>

## :dart: About

Routes are first written in the comment lines. Then the library I prepared reads these comment lines and creates automatic route functions.
It provides a Symfony framework or Flask-style routing environment.
It does not cause performance problems because the route function is already created.

You can help me to make it work within routing libraries other than fiber.

## :sparkles: Features
:heavy_check_mark: gofiber/fiber route template;\
:hammer_and_wrench: net/http route template;\
:hammer_and_wrench: gorilla/mux route template;\
:hammer_and_wrench: Fjulienschmidt/httprouter route template;

## :rocket: Technologies

The following tools were used in this project:

- [Golang](https://golang.org/): The main language used in this project.

## :white_check_mark: Requirements
Before you begin, you 🏁 need to have [Git](https://git-scm.com) and [Golang](https://golang.org/) installed.
## :checkered_flag: Starting

```bash
# Go get the dependencies
$ go get github.com/ugurkorkmaz/annotation
```

```go
package main

//go:generate go run github.com/ugurkorkmaz/annotation --directory=./example --package=example --mode=fiber --output=./example/routes.go
```

```bash
# Generate routes
$ go generate
```

## :memo: License

This project is under license from MIT. For more details, see the [LICENSE](LICENSE.md) file.

Made with :heart: by <a href="https://github.com/ugurkorkmaz" target="_blank">Uğur Korkmaz</a>

&#xa0;

<a href="#top">Back to top</a>
