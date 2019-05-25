#!/bin/sh

type="$1"
filename="$2"

dir="$(dirname $filename)"
fileNoExtension="$(basename -s .hdl -s .asm $filename)"

testfile="$dir/$fileNoExtension".tst
cmpfile="$dir/$fileNoExtension".cmp
outfile="$dir/$fileNoExtension".out

if [ ${type} = "hdl" ]; then
    ../tools/HardwareSimulator.sh $testfile
else
    ../tools/Assembler.sh $filename
    ../tools/CPUEmulator.sh $testfile
fi

success=$?

if [ "$success" -eq 0 ]; then
    echo "Success!"
    exit 0
fi

# Failed case
echo ""
echo "Expected:"
cat "$cmpfile"
echo ""
echo "Got:"
cat "$outfile"
