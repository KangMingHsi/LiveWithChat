APP=LiveWithChat

.PHONY: test-backend
test-backend:
	${MAKE} -C backend/api_gateway test
	${MAKE} -C backend/auth_subsystem test
	${MAKE} -C backend/stream_subsystem test
	${MAKE} -C backend/content_subsystem test

.PHONY: test-backend-without-io
test-backend-without-io:
	${MAKE} -C backend/api_gateway test-without-io
	${MAKE} -C backend/auth_subsystem test-without-io
	${MAKE} -C backend/stream_subsystem test-without-io
	${MAKE} -C backend/content_subsystem test-without-io