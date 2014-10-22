MRUBY_DIR = mruby

all: submodules libmruby.a test

submodules:
	git submodule init
	git submodule update

libmruby.a:
	cd ${MRUBY_DIR} && (MRUBY_CONFIG=../build_config.rb make)

test:
	go test

clean:
	cd ${MRUBY_DIR} && make clean
	go clean
	rm ${CLI_NAME}
