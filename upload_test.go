package talon

// Tests for file uploads (Talon.One template customization, see
// sdks/generator/codegen/go in talon-service).

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"testing/iotest"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type receivedUpload struct {
	formName string
	fileName string
	content  string
	err      error
}

// newUploadServer returns a client pointed at a test server that records the
// first multipart part of each request.
func newUploadServer(t *testing.T) (*APIClient, <-chan receivedUpload) {
	t.Helper()
	received := make(chan receivedUpload, 1)
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var got receivedUpload
		defer func() { received <- got }()

		mr, err := r.MultipartReader()
		if err != nil {
			got.err = err
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		part, err := mr.NextPart()
		if err != nil {
			got.err = err
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		b, err := io.ReadAll(part)
		got.formName, got.fileName, got.content = part.FormName(), part.FileName(), string(b)
		if err == nil {
			_, err = io.Copy(io.Discard, r.Body)
		}
		if err != nil {
			got.err = err
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK) // empty body: nothing to decode
	}))
	t.Cleanup(srv.Close)

	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{
		{URL: srv.URL},
	}
	return NewAPIClient(cfg), received
}

// waitUpload returns the upload recorded by the test server, failing the test
// if none arrives in time.
func waitUpload(t *testing.T, received <-chan receivedUpload) receivedUpload {
	t.Helper()
	select {
	case got := <-received:
		return got
	case <-time.After(5 * time.Second):
		require.FailNow(t, "server did not receive the upload")
		return receivedUpload{}
	}
}

// assertNoRequest fails if the server received a request.
func assertNoRequest(t *testing.T, received <-chan receivedUpload) {
	t.Helper()
	select {
	case got := <-received:
		assert.Failf(t, "server received a request, want none", "%+v", got)
	default:
	}
}

const uploadCSV = "value\nCOUP1\nCOUP2\n"

func TestUploadInMemoryReader(t *testing.T) {
	client, received := newUploadServer(t)

	_, _, err := client.ManagementAPI.ImportCoupons(context.Background(), 1, 2).
		UpFile(NamedReader(strings.NewReader(uploadCSV), "coupons.csv")).
		Execute()
	require.NoError(t, err)

	got := waitUpload(t, received)
	require.NoError(t, got.err)
	assert.Equal(t, "upFile", got.formName)
	assert.Equal(t, "coupons.csv", got.fileName, "Content-Disposition filename")
	assert.Equal(t, uploadCSV, got.content)
}

func TestUploadOSFileStillWorksAndIsNotClosed(t *testing.T) {
	client, received := newUploadServer(t)

	path := filepath.Join(t.TempDir(), "my-coupons.csv")
	require.NoError(t, os.WriteFile(path, []byte(uploadCSV), 0o600))
	f, err := os.Open(path)
	require.NoError(t, err)
	t.Cleanup(func() { _ = f.Close() })

	_, _, err = client.ManagementAPI.ImportCoupons(context.Background(), 1, 2).UpFile(f).Execute()
	require.NoError(t, err)

	got := waitUpload(t, received)
	assert.Equal(t, "my-coupons.csv", got.fileName)
	assert.Equal(t, uploadCSV, got.content)
	_, err = f.Seek(0, io.SeekStart)
	assert.NoError(t, err, "caller's *os.File was closed by the SDK")
}

type closeTrackingReader struct {
	io.Reader
	closed bool
}

func (c *closeTrackingReader) Close() error {
	c.closed = true
	return nil
}

func TestUploadDoesNotCloseReader(t *testing.T) {
	client, received := newUploadServer(t)
	r := &closeTrackingReader{Reader: strings.NewReader(uploadCSV)}

	_, _, err := client.ManagementAPI.ImportCoupons(context.Background(), 1, 2).UpFile(r).Execute()
	require.NoError(t, err)

	waitUpload(t, received)
	assert.False(t, r.closed, "SDK closed the caller's reader")
}

func TestUploadWithoutNameFallsBackToFieldName(t *testing.T) {
	client, received := newUploadServer(t)

	_, _, err := client.ManagementAPI.ImportCoupons(context.Background(), 1, 2).UpFile(strings.NewReader(uploadCSV)).Execute()
	require.NoError(t, err)

	got := waitUpload(t, received)
	assert.Equal(t, "upFile", got.fileName)
	assert.Equal(t, uploadCSV, got.content)
}

var errUploadRead = errors.New("disk on fire")

// failingReader returns some data, then an error.
type failingReader struct{ sent bool }

func (f *failingReader) Read(p []byte) (int, error) {
	if !f.sent {
		f.sent = true
		return copy(p, "value\nCOUP1\n"), nil
	}
	return 0, errUploadRead
}

func TestUploadReadErrorIsReturned(t *testing.T) {
	client, received := newUploadServer(t)

	_, _, err := client.ManagementAPI.ImportCoupons(context.Background(), 1, 2).
		UpFile(NamedReader(&failingReader{}, "coupons.csv")).
		Execute()
	require.ErrorIs(t, err, errUploadRead)

	// Nothing is sent, so a truncated file can never be imported.
	assertNoRequest(t, received)
}

func TestUploadNilFileReturnsErrorWithoutPanic(t *testing.T) {
	client, received := newUploadServer(t)

	var f *os.File // e.g. left nil after a failed os.Open
	_, _, err := client.ManagementAPI.ImportCoupons(context.Background(), 1, 2).UpFile(f).Execute()
	require.Error(t, err, "Execute with a nil *os.File returned no error")

	assertNoRequest(t, received)
}

func TestUploadNamedReaderWithEmptyNameFallsBackToFieldName(t *testing.T) {
	client, received := newUploadServer(t)

	_, _, err := client.ManagementAPI.ImportCoupons(context.Background(), 1, 2).
		UpFile(NamedReader(strings.NewReader(uploadCSV), "")).
		Execute()
	require.NoError(t, err)

	got := waitUpload(t, received)
	assert.Equal(t, "upFile", got.fileName)
	assert.Equal(t, uploadCSV, got.content)
}

func TestUploadNamedReaderWithNilReaderSendsNoFile(t *testing.T) {
	client, received := newUploadServer(t)

	// NamedReader(nil, ...) returns an untyped nil, so the SDK must skip the part
	// instead of panicking on a nil Read. The test server rejects a request without
	// a file part, so Execute may return an HTTP error here; what matters is that
	// no file part was sent and nothing panicked.
	_, _, _ = client.ManagementAPI.ImportCoupons(context.Background(), 1, 2).
		UpFile(NamedReader(nil, "coupons.csv")).
		Execute()

	got := waitUpload(t, received)
	assert.Empty(t, got.fileName, "server received a file part")
	assert.Empty(t, got.content, "server received a file part")
}

func TestUploadReaderReturningSmallChunks(t *testing.T) {
	client, received := newUploadServer(t)

	// The reader yields one byte per Read call; the whole content must still arrive.
	r := NamedReader(iotest.OneByteReader(strings.NewReader(uploadCSV)), "coupons.csv")
	_, _, err := client.ManagementAPI.ImportCoupons(context.Background(), 1, 2).UpFile(r).Execute()
	require.NoError(t, err)

	got := waitUpload(t, received)
	assert.Equal(t, "coupons.csv", got.fileName)
	assert.Equal(t, uploadCSV, got.content)
}
