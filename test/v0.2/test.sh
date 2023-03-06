cd $MG_ROOT

# initalize the repo-1 and repo-2 repositories
./multi-git --command init

# add a file to each repo
echo 123 > ./repo-1/data.txt
echo 123 > ./repo-2/data.txt

# Check the status. No file should be ready to commit yet
./multi-git --command status

# Add all files
./multi-git --command "add -A"

# Commit the files
./multi-git --command 'commit -m "added a file"'

# Check the log
./multi-git --command "log --oneline"
