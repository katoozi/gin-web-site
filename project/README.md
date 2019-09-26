# Test Project Configuration

## if you start project from `bash` set this environment variables

    export TEST_PROJECT_SERVER_HOST=127.0.0.1
    export TEST_PROJECT_SERVER_PORT=8000
    export TEST_PROJECT_MODE=release

    export TEST_PROJECT_DATABASE_DB_NAME=postgres
    export TEST_PROJECT_DATABASE_PASS=docker
    export TEST_PROJECT_DATABASE_PORT=5432
    export TEST_PROJECT_DATABASE_HOST=127.0.0.1
    export TEST_PROJECT_DATABASE_USER=postgres

    export TEST_PROJECT_RABBITMQ_HOST=127.0.0.1
    export TEST_PROJECT_RABBITMQ_USER=guest
    export TEST_PROJECT_RABBITMQ_PORT=5672
    export TEST_PROJECT_RABBITMQ_PASS=guest

    # connect to sentinel
    export TEST_PROJECT_REDIS_TYPE=sentinel
    export TEST_PROJECT_REDIS_SENTINELS=127.0.0.1:26379,127.0.0.2:26379,...
    export TEST_PROJECT_REDIS_SENTINELS_MASTER_NAME=mymaster

    # connect to simple redis instance
    export TEST_PROJECT_REDIS_TYPE=simple

    # redis required variables
    export TEST_PROJECT_REDIS_DB=0
    export TEST_PROJECT_REDIS_HOST=127.0.0.1
    export TEST_PROJECT_REDIS_PASS=

## if you using `docker-compose up` then create `.env` file and put this variables in it. `.env` will be ignored by git

    TEST_PROJECT_DATABASE_DB_NAME=postgres
    TEST_PROJECT_DATABASE_PASS=docker
    TEST_PROJECT_DATABASE_PORT=5432
    TEST_PROJECT_DATABASE_HOST=127.0.0.1
    TEST_PROJECT_DATABASE_USER=postgres

    TEST_PROJECT_SERVER_HOST=127.0.0.1
    TEST_PROJECT_SERVER_PORT=8000
    TEST_PROJECT_MODE=release

    TEST_PROJECT_RABBITMQ_HOST=127.0.0.1
    TEST_PROJECT_RABBITMQ_USER=guest
    TEST_PROJECT_RABBITMQ_PORT=5672
    TEST_PROJECT_RABBITMQ_PASS=guest

    TEST_PROJECT_REDIS_SENTINEL=sentinel_1
    TEST_PROJECT_REDIS_MASTER_NAME=mymaster
    TEST_PROJECT_REDIS_DB=0
    TEST_PROJECT_REDIS_HOST=127.0.0.1
    TEST_PROJECT_REDIS_PASS=
