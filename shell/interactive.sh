#!/bin/bash

case $- in
    *i*) echo interactive shell ;;
    *) echo non-interactive shell ;;
esac
