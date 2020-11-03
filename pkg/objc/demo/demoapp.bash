#!/bin/bash
IBTOOL="xcrun ibtool"
SDK=$(xcode-select --print-path)/Platforms/MacOSX.platform/Developer/SDKs/MacOSX10.8.sdk

go build
${IBTOOL} --errors --warnings --notices --output-format human-readable-text \
	--compile demo.app/Contents/Resources/en.lproj/MainMenu.nib MainMenu.xib \
	--sdk ${SDK}
cp demo demo.app/Contents/MacOS/demo
