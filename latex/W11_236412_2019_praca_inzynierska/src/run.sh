#!/bin/bash

rm -f img_*_it_*_best.png

go run main.go -from-file ./config.json
