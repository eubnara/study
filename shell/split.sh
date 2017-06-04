#!/bin/bash

STR="a,b,c,d,e"
array=(${STR//,/ })


echo "${array[0]}"
echo "${array[1]}"
echo ${array[2]}
echo ${array[3]}
echo ${array[4]}
echo ${array[5]}
