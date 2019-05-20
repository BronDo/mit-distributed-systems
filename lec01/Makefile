
sec1:
	env "GOPATH=${PWD}" go test -v -run Sequential mapreduce/...

sec2:
	env "GOPATH=${PWD}" go run src/wc/wc.go master sequential src/wc/pg-*.txt
	sort -n -k2 mrtmp.wcseq | tail -10 | diff - src/wc/mr-testout.txt

sec3:

clean:
	rm mrtmp.*
