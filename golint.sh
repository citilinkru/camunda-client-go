#!/bin/bash

golint ./... | \
 grep -v 'should have comment (or a comment on this block) or be unexported' | \
 grep -v 'should have comment or be unexported'
