#!/bin/bash

outFile="./docs/swagger/swagger.yaml"
workDir="./cmd"

echo "Generating api documentation..."

swagger generate spec -o $outFile -w $workDir -m

if [ $? -gt 0 ]; then 
    echo "Error when generating api documentation. Exited with code: $?"
    exit 1
else 
    cat $outFile
    echo ""
    echo "Generated successfully."
fi