#!/bin/bash

gencheckBench(){
    echo "# GENCHECK $options"
    echo "# GENCHECK $options" > "$benchMD"
    echo "\`\`\`" >> "$benchMD"
    f="-f=./internal/benchmark/types.go"
    if [[ -n $options ]]; then
        gencheck "$f" "$options"
    else
        gencheck "$f"
    fi
    go test -v -run=XXX -bench=BenchmarkCompareGencheck -benchmem ./internal/benchmark >> "$benchMD"
    echo "\`\`\`" >> "$benchMD"
}

benchMD=benchmark_nooptions.md
options=
gencheckBench

benchMD=benchmark_noprealloc.md
options="--noprealloc"
gencheckBench

benchMD=benchmark_failfast.md
options="--failfast"
gencheckBench

benchMD=benchmark_playground.md

echo "# PLAYGROUND"
echo "# PLAYGROUND" > "$benchMD"
echo "\`\`\`" >> "$benchMD"
go test -v -run=XXX -bench=BenchmarkComparePlayground -benchmem ./internal/benchmark >> "$benchMD"
echo "\`\`\`" >> "$benchMD"