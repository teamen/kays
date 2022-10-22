# kays

```bash
MYSQL_HOST=${KAYS_MYSQL_HOST} \
MYSQL_USER=${KAYS_MYSQL_USER} \
MYSQL_PWD=${KAYS_MYSQL_PWD} \
MYSQL_DB=${KAYS_MYSQL_DB} \
go test -timeout 30s \
-run ^TestCategoryCreateRootNode$ github.com/teamen/kays/internal/apiserver/store/mysql \
-v -count=1

```