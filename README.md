### Intellias Go Course

1. Git practice

- create 5 branches with unique commits  
- merge between at least 2 branches  
- merge all branches into main  

```

git log --pretty=format:'%h %s' --graph

*   e528668 Merge pull request #4 from 3u128/dev
|\  
| *   23e4ad1 Merge pull request #3 from 3u128/feature-add-root-dir
| |\  
| | *   c4736b9 Merge branch 'dev' into feature-add-root-dir
| | |\  
| | |/  
| |/|   
| * |   489d45b Merge pull request #2 from 3u128/feature-add-tls-support
| |\ \  
| | * | 3b42e98 add tls support and test by self-generated certificates
| |/ /  
| | * 9e93195 add root dir ./static with temp index.html
| |/  
* |   249bcd5 Merge pull request #1 from 3u128/feature-test
|\ \  
| * | b3981a7 just for test create/merge branches
|/ /  
* / efc612b fix syntax
|/  
* 0690364 copy from doc http server
* fbf4e70 Initial commit

```
