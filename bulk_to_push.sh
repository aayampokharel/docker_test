 # alias push=./bulk_to_push.sh
 git add .
 commit_message="$*"
 git commit -m "${commit_message:-"fix(docker):revising docker"}"
 git push origin testing

