APP=LiveWithChat

.PHONY: test-backend
test-backend:
	${MAKE} -C backend/api_gateway test
	${MAKE} -C backend/auth_subsystem test
	${MAKE} -C backend/stream_subsystem test