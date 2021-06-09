APP=LiveWithChat

.PHONY: test-backend
test-backend:
	${MAKE} -C backend/api_gateway test
	${MAKE} -C backend/auth_subsystem test
	${MAKE} -C backend/stream_subsystem test
	${MAKE} -C backend/content_subsystem test
	${MAKE} -C backend/chat_subsystem test

.PHONY: test-backend-without-io
test-backend-without-io:
	${MAKE} -C backend/api_gateway test-without-io
	${MAKE} -C backend/auth_subsystem test-without-io
	${MAKE} -C backend/stream_subsystem test-without-io
	${MAKE} -C backend/content_subsystem test-without-io
	${MAKE} -C backend/chat_subsystem test-without-io

.PHONY: all-services-up
all-services-up:
	${MAKE} -C backend/api_gateway service-up
	${MAKE} -C backend/auth_subsystem service-up
	${MAKE} -C backend/stream_subsystem service-up
	${MAKE} -C backend/content_subsystem service-up
	${MAKE} -C frontend/chat_subsystem service-up
	${MAKE} -C frontend/live-with-chat service-up

.PHONY: all-services-down
all-services-down:
	${MAKE} -C backend/auth_subsystem service-down
	${MAKE} -C backend/stream_subsystem service-down
	${MAKE} -C backend/content_subsystem service-down
	${MAKE} -C frontend/chat_subsystem service-down
	${MAKE} -C frontend/live-with-chat service-down
	${MAKE} -C backend/api_gateway service-down
