.PHONY: all parser clean capnp-genrated capnp

all: fmt parser

test: fmt parser
	sh test.sh

parser: parser.go hintparser.go

karmem-generated: karmem
	karmem build --golang -o model/km model/km/*.km

capnp-generated: capnp
	capnp compile -I$(GOPATH)/src/capnproto.org/go/capnp/std -ogo model/capnp/*.capnp

proto-generated: protobuf
	protoc -I=model/proto --go_out=model model/proto/*.proto

capnp:
	go install capnproto.org/go/capnp/v3/capnpc-go@latest

protobuf:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

karmem:
	 go install karmem.org/cmd/karmem@latest

%arser.go: prefix = $(@:parser.go=)
%arser.go: %arser.y bin/goyacc
	@echo "bin/goyacc -o $@ -p yy$(prefix) -t $(prefix)Parser $<"
	@bin/goyacc -o $@ -p yy$(prefix) -t $(prefix)Parser $< || ( rm -f $@ && echo 'Please check y.output for more information' && exit 1 )
	@rm -f y.output

%arser_golden.y: %arser.y
	@bin/goyacc -fmt -fmtout $@ $<
	@(git diff --no-index --exit-code $< $@ && rm $@) || (mv $@ $< && >&2 echo "formatted $<" && exit 1)

bin/goyacc: goyacc/main.go goyacc/format_yacc.go
	GO111MODULE=on go build -o bin/goyacc goyacc/main.go goyacc/format_yacc.go

fmt: bin/goyacc parser_golden.y hintparser_golden.y
	@echo "gofmt (simplify)"
	@gofmt -s -l -w . 2>&1 | awk '{print} END{if(NR>0) {exit 1}}'

clean:
	go clean -i ./...
	rm -rf *.out
	rm -f parser.go hintparser.go
