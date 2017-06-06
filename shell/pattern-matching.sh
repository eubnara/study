#!/bin/bash

AA="inventory.tar.gz"

[[ $AA = *.tar.gz ]]; echo $?
[[ $AA = "*.tar.gz" ]]; echo $?



[[ $AA = inventory.tar.?? ]]; echo $?
[[ $AA = "inventory.tar.??" ]]; echo $?
