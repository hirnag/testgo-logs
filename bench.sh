go test -bench . -count 2 -benchtime 1s -benchmem>bench.log
sed -e "/Hello from /d" bench.log
rm -f bench.log
