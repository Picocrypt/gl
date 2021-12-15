//go:build gen

//go:generate glow generate -out=./v3.2-core/gl/ -api=gl -version=3.2 -profile=core -xml=../glow/xml/ -tmpl=../glow/tmpl/

// This is an empty pseudo-package with the sole purpose of containing go generate directives
// that generate all gl binding packages inside this repository.
package gl
