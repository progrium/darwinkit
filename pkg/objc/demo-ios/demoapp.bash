#!/bin/bash
IBTOOL="xcrun ibtool"
SDK=$(xcode-select --print-path)/Platforms/iPhoneSimulator.platform/Developer/SDKs/iPhoneSimulator7.0.sdk/

# use -v for debugging (but also because there is a subtle race
# somewhere, causing the build to fail with missing UIKit symbols
# without it)
go build -ldflags "-extldflags -v" >/dev/null 2>/dev/null

${IBTOOL} --errors --warnings --notices --output-format human-readable-text \
	--compile demo.app/MainWindow.nib MainWindow.xib \
	--sdk ${SDK}

cp demo-ios demo.app/demo
