#!/bin/bash

tanzu plugin delete harbor
tanzu plugin source delete harbor-local-source
rm -r $HOME/.config/tanzu/harbor