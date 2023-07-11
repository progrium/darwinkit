
technologies.json:
	curl -sLo $@ https://developer.apple.com/tutorials/data/documentation/technologies.json

frameworks.json: technologies.json
	jq -r '.sections[] | select(.kind == "technologies") | .groups[] | select(.name == "App Services") | .technologies' $< > $@

frameworks.csv: technologies.json Makefile
	# print header
	echo "title,identifier,objc,swift" > $@
	jq -r '.sections[] | select(.kind == "technologies") | .groups[] | select(.name == "App Services") | .technologies[] | [.title, .destination.identifier, (.languages | contains(["objc"])), (.languages | contains(["swift"]))] | @csv' $< >> $@

frameworks.md: technologies.json render-framework-coverage.go Makefile
	go run render-framework-coverage.go > $@
