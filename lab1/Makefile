
sec1:
	$(SETENV) go test -v -run Sequential mapreduce/...

sec2:
	$(SETENV) go run src/wc/wc.go master sequential src/wc/pg-*.txt
	sort -n -k2 mrtmp.wcseq | tail -10 | diff - src/wc/mr-testout.txt

sec3:
	$(SETENV) go test -v -run TestBasic mapreduce/...

sec4:
	$(SETENV) go test -v -run Failure mapreduce/...

sec5:
	$(SETENV) go run src/ii/ii.go master sequential src/wc/pg-*.txt
	sort -k1,1 mrtmp.iiseq | sort -snk2,2 | grep -v '16' | tail -10 | diff - src/ii/mr-challenge.txt

SETENV=env "GOPATH=${PWD}"

clean:
	rm mrtmp.*
