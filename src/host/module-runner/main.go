package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

type Output struct {
	Stdout string `json:"stdout"`
}

type server struct{}

const serviceName = "workshop-host"

func main() {
	ctx := context.Background()

	// create a host Otel exporter (unrelated to the wasm traces)
	exporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithEndpoint("localhost:4317"), otlptracegrpc.WithInsecure())
	if err != nil {
		log.Fatalln("failed to create host tracer", err)
	}
	exporter.Start(ctx)
	defer exporter.Shutdown(ctx)

	// create host Otel tracer provider (unrelared to the wasm traces)
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resource.NewWithAttributes(semconv.SchemaURL, semconv.ServiceName(serviceName))),
	)
	otel.SetTracerProvider(tp)

	// create a server
	s := server{}
	http.Handle(
		"/",
		otelhttp.NewHandler(http.HandlerFunc(index), "Index"),
	)
	http.Handle(
		"/upload",
		otelhttp.NewHandler(http.HandlerFunc(upload), "Upload"),
	)
	http.Handle(
		"/run",
		otelhttp.NewHandler(http.HandlerFunc(s.runModule), "Run"),
	)

	log.Println("starting server on :3000")
	http.ListenAndServe(":3000", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	res.Header().Set("Content-Type", "application/text")
	res.Write([]byte("Hello, World!\n"))
}

func upload(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		res.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	mpFile, _, err := req.FormFile("wasm")
	if err != nil {
		log.Println("upload error:", err)
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Bad Request"))
		return
	}

	name := req.URL.Query().Get("name")
	if name == "" {
		log.Println("upload error:", err)
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Bad Request"))
		return
	}

	path := filepath.Join(os.TempDir(), name)
	tmpFile, err := os.Create(path)
	if err != nil {
		log.Println("file error:", err)
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte("Internal Service Error"))
		return
	}
	defer tmpFile.Close()

	n, err := io.Copy(tmpFile, mpFile)
	if err != nil {
		log.Println("copy error:", err)
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte("Internal Service Error"))
		return
	}

	fmt.Printf("Length of `%s` is %d bytes\n", path, n)
	res.WriteHeader(http.StatusOK)
}

func (s *server) runModule(res http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	if req.Method != http.MethodPost {
		res.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	name := req.URL.Query().Get("name")
	if name == "" {
		log.Println("name error: no name on url query")
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Bad Request"))
		return
	}

	path := filepath.Join(os.TempDir(), name)
	wasm, err := os.ReadFile(path)
	if err != nil {
		log.Println("name error: no module found", err)
		res.WriteHeader(http.StatusNotFound)
		res.Write([]byte("Not Found"))
		return
	}

	cfg := wazero.NewRuntimeConfig().WithCustomSections(true)
	rt := wazero.NewRuntimeWithConfig(ctx, cfg)
	wasi_snapshot_preview1.MustInstantiate(ctx, rt)
	output := &bytes.Buffer{}
	config := wazero.NewModuleConfig().WithStdin(req.Body).WithStdout(output).WithArgs(name)
	defer req.Body.Close()

	mod, err := rt.InstantiateWithConfig(ctx, wasm, config)
	if err != nil {
		log.Println("module instance error:", err)
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte("Internal Service Error"))
		return
	}
	defer mod.Close(ctx)

	res.WriteHeader(http.StatusOK)
	res.Header().Add("content-type", "application/json")
	res.Write(output.Bytes())
}
