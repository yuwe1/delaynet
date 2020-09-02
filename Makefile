Build:
	docker build -t delayping .
	docker run --name delaypingv1 -p 2332:2332 -d delayping