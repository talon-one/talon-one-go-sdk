package talon

import "io"

// NamedReader wraps r so that a file upload sends name as the file name.
// Without it, an *os.File is sent under its own name and any other reader
// under the form field name.
func NamedReader(r io.Reader, name string) io.Reader {
	if r == nil {
		return nil
	}
	return namedReader{Reader: r, name: name}
}

type namedReader struct {
	io.Reader
	name string
}

// Name returns the file name sent in the multipart Content-Disposition header.
func (n namedReader) Name() string {
	return n.name
}

// readFormFile reads r fully and returns it as a multipart part named field.
// The file name comes from r's Name method when it has one, otherwise field is
// used. Read errors are returned and r is never closed.
func readFormFile(r io.Reader, field string) (formFile, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return formFile{}, err
	}
	name := field
	if named, ok := r.(interface{ Name() string }); ok && named.Name() != "" {
		name = named.Name()
	}
	return formFile{fileBytes: b, fileName: name, formFileName: field}, nil
}
