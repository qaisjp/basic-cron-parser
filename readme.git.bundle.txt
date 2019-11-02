# How to create a git bundle

Just run `git bundle create git.bundle master`

# How to import a git bundle

Assuming `git.bundle` lives in the `expr-parser` folder...

1. `git clone -b master expr-parser/git.bundle qaisjp`
2. `cd qaisjp`
3. `ls # files have been added`
4. `git log # with history``
