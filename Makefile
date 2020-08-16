unit-tests:
	go test ./adder

tag:
	git tag v0.1.0

tag-push:
	git push origin v0.1.0