.gitignore having only files which need to be included.
-------------------------------------------------------

# Why this is required ?

    Sometimes in a repository lots of supporting files gets generated along with source code. For example in this repo
    build is generating lots of files which are either binary or wasm (web assembly binary). In such case we need 
    .gitinclude type of system. 

# How this is done ?

    Although git don't provide such fucntionality by default, but we can always implement things by hiding all the 
    files by default and then adding only those files in with ! mark and adding them along